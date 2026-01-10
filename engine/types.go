package engine

// Module represents one installable component (python, flask, git, etc)
type Module struct {
	ID        string            `json:"id"`
	Requires  []string          `json:"requires"`
	Conflicts []string          `json:"conflicts"`
	Install   []string          `json:"install"`
	Files     map[string]string `json:"files"`
	Test      []string          `json:"test"`
}

// ProjectConfig represents what the user asked for
type ProjectConfig struct {
	Modules []string `json:"modules"`
	Name    string   `json:"name"`
}
