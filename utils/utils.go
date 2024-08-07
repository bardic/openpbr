package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	cp "github.com/otiai10/copy"
)

const BaseAssets = "./input"
const BuildDir = "./output/openpbr"
const Overrides = "./overrides"
const SettingDIr = "./settings"
const IM_CMD = "magick"
const Psds = "./psds"

var Beta bool
var DeleteAutoGen bool
var SkipDownload bool
var NormalMaps bool
var ZipOnly bool
var Crush bool
var TexturesetVersion string

var TargetAssets = []string{"blocks", "entity", "particle", "items"}

func CheckForOverride(file string) (bool, error) {
	stringSlice := strings.Split(file, string(os.PathSeparator))
	items, _ := os.ReadDir(Overrides)
	for _, item := range items {
		if stringSlice[len(stringSlice)-1] == item.Name() {
			p := filepath.Join(Overrides, item.Name())

			e := CopyF(p, file)
			if e != nil {
				return false, e
			}
			return true, nil
		}
	}
	return false, nil
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
	c1 := exec.Command(IM_CMD, in, out)
	err := c1.Run()

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func PsdPng(in string, out string) error {
	c := exec.Command(IM_CMD, in+"[0]", out)
	return c.Run()
}

func AdjustColor(in string) error {
	if b, err := CheckForOverride(in); err != nil || b {
		return nil
	}

	c2 := exec.Command(IM_CMD, in, "-modulate", "101,99,99", in)
	e := c2.Run()

	if e != nil {
		return e
	}

	c1 := exec.Command(IM_CMD, in, "-colorspace", "sRGB", "-type", "truecolor", "png32:"+in)
	e = c1.Run()

	return e
}

func CreateHeightMap(in string, out string) error {
	if b, err := CheckForOverride(in); err != nil || b {
		return nil
	}
	command := exec.Command(IM_CMD, in, "-set", "colorspace", "Gray", "-negate", "-channel", "RGB", out)
	return command.Run()
}

func Upscale(in string, out string) {
	c := exec.Command(IM_CMD, in, "-filter", "point", "-set", "option:distort:scale", "-distort", "SRT", "0", "-scale", "100%", "-unsharp", "12x6+0.5+0", "-type", "truecolor", "png32:"+out)
	c.Run()
}

func CreateNormalMap(in string, out string) error {
	if b, err := CheckForOverride(in); err != nil || b {
		return nil
	}
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
