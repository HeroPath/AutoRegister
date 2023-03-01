package main

import (
	"aoweb-auto-register/src"
	"fmt"
)

func main() {

	fmt.Println("Quests")
	src.GetQuests()

	fmt.Println("Items")
	src.GetItems()

	fmt.Println("Npcs")
	src.GetNpcs()
}
