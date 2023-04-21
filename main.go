package main

import "fmt"

func main() {
	hero := new(Character)
	hero.init("Suel", DRAGONBORN)

	fmt.Println(*hero)

	r := roll("1d20")
	fmt.Println(r)
}
