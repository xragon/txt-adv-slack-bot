package controllers

import (
	"fmt"
	"log"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

// We create a sctucture to let us use dependency injection
type AppHomeController struct {
	EventHandler *socketmode.SocketmodeHandler
}

func NewAppHomeController(eventhandler *socketmode.SocketmodeHandler) AppHomeController {
	c := AppHomeController{
		EventHandler: eventhandler,
	}

	// c.EventHandler.Handle(socketmode.EventTypeErrorBadMessage, c.recoverAppHomeOpened)

	// App Home (2)
	c.EventHandler.HandleEventsAPI(
		slackevents.Message,
		c.publishHomeTabView,
	)

	return c

}

func (c *AppHomeController) publishHomeTabView(evt *socketmode.Event, clt *socketmode.Client) {
	// we need to cast our socketmode.Event into slackevents.AppHomeOpenedEvent
	evt_api, ok := evt.Data.(slackevents.EventsAPIEvent)

	clt.Ack(*evt.Request)

	fmt.Printf("PROCESSING MESSAGE RECIEVED")

	if ok != true {
		log.Printf("ERROR converting event to slackevents.EventsAPIEvent")
	}

	evt_app_home_opened, ok := evt_api.InnerEvent.Data.(*slackevents.MessageEvent)

	if ok != true {
		log.Printf("ERROR converting event to slackevents.MessageEvent: %v", ok)
	}

	fmt.Printf(evt_app_home_opened.User)
	fmt.Printf("%t", ok)

	// var user string

	// if ok != true {
	// 	log.Printf("ERROR converting inner event to slackevents.AppHomeOpenedEvent")
	// 	//Patch the fact that we are not able to cast evt_api.InnerEvent.Data to AppHomeOpenedEvent
	// 	user = reflect.ValueOf(evt_api.InnerEvent.Data).Elem().FieldByName("User").Interface().(string)
	// } else {
	// 	user = evt_app_home_opened.User
	// }

	// log.Printf("ERROR publishHomeTabView: %v", evt_app_home_opened)

	// // create the view using block-kit
	// view := views.AppHomeTabView()

	// // Publish the view (3)
	// // We get the Api client from `clt` and post our view
	// _, err := clt.GetApiClient().PublishView(user, view, "")

	if evt_app_home_opened.User != "U03AN9C3NV7" {
		_, _, err := clt.GetApiClient().PostMessage(
			evt_app_home_opened.Channel,
			slack.MsgOptionText("Hello World", false),
		)

		//Handle errors
		if err != nil {
			log.Printf("ERROR publishHomeTabView: %v", err)
		}
	}

}
