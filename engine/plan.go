package engine

type Plan struct {
	Installs []string
	Files    []string
	Tests    []string
	Script   string
}
