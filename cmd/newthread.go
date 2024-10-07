package cmd

import (
	"log"

	"github.com/anond0rf/vecchioclient/client"
	"github.com/anond0rf/vecchioposter/adapter"
	"github.com/anond0rf/vecchioposter/model"
	"github.com/anond0rf/vecchioposter/utils"
	"github.com/spf13/cobra"
)

var threadOpts model.Thread
var tMsgFile string
var tUserAgent string
var tVerbose bool

var newThreadCmd = &cobra.Command{
	Use:     "new-thread",
	Short:   "Post a new thread",
	Run:     newThread,
	Example: "vecchioposter new-thread -b b -B \"This is a new thread on board /b/\" -f file.jpg",
	Aliases: []string{"thread", "nt", "newthread", "create-thread", "post-thread"},
}

func init() {
	utils.SetCommonFlags(newThreadCmd, &threadOpts.PostOptions)
	newThreadCmd.Flags().StringVarP(&threadOpts.Subject, "subject", "s", "", "Subject of the thread")
	newThreadCmd.Flags().StringVarP(&tMsgFile, "msg-file", "m", "", "Path to a file containing the message text")
	newThreadCmd.Flags().StringVarP(&tUserAgent, "user-agent", "u", "", "The custom User-Agent to use for HTTP requests")
	newThreadCmd.Flags().BoolVarP(&tVerbose, "verbose", "v", false, "Enable verbose logging")
	newThreadCmd.MarkFlagRequired("board")

	rootCmd.AddCommand(newThreadCmd)
}

func newThread(cmd *cobra.Command, args []string) {
	log.Println("Executing newThread command")

	if err := utils.ValidateFlags(threadOpts.Body, tMsgFile); err != nil {
		cmd.PrintErrf("Invalid flags: %v\n", err)
		return
	}

	if tMsgFile != "" {
		content, err := utils.GetFileContent(tMsgFile)
		if err != nil {
			cmd.PrintErrf("Error reading file: %v\n", err)
			return
		}
		threadOpts.Body = content
	}

	conf := client.DefaultConfig
	if tVerbose {
		conf.Verbose = true
	}
	if tUserAgent != "" {
		conf.UserAgent = tUserAgent
	}
	vc := client.NewVecchioClientWithConfig(conf)

	thread := adapter.ToClientThread(threadOpts)
	log.Printf("Thread to be posted: %+v\n", thread)

	id, err := vc.NewThread(thread)
	if err != nil {
		cmd.PrintErrf("Unable to post thread: %+v. Error: %v\n", thread, err)
		return
	}

	log.Printf("Thread #%d posted successfully\n", id)
}
