package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelID, timestamp, err := api.PostMessage(
		"C03A30QM5U6",
		slack.MsgOptionText("Hello World", false),
	)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message sent succesfully to channel %s at %s", channelID, timestamp)
}
