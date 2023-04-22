package main

import (
	"github.com/conorkenn/dnd_go/character"
)

func main() {

	me := new(character.Character)
	me.Init("Suel", "human", "barbarian")
	me.Print()

}
