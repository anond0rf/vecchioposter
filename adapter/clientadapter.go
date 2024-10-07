package adapter

import (
	"github.com/anond0rf/vecchioclient/client"
	"github.com/anond0rf/vecchioposter/model"
)

func ToClientThread(thread model.Thread) client.Thread {
	return client.Thread{
		Board:    thread.Board,
		Name:     thread.Name,
		Email:    thread.Email,
		Subject:  thread.Subject,
		Spoiler:  thread.Spoiler,
		Body:     thread.Body,
		Embed:    thread.Embed,
		Password: thread.Password,
		Sage:     thread.Sage,
		Files:    thread.Files,
	}
}

func ToClientReply(reply model.Reply) client.Reply {
	return client.Reply{
		Thread:   reply.Thread,
		Board:    reply.Board,
		Name:     reply.Name,
		Email:    reply.Email,
		Spoiler:  reply.Spoiler,
		Body:     reply.Body,
		Embed:    reply.Embed,
		Password: reply.Password,
		Sage:     reply.Sage,
		Files:    reply.Files,
	}
}
