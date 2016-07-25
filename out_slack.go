package main

import (
	"fmt"
)

func OutputToSlack(pokemon PokemonMeta, config Config) {
	slackClient := CreateSlackClient(config.Slack.WebhookURL)
	slackClient.SendPayload(SlackPayload{
		IconURL: pokemon.Icon,
		Markdown: true,
		Channel: config.Slack.Channel,
		Attachments: []SlackAttachment{
			{
				Title: pokemon.Name,
				TitleLink: pokemon.URL,
				Fallback: fmt.Sprintf("%v located near %v", pokemon.Name, pokemon.Location),
				Fields: []SlackAttachmentField{
					{Title: "Expires in", Value: HumanTime(pokemon.ExpiresAt), Short: true},
					{Title: "Location", Value: pokemon.Location, Short: true},
					{Title: "Distance", Value: fmt.Sprintf("%vm", pokemon.Distance), Short: true},
				},
			},
		},
	})
}