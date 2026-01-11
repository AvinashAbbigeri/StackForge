package engine

type Plan struct {
	Installs []string
	Commands []string
	Files    []string
	Tests    []string
	Script   string
}
