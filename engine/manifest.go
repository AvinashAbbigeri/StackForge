package engine

import (
	"encoding/json"
	"os"
)

type Manifest struct {
	Modules []string `json:"modules"`
	OS      string   `json:"os"`
}

func LoadManifest(path string) (*Manifest, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var m Manifest
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	return &m, nil
}

func SaveManifest(path string, m *Manifest) error {
	data, _ := json.MarshalIndent(m, "", "  ")
	return os.WriteFile(path, data, 0644)
}
