package adventure

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
	"strings"

	"github.com/slack-go/slack"
)

type GameState struct {
	Players map[string]*Player
	Rooms   map[string]*Room
}

type Room struct {
	Description string
	Exits       map[string]Exit
}

type Exit struct {
	Description string
	Destination string // destination key/id
}

type Player struct {
	Name        string
	CurrentRoom string
}

func NewAdventure() GameState {
	log.Printf("NEW ADVENTURE!!!")
	var gs GameState
	file, err := ioutil.ReadFile("data/rooms.json")
	if err != nil {
		log.Printf("ERROR Reading Rooms File: %v", err)
	}

	err = json.Unmarshal([]byte(file), &gs.Rooms)
	if err != nil {
		log.Printf("ERROR unmarshal: %v", err)
	}

	file, err = ioutil.ReadFile("data/players.json")
	if err != nil {
		log.Printf("ERROR Reading Rooms File: %v", err)
	}

	err = json.Unmarshal([]byte(file), &gs.Players)
	if err != nil {
		log.Printf("ERROR unmarshal: %v", err)
	}

	return gs
}

func (gs *GameState) ProcessCommand(u string, cmd string) []slack.Block {
	moves := []string{"north", "n", "east", "e", "south", "s", "west", "w"}
	r := "Command Not Understood"

	player := gs.Players[u]
	log.Printf("PLAYER ProcessCommand: %v", player)

	room := gs.Rooms[player.CurrentRoom]
	log.Printf("ROOM ProcessCommand: %v", room)
	log.Printf("Exits ProcessCommand: %v", room.Exits)

	if itemExists(moves, cmd) {
		switch cmd {
		case "north", "n":
			player.move("n", room)
		case "east", "e":
			player.move("e", room)
		case "south", "s":
			player.move("s", room)
		case "west", "w":
			player.move("w", room)
		}

		room = gs.Rooms[player.CurrentRoom]
		log.Printf("ROOM ProcessCommand: %v", &room)
		r = "You move " + cmd + ". You have entered: " + room.Description
	}

	if cmd == "look" {
		blocks := RoomMessage(*room)
		return blocks
	}

	log.Printf("Response ProcessCommand: %v", r)
	return TextMessage(r)
}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Slice {
		return false
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func (p *Player) move(cmd string, r *Room) {
	newRoom := r.Exits[cmd]
	if newRoom != (Exit{}) {
		p.CurrentRoom = newRoom.Destination
	}
}

func (r *Room) listExists() string {
	log.Printf("Exits: %v", r.Exits)
	exits := make([]string, 0, len(r.Exits))
	for k := range r.Exits {
		exits = append(exits, k)
	}
	return strings.Join(exits, ",")
}
