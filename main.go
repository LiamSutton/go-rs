package main

import (
	"fmt"
	"github.com/LiamSutton/go-rs/services"
	"github.com/LiamSutton/go-rs/utils"
	"log"
)

func main() {

	username, err := utils.ReadUsername()
	if err != nil {
		log.Fatalf("[ReadUsername]: %v\n", err)
	}

	runescapeProfile, err := services.GetRunescapeProfile(username)
	if err != nil {
		log.Fatalf("[GetRunescapeProfile]: %v\n", err)
	}

	fmt.Println(runescapeProfile.Name)

}
