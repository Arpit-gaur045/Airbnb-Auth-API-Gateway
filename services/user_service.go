package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"
)

type UserService interface {
	GetUserById() error
	CreateUser() error
	LoginUser() error
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userReposiotry db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userReposiotry,
	}
}

func (u *UserServiceImpl) GetUserById() error {
	fmt.Println("Fetching user in UserService")
	//yaha jo create call hoga usmein mughe encrypted password paas karna hai
    //hashed, _ := bcrypt.GenerateFromPassword([]byte("Test@123"), 8)

    //password := string(hashed)
	//u.userRepository.Create(password)
	//u.userRepository.DeleteByID(3)
	u.userRepository.GetByID()
	return nil
}

func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("Creating user in UserService")

	password:="example_password"
	hashedPassword,err:= utils.HashPassword(password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return err
	}
	u.userRepository.Create(
		"username_example001",
		"user@example001.com",
		 hashedPassword,
	)
  return nil
}

func (u *UserServiceImpl) LoginUser() error {
	fmt.Println("Logging in user in UserService")
	response:=utils.CheckPasswordHash("example_password_wrong", "")
	fmt.Println("Login response:", response)
	return nil
}


