//go:build windows

package main

import (
	"errors"
	"unicode/utf16"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	LOCALE_NAME_USER_DEFAULT = 0
	LOCALE_SNAME             = 0x0000005c
)

func GetOSLanguage() (string, error) {
	lang, err := getWindowsLanguage()
	if err != nil {
		lang, regErr := getLanguageFromRegistry()
		if regErr != nil {
			return "", errors.New(localize("ErrorWinRetrieveLang", map[string]interface{}{
				"Err":   err,
				"Error": regErr,
			}))
		}
		return lang, nil
	}
	return lang, nil
}

func getWindowsLanguage() (string, error) {
	kernel32 := windows.NewLazyDLL("Kernel32.dll")
	procGetLocaleInfoEx := kernel32.NewProc("GetLocaleInfoEx")

	var buf [85]uint16
	ret, _, err := procGetLocaleInfoEx.Call(
		uintptr(LOCALE_NAME_USER_DEFAULT),
		uintptr(LOCALE_SNAME),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
	)

	if ret == 0 {
		return "", errors.New(localize("ErrorWinLocaleInfo", map[string]interface{}{
			"Error": err,
		}))

	}

	return utf16ToString(buf[:]), nil
}

func utf16ToString(buf []uint16) string {
	n := 0
	for n < len(buf) && buf[n] != 0 {
		n++
	}
	return string(utf16.Decode(buf[:n]))
}

func getLanguageFromRegistry() (string, error) {
	var key windows.Handle
	err := windows.RegOpenKeyEx(windows.HKEY_CURRENT_USER, windows.StringToUTF16Ptr(`Control Panel\International`), 0, windows.KEY_READ, &key)
	if err != nil {
		return "", errors.New(localize("ErrorWinOpenRegKey", map[string]interface{}{
			"Error": err,
		}))
	}
	defer windows.RegCloseKey(key)

	var buf [256]uint16
	var bufLen uint32 = uint32(len(buf) * 2)
	err = windows.RegQueryValueEx(key, windows.StringToUTF16Ptr("LocaleName"), nil, nil, (*byte)(unsafe.Pointer(&buf[0])), &bufLen)
	if err != nil {
		return "", errors.New(localize("ErrorWinQueryRegKey", map[string]interface{}{
			"Error": err,
		}))
	}

	return utf16ToString(buf[:]), nil
}
