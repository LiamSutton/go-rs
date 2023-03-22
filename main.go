package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RunescapeProfileResponse struct {
	Magic            int `json:"magic"`
	Questsstarted    int `json:"questsstarted"`
	Totalskill       int `json:"totalskill"`
	Questscomplete   int `json:"questscomplete"`
	Questsnotstarted int `json:"questsnotstarted"`
	Totalxp          int `json:"totalxp"`
	Ranged           int `json:"ranged"`
	Activities       []struct {
		Date    string `json:"date"`
		Details string `json:"details"`
		Text    string `json:"text"`
	} `json:"activities"`
	Skillvalues []struct {
		Level int `json:"level"`
		Xp    int `json:"xp"`
		Rank  int `json:"rank"`
		ID    int `json:"id"`
	} `json:"skillvalues"`
	Name        string `json:"name"`
	Rank        string `json:"rank"`
	Melee       int    `json:"melee"`
	Combatlevel int    `json:"combatlevel"`
	LoggedIn    string `json:"loggedIn"`
}

func main() {
  r, err := http.Get("https://apps.runescape.com/runemetrics/profile/profile?user=Natria&activities=20")
  if err != nil {
    fmt.Println("No response from request")
  }

  defer r.Body.Close()
  body, err := ioutil.ReadAll(r.Body)

  var profile RunescapeProfileResponse
  if err := json.Unmarshal(body, &profile); err != nil {
    fmt.Println("Cannot unmarshal JSON")
  }
  
  prettyResponse := PrettyPrint(profile)
  fmt.Println(prettyResponse)
}

func PrettyPrint(i interface{}) string {
  s,_ := json.MarshalIndent(i, "", "\t")
  return string(s)
}

