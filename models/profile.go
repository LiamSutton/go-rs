package profile

type RunescapeProfile struct {
	Magic            int          `json:"magic"`
	QuestsStarted    int          `json:"questsstarted"`
	Totalskill       int          `json:"totalskill"`
	QuestsComplete   int          `json:"questscomplete"`
	QuestsNotStarted int          `json:"questsnotstarted"`
	TotalXP          int          `json:"totalxp"`
	Ranged           int          `json:"ranged"`
	Activities       []Activity   `json:"activities"`
	SkillValues      []SkillValue `json:"skillvalues"`
	Name             string       `json:"name"`
	Rank             string       `json:"rank"`
	Melee            int          `json:"melee"`
	Combatlevel      int          `json:"combatlevel"`
	LoggedIn         string       `json:"loggedIn"`
}

type Activity struct {
	Date    string `json:"date"`
	Details string `json:"details"`
	Text    string `json:"text"`
}

type SkillValue struct {
	Level int `json:"level"`
	XP    int `json:"xp"`
	Rank  int `json:"rank"`
	ID    int `json:"id"`
}

func ProfileIsValid(p RunescapeProfile) bool { return len(p.Name) != 0 }
