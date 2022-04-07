package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/slack-go/slack/socketmode"
	"github.com/xragon/txt-adv-slack-bot/controllers"
	"github.com/xragon/txt-adv-slack-bot/drivers"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Instanciate deps
	client, err := drivers.ConnectToSlackViaSocketmode()
	if err != nil {
		log.Error().
			Str("error", err.Error()).
			Msg("Unable to connect to slack")

		os.Exit(1)
	}

	// Inject Deps in router
	socketmodeHandler := socketmode.NewsSocketmodeHandler(client)

	controllers.NewAppAdventureController(socketmodeHandler)

	socketmodeHandler.RunEventLoop()

}
