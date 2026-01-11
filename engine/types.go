package engine

// Module represents one installable component (python, flask, git, etc)
type Module struct {
	ID        string   `json:"id"`
	Requires  []string `json:"requires"`
	Conflicts []string `json:"conflicts"`

	Install  map[string][]string `json:"install,omitempty"`
	Commands []string            `json:"commands,omitempty"`

	Files map[string]string `json:"files,omitempty"`
	Tests []string          `json:"test,omitempty"`
}

// ProjectConfig represents what the user asked for
type ProjectConfig struct {
	Modules []string `json:"modules"`
	Name    string   `json:"name"`
}
