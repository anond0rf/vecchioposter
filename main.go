package main

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/anond0rf/vecchioclient/client"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

//go:embed active.en.json active.it.json
var translationsFS embed.FS
var localizer *i18n.Localizer

var postThreadCmd *cobra.Command
var postReplyCmd *cobra.Command

var threadOpts = client.Thread{}
var replyOpts = client.Reply{}

var msgFile string
var userAgent string
var verbose bool

func init() {
	lang, err := GetOSLanguage()
	if err != nil || (!strings.EqualFold(lang, "it") && !strings.EqualFold(lang, "it-it") && !strings.EqualFold(lang, "it_it")) {
		lang = "en"
	}

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	bundle.LoadMessageFileFS(translationsFS, "active.en.json")
	bundle.LoadMessageFileFS(translationsFS, "active.it.json")

	localizer = i18n.NewLocalizer(bundle, lang)

	postThreadCmd = &cobra.Command{
		Use:     "new-thread",
		Short:   localize("NewThreadDescription", nil),
		Run:     newThread,
		Example: localize("NewThreadExample", nil),
		Aliases: []string{"thread", "nt", "newthread", "create-thread", "post-thread"},
	}

	postReplyCmd = &cobra.Command{
		Use:     "post-reply",
		Short:   localize("PostReplyDescription", nil),
		Run:     postReply,
		Example: localize("PostReplyExample", nil),
		Aliases: []string{"reply", "pr", "postreply", "add-reply", "post-reply"},
	}

	postThreadCmd.Flags().StringVarP(&threadOpts.Board, "board", "b", "", localize("FlagBoard", nil))
	postThreadCmd.Flags().StringVarP(&threadOpts.Name, "name", "n", "", localize("FlagName", nil))
	postThreadCmd.Flags().StringVarP(&threadOpts.Subject, "subject", "s", "", localize("FlagSubject", nil))
	postThreadCmd.Flags().StringVarP(&threadOpts.Email, "email", "e", "", localize("FlagEmail", nil))
	postThreadCmd.Flags().BoolVarP(&threadOpts.Spoiler, "spoiler", "S", false, localize("FlagSpoiler", nil))
	postThreadCmd.Flags().StringVarP(&threadOpts.Body, "body", "B", "", localize("FlagBody", nil))
	postThreadCmd.Flags().StringVarP(&msgFile, "msg-file", "m", "", localize("FlagMsgFile", nil))
	postThreadCmd.Flags().StringVarP(&threadOpts.Embed, "embed", "E", "", localize("FlagEmbed", nil))
	postThreadCmd.Flags().StringVarP(&threadOpts.Password, "password", "p", "", localize("FlagPassword", nil))
	postThreadCmd.Flags().BoolVarP(&threadOpts.Sage, "prevent-bump", "P", false, localize("FlagPreventBump", nil))
	postThreadCmd.Flags().StringSliceVarP(&threadOpts.Files, "files", "f", []string{}, localize("FlagFiles", nil))
	postThreadCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, localize("FlagVerbose", nil))
	postThreadCmd.Flags().StringVarP(&userAgent, "user-agent", "u", "", localize("FlagUserAgent", nil))

	postThreadCmd.MarkFlagRequired("board")

	postReplyCmd.Flags().IntVarP(&replyOpts.Thread, "thread", "t", 0, localize("FlagThread", nil))
	postReplyCmd.Flags().StringVarP(&replyOpts.Board, "board", "b", "", localize("FlagBoard", nil))
	postReplyCmd.Flags().StringVarP(&replyOpts.Name, "name", "n", "", localize("FlagName", nil))
	postReplyCmd.Flags().StringVarP(&replyOpts.Email, "email", "e", "", localize("FlagEmail", nil))
	postReplyCmd.Flags().BoolVarP(&replyOpts.Spoiler, "spoiler", "S", false, localize("FlagSpoiler", nil))
	postReplyCmd.Flags().StringVarP(&replyOpts.Body, "body", "B", "", localize("FlagBody", nil))
	postReplyCmd.Flags().StringVarP(&msgFile, "msg-file", "m", "", localize("FlagMsgFile", nil))
	postReplyCmd.Flags().StringVarP(&replyOpts.Embed, "embed", "E", "", localize("FlagEmbed", nil))
	postReplyCmd.Flags().StringVarP(&replyOpts.Password, "password", "p", "", localize("FlagPassword", nil))
	postReplyCmd.Flags().BoolVarP(&replyOpts.Sage, "prevent-bump", "P", false, localize("FlagPreventBump", nil))
	postReplyCmd.Flags().StringSliceVarP(&replyOpts.Files, "files", "f", []string{}, localize("FlagFiles", nil))
	postReplyCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, localize("FlagVerbose", nil))
	postReplyCmd.Flags().StringVarP(&userAgent, "user-agent", "u", "", localize("FlagUserAgent", nil))

	postReplyCmd.MarkFlagRequired("thread")
	postReplyCmd.MarkFlagRequired("board")
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "vecchioposter",
		Short: localize("ShortDescription", nil),
		Long:  localize("LongDescription", nil),
	}

	rootCmd.AddCommand(postThreadCmd)
	rootCmd.AddCommand(postReplyCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newThread(cmd *cobra.Command, args []string) {
	if err := validateFlags(threadOpts.Body, msgFile); err != nil {
		log.Fatalln(localize("ErrorInvalidFlags", map[string]interface{}{
			"Error": err,
		}))
	}

	if msgFile != "" {
		content, err := getFileContent(msgFile)
		if err != nil {
			log.Fatalln(localize("ErrorReadingFile", map[string]interface{}{
				"Error": err,
			}))
		}
		threadOpts.Body = content
	}

	conf := client.DefaultConfig
	if verbose {
		conf.Verbose = true
	}
	if userAgent != "" {
		conf.UserAgent = userAgent
	}
	vc := client.NewVecchioClientWithConfig(conf)

	id, err := vc.NewThread(threadOpts)
	if err != nil {
		log.Fatalln(localize("ErrorPostingThread", map[string]interface{}{
			"ThreadOpts": fmt.Sprintf("%+v", threadOpts),
			"Error":      err,
		}))
	}
	log.Println(localize("ThreadPostedSuccessfully", map[string]interface{}{
		"Id":         id,
		"ThreadOpts": fmt.Sprintf("%+v", threadOpts),
	}))
}

func postReply(cmd *cobra.Command, args []string) {
	if err := validateFlags(replyOpts.Body, msgFile); err != nil {
		log.Fatalln(localize("ErrorInvalidFlags", map[string]interface{}{
			"Error": err,
		}))
	}

	if msgFile != "" {
		content, err := getFileContent(msgFile)
		if err != nil {
			log.Fatalln(localize("ErrorReadingFile", map[string]interface{}{
				"Error": err,
			}))
		}
		replyOpts.Body = content
	}

	conf := client.DefaultConfig
	if verbose {
		conf.Verbose = true
	}
	if userAgent != "" {
		conf.UserAgent = userAgent
	}
	vc := client.NewVecchioClientWithConfig(conf)

	id, err := vc.PostReply(replyOpts)
	if err != nil {
		log.Fatalln(localize("ErrorPostingReply", map[string]interface{}{
			"ReplyOpts": fmt.Sprintf("%+v", replyOpts),
			"Error":     err,
		}))
	}
	log.Println(localize("ReplyPostedSuccessfully", map[string]interface{}{
		"Id":        id,
		"ReplyOpts": fmt.Sprintf("%+v", replyOpts),
	}))
}

func validateFlags(body, msgFile string) error {
	if body != "" && msgFile != "" {
		return errors.New(localize("ErrorBodyFlags", nil))
	}
	return nil
}

func getFileContent(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(absPath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func localize(messageID string, templateData map[string]interface{}) string {
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
	})
}
