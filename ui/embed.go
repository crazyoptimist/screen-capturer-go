package ui

import (
	"embed"
	"io/fs"
)

//go:embed dist
var EmbeddedFS embed.FS

func Root() (fs.FS, error) {
	return fs.Sub(EmbeddedFS, "dist")
}
