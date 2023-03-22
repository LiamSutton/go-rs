package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/LiamSutton/go-rs/models"
)


func main() {
	r, err := http.Get("https://apps.runescape.com/runemetrics/profile/profile?user=Natria&activities=20")
	if err != nil {
    log.Fatalf("Request failed: %v", err)
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

  p := profile.RunescapeProfile{}
	if err := json.Unmarshal(body, &p); err != nil {
    log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
  
  isValid := profile.ProfileIsValid(p)
  if isValid {
    fmt.Printf("Retrieved profile: %s\n", p.Name)
  } else {
    log.Fatalf("Failed to retrieve profile.")
  }
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
