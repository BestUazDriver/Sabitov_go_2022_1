package main

import (
	"fmt"
	"web1/internal/core"
)

func main() {
	fmt.Println("hello")
	user := &core.User{
		Id:          1,
		Name:        "Ivan",
		Age:         20,
		NumberPhone: "1234",
	}
	user.ChangeNumber("+7238462")
	user.PrintInfo()
}
