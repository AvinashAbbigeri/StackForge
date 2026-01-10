package engine

import "fmt"

func Resolve(selected []string, all map[string]Module) ([]Module, error) {
	visited := make(map[string]bool)
	resolving := make(map[string]bool)
	var ordered []Module

	var visit func(string) error
	visit = func(id string) error {
		if resolving[id] {
			return fmt.Errorf("circular dependency at %s", id)
		}
		if visited[id] {
			return nil
		}

		m, ok := all[id]
		if !ok {
			return fmt.Errorf("unknown module: %s", id)
		}

		resolving[id] = true

		for _, dep := range m.Requires {
			if err := visit(dep); err != nil {
				return err
			}
		}

		resolving[id] = false
		visited[id] = true
		ordered = append(ordered, m)
		return nil
	}

	for _, id := range selected {
		if err := visit(id); err != nil {
			return nil, err
		}
	}

	return ordered, nil
}
