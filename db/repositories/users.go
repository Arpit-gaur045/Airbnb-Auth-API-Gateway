package db

import (
	"fmt"
)

type UserRepository interface {
	Create() error
}

type UserRepositoryImpl struct {
	
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{
		
	}
}

func (u *UserRepositoryImpl) Create() error {
	fmt.Println("Creating user in UserRepository")
	return nil
}
