package utils

import (
	"errors"
	"os"

	"github.com/anond0rf/vecchioposter/model"
	"github.com/spf13/cobra"
)

func ValidateFlags(body, msgFile string) error {
	if body != "" && msgFile != "" {
		return errors.New("cannot use both -m and -M options at the same time")
	}
	return nil
}

func GetFileContent(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func SetCommonFlags(cmd *cobra.Command, opts *model.PostOptions) {
	cmd.Flags().StringVarP(&opts.Board, "board", "b", "", "Board name (e.g., 'b') (required)")
	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Name of the user")
	cmd.Flags().StringVarP(&opts.Email, "email", "e", "", "Email of the user")
	cmd.Flags().BoolVarP(&opts.Spoiler, "spoiler", "S", false, "Marks attached files as spoilers")
	cmd.Flags().StringVarP(&opts.Body, "body", "B", "", "The message of the post")
	cmd.Flags().StringVarP(&opts.Embed, "embed", "E", "", "Embed URL (YouTube, Spotify...)")
	cmd.Flags().StringVarP(&opts.Password, "password", "p", "", "Password used to delete or edit the post")
	cmd.Flags().BoolVarP(&opts.Sage, "prevent-bump", "P", false, "Replaces email with 'rabbia' and prevents bumping the thread")
	cmd.Flags().StringSliceVarP(&opts.Files, "files", "f", []string{}, "Paths of the files to attach")
}
