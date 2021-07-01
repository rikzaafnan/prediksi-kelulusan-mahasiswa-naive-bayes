package models

type User struct {
	ID       int    `json:"id" form:"id"`
	Email    string `json:"email" form:"email"`
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

type UserResponse struct {
	ID    int    `json:"id" form:"id"`
	Email string `json:"email" form:"email"`
	Name  string `json:"name" form:"name"`
	Token string `json:"token" form:"token"`
}
