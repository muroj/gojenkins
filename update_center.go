package gojenkins

type UpdateCenter struct {
	Class           string `json:"_class"`
	Availables      []interface{}
	Jobs            []UpdateCenterJob `json:"jobs"`
	RestartRequired bool              `json:"restartRequiredForCompletion"`
	Sites           []UpdateSite
}

type UpdateCenterJob struct {
	Class        string `json:"_class"`
	ErrorMessage string `json:"errorMessage"`
	Id           string `json:"id"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	Status       struct {
		Class   string `json:"_class"`
		Success bool   `json:"success"`
		Type    string `json:"type"`
	}
}

type UpdateSite struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}
