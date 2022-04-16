package controllers

import (
	"log"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"github.com/xragon/txt-adv-slack-bot/adventure"
)

// We create a sctucture to let us use dependency injection
type AppAdventureController struct {
	EventHandler *socketmode.SocketmodeHandler
	GameState    adventure.GameState
}

func NewAppAdventureController(eventhandler *socketmode.SocketmodeHandler) AppAdventureController {

	log.Printf("NEW CONTROLLER!!!")

	c := AppAdventureController{
		EventHandler: eventhandler,
		GameState:    adventure.NewAdventure(),
	}

	c.EventHandler.HandleEventsAPI(
		slackevents.Message,
		c.processMessage,
	)

	return c
}

func (c *AppAdventureController) processMessage(evt *socketmode.Event, clt *socketmode.Client) {
	// we need to cast our socketmode.Event into slackevents.MessageEvent
	evt_api, ok := evt.Data.(slackevents.EventsAPIEvent)

	clt.Ack(*evt.Request) // Acknowlege message otherwise it will retry

	if !ok {
		log.Printf("ERROR converting event to slackevents.EventsAPIEvent")
	}

	evt_app_message, ok := evt_api.InnerEvent.Data.(*slackevents.MessageEvent)

	if !ok {
		log.Printf("ERROR converting event to slackevents.MessageEvent: %v", ok)
	}

	// "U03AN9C3NV7"
	if evt_app_message.User == "U03C11FNTH7" {
		return // do nothing if bots own message
	}

	command := evt_app_message.Text
	log.Printf("command is: %v", command)

	respondWithBlock(clt, c.GameState.ProcessCommand(evt_app_message.User, strings.ToLower(command)), evt_app_message.Channel)
}

func respondToMessage(clt *socketmode.Client, message string, channel string) {
	log.Printf("respondToMessage Triggered")
	_, _, err := clt.GetApiClient().PostMessage(
		channel,
		slack.MsgOptionText(message, false),
	)

	if err != nil {
		log.Printf("ERROR respondToMessage: %v", err)
	}
}

func respondWithBlock(clt *socketmode.Client, blocks []slack.Block, channel string) {
	_, _, err := clt.GetApiClient().PostMessage(
		channel,
		slack.MsgOptionBlocks(blocks...),
	)

	if err != nil {
		log.Printf("ERROR postGreetingMessage: %v", err)
	}
}
