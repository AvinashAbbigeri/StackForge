package main

import "embed"

//go:embed modules templates presets.json
var embeddedFS embed.FS
