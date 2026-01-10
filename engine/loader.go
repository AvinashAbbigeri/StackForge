package engine

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func LoadModules(root string) (map[string]Module, error) {
	modules := make(map[string]Module)

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip folders
		if d.IsDir() {
			return nil
		}

		// Only JSON files
		if filepath.Ext(path) != ".json" {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		var m Module
		if err := json.Unmarshal(data, &m); err != nil {
			return fmt.Errorf("error in %s: %w", path, err)
		}

		if m.ID == "" {
			return fmt.Errorf("module %s has no id", path)
		}

		if _, exists := modules[m.ID]; exists {
			return fmt.Errorf("duplicate module id: %s", m.ID)
		}

		modules[m.ID] = m
		return nil
	})

	if err != nil {
		return nil, err
	}

	return modules, nil
}
