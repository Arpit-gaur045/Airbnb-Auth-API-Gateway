package controllers

import (
	"AuthInGo/services"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Logic to register a user using uc.UserService
	fmt.Println("Registering user in UserController")
	uc.UserService.CreateUser()
	w.Write([]byte("User registeration endpoint"))
}

