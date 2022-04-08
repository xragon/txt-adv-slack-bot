package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"

	"github.com/slack-go/slack/socketmode"
)

// We create a sctucture to let us use dependency injection
type AppAdventureController struct {
	EventHandler *socketmode.SocketmodeHandler
	Rooms        []Rooms
}

func NewAppAdventureController(eventhandler *socketmode.SocketmodeHandler) AppAdventureController {
	file, err := ioutil.ReadFile("data/rooms.json")
	if err != nil {
		log.Printf("ERROR ReadFile: %v", err)
	}

	c := AppAdventureController{
		EventHandler: eventhandler,
	}

	err = json.Unmarshal([]byte(file), &c.Rooms)
	if err != nil {
		log.Printf("ERROR unmarshal: %v", err)
	}

	log.Printf("ERROR unmarshal: %v", c.Rooms)

	c.EventHandler.HandleEventsAPI(
		slackevents.Message,
		c.processMessage,
	)

	return c
}

type Rooms struct {
	ID          string // guid?
	Description string
	Exits       []Exits
}

// Exits:  n,e,s,w [10        ,11,0,34] ? zero as null/nothing?
// Exits2: array/slice of type exits, allowing for 0 to n exits from a room.

type Exits struct {
	Key         string // eg nrth/n
	Description string
	Destination string // destination key/id
}

type Players struct {
	Players map[string]Player
}

type Player struct {
	Name      string
	CurrentRooom string
}

type GameState struct {
	Players map[string]Player
	Rooms   []Rooms
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

	if evt_app_message.User == "U03AN9C3NV7" {
		return // do nothing if bots own message
	}

	command := evt_app_message.Text
	log.Printf("command is: %v", command)
	switch command {
	case "n":
		respondToMessage(clt, "you went north", evt_app_message.Channel)
	}
}

func respondToMessage(clt *socketmode.Client, message string, channel string) {
	log.Printf("respondToMessage Triggered")
	_, _, err := clt.GetApiClient().PostMessage(
		channel,
		slack.MsgOptionText(message, false),
	)
	//Handle errors
	if err != nil {
		log.Printf("ERROR publishHomeTabView: %v", err)
	}
}
