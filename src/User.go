package src

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
