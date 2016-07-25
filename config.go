package main

type Config struct {
	Lat, Lon float64
	Name string
	Interval int
	ForceInitial bool
	Range int
	Forever bool
	IgnoreCommon bool
	Slack SlackConfig
	Notification NotificationConfig
}

type SlackConfig struct {
	Enable bool
	WebhookURL string
	Channel string
}

type NotificationConfig struct {
	// TODO
}