package engine

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func LoadModules(root string) (map[string]Module, error) {
	modules := make(map[string]Module)

	var fsys fs.FS = os.DirFS(root)
	if EmbeddedFS != nil {
		fsys = EmbeddedFS
	}

	err := fs.WalkDir(fsys, "modules", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}

		data, err := fs.ReadFile(fsys, path)
		if err != nil {
			return err
		}

		var m Module
		if err := json.Unmarshal(data, &m); err != nil {
			return fmt.Errorf("error in %s: %w", path, err)
		}

		modules[m.ID] = m
		return nil
	})

	return modules, err
}
