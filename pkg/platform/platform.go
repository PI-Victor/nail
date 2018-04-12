package platform

// Platform represents the top level resource that
type Platform struct {
	Name             string       `json:"name"`
	PreviousRevision string       `json:"previousVersion"`
	CurrentRevision  string       `json:"currentVersion"`
	PreviousTag      string       `json:"previousTag"`
	CurrentTag       string       `json:"currentTag"`
	LockedRepos      []Repository `json:"LockedRepos,omitempty"`
}

// Repository is a repository at a specific revision that is currently locked by
// the top level `platform` resource.
type Repository struct {
	Name      string `json:"name"`
	Validated bool   `json:"validated"`
}

// Revision contains information about a repository revision.
type Revision struct {
	Name    string `json:"name,omitempty"`
	Message string `json:"message,omitempty"`
	Tag     string `json:"tag"`
}

func NewPlatform(name string) (platform *Platform) {
	return &Platform{
		Name:            name,
		CurrentTag:      "v0.0.0",
		CurrentRevision: "",
	}
}
