package adventure

import (
	"bytes"
	"embed"
	"encoding/json"
	"html/template"
	"io/fs"
	"io/ioutil"

	"github.com/slack-go/slack"
)

//go:embed messageTemplates
var messageAssets embed.FS

func RoomMessage(room Room) []slack.Block {

	// we need a stuct to hold template arguments
	// type args struct {
	// 	User string
	// }

	tpl := renderTemplate(messageAssets, "room.json", room)

	// we convert the view into a message struct
	view := slack.Msg{}

	str, _ := ioutil.ReadAll(&tpl)
	json.Unmarshal(str, &view)

	// We only return the block because of the way the PostEphemeral function works
	// we are going to use slack.MsgOptionBlocks in the controller
	return view.Blocks.BlockSet
}

func TextMessage(m string) []slack.Block {
	type args struct {
		Text string
	}
	tpl := renderTemplate(messageAssets, "simpleText.json", args{Text: m})

	// we convert the view into a message struct
	view := slack.Msg{}

	str, _ := ioutil.ReadAll(&tpl)
	json.Unmarshal(str, &view)

	// We only return the block because of the way the PostEphemeral function works
	// we are going to use slack.MsgOptionBlocks in the controller
	return view.Blocks.BlockSet
}

func renderTemplate(fs fs.FS, file string, args interface{}) bytes.Buffer {
	var tpl bytes.Buffer
	path := "messageTemplates/"
	// read the block-kit definition as a go template
	t, err := template.ParseFS(fs, path+file)
	if err != nil {
		panic(err)
	}
	// render the template using provided datas
	err = t.Execute(&tpl, args)
	if err != nil {
		panic(err)
	}
	return tpl
}
