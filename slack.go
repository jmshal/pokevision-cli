package main

import (
	"net/http"
	"strings"
	"encoding/json"
)

type SlackPayload struct {
	Username string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	IconURL string `json:"icon_url"`
	Channel string `json:"channel"`
	Text string `json:"text"`
	Attachments []SlackAttachment `json:"attachments"`
	Markdown bool `json:"mrkdwn"`
}

type SlackAttachment struct {
	Fallback string `json:"fallback"`
	Color string `json:"color"`
	Pretext string `json:"pretext"`
	Title string `json:"title"`
	TitleLink string `json:"title_link"`
	Text string `json:"text"`
	Fields []SlackAttachmentField `json:"fields"`
	MarkdownIn []string `json:"mrkdwn_in"`
}

type SlackAttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool `json:"short"`
}

type SlackClient struct {
	WebhookURL string
}

func CreateSlackClient(webhookURL string) SlackClient {
	return SlackClient{webhookURL}
}

func (c *SlackClient) SendPayload(payload SlackPayload) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	reader := strings.NewReader(string(data))
	_, err = http.Post(c.WebhookURL, "application/json", reader)
	if err != nil {
		return err
	}
	// Maybe do something here like check response? For now it's OK
	return nil
}

//func (c *SlackClient) SendMessage(message string) error {
//	return c.SendPayload(SlackPayload{
//		Text: message,
//	})
//}
//
//func (c *SlackClient) SendMessageToChannel(channel, message string) error {
//	return c.SendPayload(SlackPayload{
//		Channel: channel,
//		Text: message,
//	})
//}