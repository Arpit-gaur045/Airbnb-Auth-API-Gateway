package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"
     env "AuthInGo/config/env"
	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById() error
	CreateUser() error
	LoginUser() (string,error)
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

func (u *UserServiceImpl) LoginUser() (string,error) {
	//Pre-requisite: This function will be given email and password as 
	//parameter, which we can hardcode for now.
	email:="user@example001.com"
	password:="example_password"

	//Step 1: Make a repository call to get user by email
	user,err:=u.userRepository.GetByEmail(email)
	//Step 2: if user exists, or not. if not exists,return error
	if err != nil {
		fmt.Println("Error fetching user by email:", err)
		return "",err
	}

	if user == nil {
		fmt.Println("no user found with the given email")
		return "",err
	}

	//Step 3: if user exists, check the password using utils.CheckPasswordHash
	isPasswordValid := utils.CheckPasswordHash(password, user.Password)

	// Step 4: if password matches, return a JWT token , else return error saying password does not match
	if(!isPasswordValid) {
		fmt.Println("Password does not match")
		return "", err
	}

	// token, err := createToken(user.Username)
	// if err != nil {
	// 	fmt.Println("Error creating token:", err)
	// 	return "", err
	// }
	// fmt.Println("Token created successfully:", token)

    jwtPayload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)

	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "TOKEN")))

	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}

	fmt.Println("JWT Token:", tokenString)

	return tokenString, nil

}


