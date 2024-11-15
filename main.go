package main

import (
	"embed"

	"github.com/bardic/openpbr/ui"
)

//go:embed templates/*.tmpl
var templates embed.FS

//go:embed defaults/*.json
var defaults embed.FS

func main() {
	ui := ui.UI{}
	ui.Build(templates, defaults)
}
