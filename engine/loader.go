package engine

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func LoadModules(dir string) (map[string]Module, error) {
	modules := make(map[string]Module)

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".json" {
			continue
		}

		data, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}

		var m Module
		if err := json.Unmarshal(data, &m); err != nil {
			return nil, err
		}

		if m.ID == "" {
			return nil, fmt.Errorf("module %s has no id", file.Name())
		}

		modules[m.ID] = m
	}

	return modules, nil
}
