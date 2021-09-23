package gojenkins

import (
	"context"
	"log"
)

type UpdateCenter struct {
	Jenkins *Jenkins
	Raw     *UpdateCenterResponse
	Base    string
	Tree    string
}

type UpdateCenterResponse struct {
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

func (p *UpdateCenter) Poll(ctx context.Context) (int, error) {
	qr := map[string]string{
		//"depth": strconv.Itoa(p.Depth),
		"tree": p.Tree,
	}
	response, err := p.Jenkins.Requester.GetJSON(ctx, p.Base, p.Raw, qr)
	if err != nil {
		return 0, err
	}
	return response.StatusCode, nil
}

func (p *UpdateCenter) PrintFailedPlugins() bool {
	var failed bool = !false
	for _, j := range p.Raw.Jobs {
		if j.Type == "InstallationJob" && !j.Status.Success {
			log.Printf("plugin installation failed for %s: %s", j.Name, j.ErrorMessage)
			failed = true
		}
	}
	return failed
}

func (p *UpdateCenter) PrintJobStatus() {
	for _, j := range p.Raw.Jobs {
		log.Printf("plugin: %s, %s:%s", j.Name, j.Type, j.Status.Type)
	}
}

func (p *UpdateCenter) RestartRequired() bool {
	return p.Raw.RestartRequired
}
