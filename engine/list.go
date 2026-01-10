package engine

import "sort"

func ListModules(mods map[string]Module) []string {
	var names []string
	for k := range mods {
		names = append(names, k)
	}
	sort.Strings(names)
	return names
}
