package engine

import "io/fs"

var EmbeddedFS fs.FS

func SetEmbeddedFS(fsys fs.FS) {
	EmbeddedFS = fsys
}
