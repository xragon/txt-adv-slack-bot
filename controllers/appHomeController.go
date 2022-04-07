package controllers

import (
	"log"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

// We create a sctucture to let us use dependency injection
type AppAdventureController struct {
	EventHandler *socketmode.SocketmodeHandler
}

func NewAppAdventureController(eventhandler *socketmode.SocketmodeHandler) AppAdventureController {
	c := AppAdventureController{
		EventHandler: eventhandler,
	}

	c.EventHandler.HandleEventsAPI(
		slackevents.Message,
		c.respondToMessage,
	)

	return c
}

func (c *AppAdventureController) respondToMessage(evt *socketmode.Event, clt *socketmode.Client) {
	// we need to cast our socketmode.Event into slackevents.AppHomeOpenedEvent
	evt_api, ok := evt.Data.(slackevents.EventsAPIEvent)

	clt.Ack(*evt.Request) // Acknowlege message otherwise it will retry

	if !ok {
		log.Printf("ERROR converting event to slackevents.EventsAPIEvent")
	}

	evt_app_message, ok := evt_api.InnerEvent.Data.(*slackevents.MessageEvent)

	if !ok {
		log.Printf("ERROR converting event to slackevents.MessageEvent: %v", ok)
	}

	if evt_app_message.User != "U03AN9C3NV7" {
		_, _, err := clt.GetApiClient().PostMessage(
			evt_app_message.Channel,
			slack.MsgOptionText("Hello World", false),
		)

		//Handle errors
		if err != nil {
			log.Printf("ERROR publishHomeTabView: %v", err)
		}
	}
}
