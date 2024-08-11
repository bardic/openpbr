package utils

import (
	"os"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2/widget"
	cp "github.com/otiai10/copy"
)

// Folders
const BaseAssets = "input"
const OutDir = "openpbr"
const Overrides = "overrides"
const SettingDir = "settings"
const IM_CMD = "magick"
const Psds = "psds"

var Beta bool
var DeleteAutoGen bool
var SkipDownload bool
var NormalMaps bool
var ZipOnly bool
var Crush bool
var TexturesetVersion string
var Basedir string

var TargetAssets = []string{"blocks", "entity", "particle", "items"}

var LoadStdOut *widget.TextGrid

func LocalPath(partialPath string) string {
	return Basedir + string(os.PathSeparator) + partialPath
}

func CopyF(in string, out string) error {
	data, err := os.ReadFile(in)

	if err != nil {
		return err
	}

	err = os.WriteFile(out, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func CopyD(in string, out string) error {
	return cp.Copy(in, out)
}

func TgaPng(in string, out string) error {
	c := exec.Command(IM_CMD, in, "png32:"+out)
	return c.Run()
}

func PsdPng(in string, out string) error {
	c := exec.Command(IM_CMD, in+"[0]", "png32:"+out)
	return c.Run()
}

func AdjustColor(in string) error {
	c := exec.Command(IM_CMD, in, "-modulate", "95,105,105", "png32:"+in)
	e := c.Run()
	return e
}

func CreateHeightMap(in string, out string) error {
	c := exec.Command(IM_CMD, in, "-set", "colorspace", "Gray", "-negate", "-channel", "RGB", "png32:"+out)
	return c.Run()
}

func Upscale(in string, out string) {
	c := exec.Command(IM_CMD, in, "-filter", "point", "-set", "option:distort:scale", "-distort", "SRT", "0", "-scale", "100%", "-unsharp", "12x6+0.5+0", "-type", "truecolor", "png32:"+out)
	c.Run()
}

func CreateNormalMap(in string, out string) error {
	c := exec.Command("nvtt_export.exe", in, "-p", "norm.dpf", "-o", out)
	return c.Run()
}

func CrushFiles(out string) {
	items, _ := os.ReadDir(out)
	for _, item := range items {
		if item.IsDir() {
			CrushFiles(out + "/" + item.Name())
			continue
		}

		if strings.Contains(item.Name(), ".png") {
			c := exec.Command("pngcrush", "-brute", "-ow", item.Name())
			c.Run()
		}
	}
}

func AppendLoadOut(s string) {
	LoadStdOut.SetText(LoadStdOut.Text() + "\n" + s)
}
