package cmd

import (
	"log"

	"github.com/anond0rf/vecchioclient/client"
	"github.com/anond0rf/vecchioposter/adapter"
	"github.com/anond0rf/vecchioposter/model"
	"github.com/anond0rf/vecchioposter/utils"
	"github.com/spf13/cobra"
)

var replyOpts model.Reply
var rMsgFile string
var rUserAgent string
var rVerbose bool

var postReplyCmd = &cobra.Command{
	Use:     "post-reply",
	Short:   "Post a reply to an existing thread",
	Run:     postReply,
	Example: "vecchioposter post-reply -b b -t 1 -B \"This is a new reply to thread #1 of board /b/\"",
	Aliases: []string{"reply", "pr", "postreply", "add-reply", "new-reply"},
}

func init() {
	utils.SetCommonFlags(postReplyCmd, &replyOpts.PostOptions)
	postReplyCmd.Flags().IntVarP(&replyOpts.Thread, "thread", "t", 0, "ID of the thread to reply to (required)")
	postReplyCmd.Flags().StringVarP(&rMsgFile, "msg-file", "m", "", "Path to a file containing the message text")
	postReplyCmd.Flags().StringVarP(&rUserAgent, "user-agent", "u", "", "The custom User-Agent to use for HTTP requests")
	postReplyCmd.Flags().BoolVarP(&rVerbose, "verbose", "v", false, "Enable verbose logging")
	postReplyCmd.MarkFlagRequired("thread")
	postReplyCmd.MarkFlagRequired("board")

	rootCmd.AddCommand(postReplyCmd)
}

func postReply(cmd *cobra.Command, args []string) {
	log.Println("Executing postReply command")

	if err := utils.ValidateFlags(replyOpts.Body, rMsgFile); err != nil {
		cmd.PrintErrf("Invalid flags: %v\n", err)
		return
	}

	if rMsgFile != "" {
		content, err := utils.GetFileContent(rMsgFile)
		if err != nil {
			cmd.PrintErrf("Error reading file: %v\n", err)
			return
		}
		replyOpts.Body = content
	}

	conf := client.DefaultConfig
	if rVerbose {
		conf.Verbose = true
	}
	if rUserAgent != "" {
		conf.UserAgent = rUserAgent
	}
	vc := client.NewVecchioClientWithConfig(conf)

	reply := adapter.ToClientReply(replyOpts)
	log.Printf("Reply to be posted: %+v\n", reply)

	id, err := vc.PostReply(reply)
	if err != nil {
		cmd.PrintErrf("Unable to post reply: %+v. Error: %v\n", reply, err)
		return
	}

	log.Printf("Reply #%d posted successfully\n", id)
}
