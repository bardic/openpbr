package utils

import (
	"errors"
	"os"
	"strings"

	"fyne.io/fyne/v2/widget"
)

// Folders
const BaseAssets = "input"
const OutDir = "openpbr_out"
const Overrides = "overrides"
const SettingDir = "settings"
const Psds = "psds"

var IM_CMD = "magick"
var Beta bool
var DeleteAutoGen bool
var SkipDownload bool
var NormalMaps bool
var ZipOnly bool
var Crush bool
var TexturesetVersion string
var Basedir string
var Failed bool
var TargetAssets = []string{"blocks", "entity", "particle"}

var HeightMapNameSuffix = "_height"
var NormalMapNameSuffix = "_normal"
var MerMapNameSuffix = "_mer"

var LoadStdOut *widget.TextGrid

func LocalPath(partialPath string) string {
	return Basedir + string(os.PathSeparator) + partialPath
}

func AppendLoadOut(s string) {
	LoadStdOut.SetText(LoadStdOut.Text() + "\n" + s)
}

func GetTextureSubpath(p string, key string) (string, error) {
	subpaths := strings.Split(p, string(os.PathSeparator))
	for i, subpath := range subpaths {
		if subpath == key {
			sub := strings.Join(subpaths[i:], string(os.PathSeparator))
			return sub, nil
		}
	}

	return "", errors.New("")
}
