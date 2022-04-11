package adventure

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
)

type GameState struct {
	Players
	Rooms []Rooms
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
	Name         string
	CurrentRooom string
}

func NewAdventure() GameState {
	var rooms Rooms
	file, err := ioutil.ReadFile("data/rooms.json")
	if err != nil {
		log.Printf("ERROR Reading Rooms File: %v", err)
	}

	err = json.Unmarshal([]byte(file), &rooms)
	if err != nil {
		log.Printf("ERROR unmarshal: %v", err)
	}

	log.Printf("ERROR unmarshal: %v", rooms)

	gs := GameState{
		Rooms: rooms,
	}

	return gs
}

func (gs *GameState) ProcessCommand(u string, cmd string) string {
	moves := []string{"north", "n", "east", "e", "south", "s", "west", "w"}
	// type move struct {
	// 	North []string
	// 	East  []string
	// 	South []string
	// 	West  []string
	// }

	// gameMoves := new(move)

	// gameMoves.North = append(gameMoves.North, "North", "n")
	// gameMoves.East = append(gameMoves.North, "East", "e")
	// gameMoves.South = append(gameMoves.North, "South", "s")
	// gameMoves.West = append(gameMoves.North, "West", "w")

	if itemExists(moves, cmd) {
		player := gs.Players.Players[u]

		room := gs.Rooms
	}

	return cmd
}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		return false
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
