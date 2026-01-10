package engine

func Unique(list []string) []string {
	seen := make(map[string]bool)
	var out []string
	for _, v := range list {
		if !seen[v] {
			seen[v] = true
			out = append(out, v)
		}
	}
	return out
}
