package main

import (
	"fmt"

	"github.com/conorkenn/dnd_go/character"
)

func main() {

	me := new(character.Character)
	me.Init("Suel", "gnome")

	fmt.Println(*me)

}
