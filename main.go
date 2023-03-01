package main

import (
	"aoweb-auto-register/src"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// user dotenv
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	url := os.Getenv("URL_API")
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
