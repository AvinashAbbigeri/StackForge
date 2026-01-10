package engine

import (
	"encoding/json"
	"os"
)

func LoadPresets(path string) (map[string][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var presets map[string][]string
	if err := json.Unmarshal(data, &presets); err != nil {
		return nil, err
	}

	return presets, nil
}

func PresetNames(p map[string][]string) []string {
	var names []string
	for k := range p {
		names = append(names, k)
	}
	return names
}
