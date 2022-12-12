package main

import "C"
import (
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	println(viper.Get("PORT").(string))
}
