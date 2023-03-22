package main

import (
	"fmt"
	"log"

	"github.com/LiamSutton/go-rs/services"
)

func main() {
  runescapeProfile,err := services.GetRunescapeProfile("Natria")
  if err != nil {
    log.Fatalf("[GetRunescapeProfile]: %v\n", err)
  }

  fmt.Println(runescapeProfile.Name)
}

// func PrettyPrint(i interface{}) string {
// 	s, _ := json.MarshalIndent(i, "", "\t")
// 	return string(s)
// }
