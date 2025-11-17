package modrinth

import (
	"strconv"

	"github.com/charmbracelet/log"
)

func LoadMods(mods map[string]string) []Project {
	var modProjects []Project
	log.Info("Loading " + strconv.Itoa(len(mods)) + " mods")
	for name, id := range mods {
		project := fetchProject(name, id)
		if (project == nil) {
			log.Warn("Unable to find mod project, skipping...", "name", name, "id", id)
			continue
		}
		if (project.ProjectType != "mod") {
			log.Warn("Project is not a mod, skipping...", "name", name, "id", id)
			continue
		}
		log.Info("Loaded mod project", "name", name, "id", id)
		modProjects = append(modProjects, *project)
	}
	return modProjects 
}


