package modrinth

type (
	Project struct {
		Slug string
		Id string
		Title string
		Description string
		ClientSide string `json:"client_side"`
		ServerSide string `json:"server_side"`
		Status string
		RequestedStatus string `json:"requested_status"`
		ProjectType string `json:"project_type"`
		Versions []string
		GameVersions []string `json:"game_versions"`
		Loaders []string
	}

	Version struct {
		ProjectId string `json:"project_id"`
		Name string
		Id string
		VersionNumber string `json:"version_number"`
		GameVersions []string `json:"game_versions"`
		VersionType string `json:"version_type"`
		Loaders []string
		Files []File
	}

	File struct {
		Hashes Hash
		Url string
		Filename string
		Primary bool
		Size int
	}

	Hash struct {
		Sha512 string
		Sha1 string
	}
)
