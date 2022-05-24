package slack

import "github.com/ashwanthkumar/slack-go-webhook"

func sendLogs(webhook string, message string) []error {
	payload := slack.Payload{
		Text:     message,
		Username: "Apex",
	}
	err := slack.Send(webhook, "", payload)
	return err
}
