package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	godotenv.Load("../.env")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelId, timestamp, err := api.PostMessage(
		"C0811AVLN9W",
		slack.MsgOptionText("Hello World", false),
	)

	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	fmt.Printf("Message sent successfully to channel %s at %s", channelId, timestamp)
}
