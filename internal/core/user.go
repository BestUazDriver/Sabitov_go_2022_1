package core

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id, omitempty"`
	Name        string             `bson:"name,omitempty"`
	Age         int                `bson:"age,omitempty"`
	Login       string             `bson:"login,omitempty"`
	Password    string             `bson:"password,omitempty"`
	NumberPhone string             `bson:"number_phone,omitempty"`
	Posts       []*Post            `bson:"posts,omitempty"`
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
