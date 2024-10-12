//go:build linux || darwin

package main

import (
	"errors"
	"os"
	"strings"
)

func GetOSLanguage() (string, error) {
	envVars := []string{"LANG", "LC_ALL", "LC_MESSAGES", "LANGUAGE"}

	for _, envVar := range envVars {
		if lang := os.Getenv(envVar); lang != "" {
			return strings.Split(lang, ".")[0], nil
		}
	}

	return "", errors.New(localize("ErrorUnixRetrieveLang", nil))
}
