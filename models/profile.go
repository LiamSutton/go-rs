package profile

type RunescapeProfile struct {
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

func ProfileIsValid(p RunescapeProfile) bool { return len(p.Name) != 0 }
