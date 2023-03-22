package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/LiamSutton/go-rs/models"
)

func GetRunescapeProfile(username string) (*profile.RunescapeProfile, error) {

	r, err := http.Get("https://apps.runescape.com/runemetrics/profile/profile?user=Natria&activities=20")
	if err != nil {
		return nil, errors.New("Request failed.")
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	p := profile.RunescapeProfile{}
	if err := json.Unmarshal(body, &p); err != nil {
    return nil, errors.New("Failed to unmarshal JSON %v")
	}
  
  isValid := profile.ProfileIsValid(p)
  
  if !isValid {
    return nil, errors.New("Failed to retrieve profile")
  }

  return &p, nil
}
