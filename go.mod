module github.com/xragon/txt-adv-slack-bot

go 1.17

require (
	github.com/rs/zerolog v1.26.1
	github.com/slack-go/slack v0.10.2
)

require (
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
)

replace github.com/slack-go/slack => github.com/xnok/slack v0.8.1-0.20210509200330-9b2b404dbde9
