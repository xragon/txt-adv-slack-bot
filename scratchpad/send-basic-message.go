package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	// params := {
	// 	TeamID:
	// }

	// api.ope

	params := slack.GetConversationsParameters{Types: []string{"im"}}

	channels, _, err := api.GetConversations(&params)

	// channelID, timestamp, err := api.PostMessage(
	// 	"C03A30QM5U6",
	// 	slack.MsgOptionText("Hello World", false),
	// )

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("%d\n", len(channels))

	for _, v := range channels {
		fmt.Printf(v.ID)
		fmt.Printf("%+v\n", v)
	}

	// fmt.Printf("Message sent succesfully to channel %s at %s", channelID, timestamp)
	// fmt.Printf(channels)
}
