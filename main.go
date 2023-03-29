package main

import (
	"AutoRegister/src"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	url := os.Getenv("URL_API")
	usernameAdmin := os.Getenv("USERNAME_ADMIN")
	passwordAdmin := os.Getenv("PASSWORD_ADMIN")

	fmt.Println("Registering...")

	src.RegisterUsers(url)
	time.Sleep(500 * time.Millisecond)

	token := src.LoginAdminUser(url, usernameAdmin, passwordAdmin)

	if token != nil {
		results := make(chan string, 3)
		go func() {
			results <- src.RegisterItems(url, token.(string))
		}()
		go func() {
			results <- src.RegisterNpcs(url, token.(string))
		}()
		go func() {
			results <- src.RegisterQuests(url, token.(string))
		}()
		for i := 0; i < 3; i++ {
			<-results
		}
	}
}
