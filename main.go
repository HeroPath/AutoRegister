package main

import (
	"aoweb-auto-register/src"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	url := os.Getenv("URL_API")
	usernameAdmin := os.Getenv("USERNAME_ADMIN")
	passwordAdmin := os.Getenv("PASSWORD_ADMIN")

	//src.RegisterUsers(url)
	//time.Sleep(3 * time.Second)

	token := src.LoginAdminUser(url, usernameAdmin, passwordAdmin)

	if token != nil {
		src.RegisterItems(url, token.(string))
		src.RegisterNpcs(url, token.(string))
		src.RegisterQuests(url, token.(string))
	}
}
