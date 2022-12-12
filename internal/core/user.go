package core

import (
	"fmt"
	"strings"
)

type User struct {
	Id          int
	Name        string
	Age         int
	Login       string
	Password    string
	NumberPhone string
	Posts       []*Post
}

func (user *User) ChangeNumber(newNumber string) {
	if strings.HasPrefix(newNumber, "+7") {
		user.NumberPhone = newNumber
	} else {
		user.NumberPhone = "0"
	}
}

func (user *User) PrintInfo() {
	fmt.Printf("Name : %v, Age : %v, Number : %v\n", user.Name, user.Age, user.NumberPhone)
}
