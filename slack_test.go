package slack

import (
	"flag"
	"testing"
)

var (
	client *Client
)

func init() {
	client = &Client{}
	flag.StringVar(&client.Url, "url", "", "webhook url")
	flag.Parse()
	if client.Url == "" {
		flag.PrintDefaults()
		panic("\n=================\nYou need to specify -url flag\n=================\n\n")
	}
}

func TestSendMessage(t *testing.T) {
	msg := &Message{}
	msg.MessageText = "Slack API Test from go"
	client.SendMessage(msg)
}

func TestSendMessageWithBlock(t *testing.T) {
	msg := &Message{}
	msg.MessageText = "This is a test message using Blocks in Go"

	block := msg.NewMessageBlock()
	block.BlockType = "section"
	block.BlockText.Type = "mrkdwn"
	block.BlockText.Text = "This message was sent from Go using Blocks"

	client.SendMessage(msg)
}
