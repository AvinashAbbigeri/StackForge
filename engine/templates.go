package engine

import (
	"io/fs"
	"os"
)

func ReadTemplate(path string) ([]byte, error) {
	if EmbeddedFS != nil {
		return fs.ReadFile(EmbeddedFS, path)
	}
	return os.ReadFile(path)
}
