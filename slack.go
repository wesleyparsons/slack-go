package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Url string
}

type Message struct {
	MessageText string `json:text`
	MessageBlock []*MessageBlock `json:blocks`
}

type MessageBlock struct {
	BlockType string `json:type`
	BlockText BlockText `json:text`
}

type BlockText struct {
	Type string `json:type`
	Text string `json:text`
}

type SlackError struct {
	Code int
	Body string
}

func (e *SlackError) Error() string {
	return fmt.Sprintf("SlackError: %d %s", e.Code, e.Body)
}

func NewClient(url string) *Client {
	return &Client{url}
}

func (c *Client) SendMessage(msg *Message) error {

	body, _ := json.Marshal(msg)
	buf := bytes.NewReader(body)

	http.NewRequest("POST", c.Url, buf)
	resp, err := http.Post(c.Url, "application/json", buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	htmlData, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(htmlData))

	if resp.StatusCode != 200 {
		t, _ := ioutil.ReadAll(resp.Body)
		return &SlackError{resp.StatusCode, string(t)}
	}

	return nil
}

func (m *Message) NewMessageBlock() *MessageBlock {
	mb := &MessageBlock{}
	m.AddMessageBlock(mb)
	return mb
}

func (m *Message) AddMessageBlock(mb *MessageBlock) {
	m.MessageBlock = append(m.MessageBlock, mb)
}
