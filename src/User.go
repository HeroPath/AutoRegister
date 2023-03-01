package src

import (
	"strconv"
)

type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	ClassNpc string `json:"className"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUserRegister() []UserRegister {
	users := make([]UserRegister, 0)
	sheet := readFile("Users")

	for _, row := range sheet {
		if _, err := strconv.Atoi(row.Cells[0].String()); err == nil {
			user := UserRegister{}
			user.Username = row.Cells[1].String()
			user.Password = row.Cells[2].String()
			user.Email = row.Cells[3].String()
			user.ClassNpc = row.Cells[4].String()
			users = append(users, user)
		}
	}
	return users
}

func RegisterUsers(url string) {
	users := GetUserRegister()
	for _, user := range users {
		PostRequest(url+"auth/register", user, "")
	}
}

func LoginAdminUser(url string, username string, password string) interface{} {
	user := UserLogin{}
	user.Username = username
	user.Password = password
	response, statusCode := PostRequest(url+"auth/login", user, "")
	if statusCode == 200 {
		return response.(map[string]interface{})["token"]
	}
	return nil
}
