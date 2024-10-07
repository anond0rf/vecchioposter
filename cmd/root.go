package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   filepath.Base(os.Args[0]),
	Short: "VecchioPoster v1.0.0 - A CLI for posting on vecchiochan.com",
	Long:  "VecchioPoster v1.0.0 - A command-line tool to post new threads or replies on vecchiochan.com",
}

func Execute() error {
	return rootCmd.Execute()
}
