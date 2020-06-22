package user_service

import (
	"fmt"
)

type UserService struct {
	// inject the repository here
}

func  GetAllUsers() string {
	fmt.Println("********************************************")
	fmt.Println("Into the GetAllUsers function, USER SERVICE")
	fmt.Println("********************************************")
	return "asdasd"
}