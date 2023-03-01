package main

import (
	"aoweb-auto-register/src"
)

const url = "http://localhost:8000/api/v1/"

func main() {
	src.RegisterUsers(url)

	/*
		fmt.Println("Quests")
		src.GetQuests()

		fmt.Println("Items")
		src.GetItems()

		fmt.Println("Npcs")
		src.GetNpcs()

	*/
}
