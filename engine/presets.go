package engine

import (
	"encoding/json"
	"io/fs"
	"os"
)

func LoadPresets(path string) (map[string][]string, error) {
	var data []byte
	var err error

	if EmbeddedFS != nil {
		data, err = fs.ReadFile(EmbeddedFS, "presets.json")
	} else {
		data, err = os.ReadFile(path)
	}

	if err != nil {
		return nil, err
	}

	var presets map[string][]string
	err = json.Unmarshal(data, &presets)
	return presets, err
}

func PresetNames(p map[string][]string) []string {
	var names []string
	for k := range p {
		names = append(names, k)
	}
	return names
}
