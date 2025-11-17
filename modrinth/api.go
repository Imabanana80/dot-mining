package modrinth

import (
	"encoding/json"
	"io"
	"net/http"
)

func fetchProject(name string, id string) *Project {
	var project *Project
	url := "https://api.modrinth.com/v2/project/" + id
	res, _ := http.Get(url)
	if (res.StatusCode == 404) {
		return project
	}
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &project)
	return project
}
