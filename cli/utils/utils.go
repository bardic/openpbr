package utils

import (
	"os"
	"os/exec"
	"strings"

	cp "github.com/otiai10/copy"
)

const BaseAssets = "./input"
const BuildDir = "./output/openpbr"
const Overrides = "../overrides"
const ConfigDIr = "./settings"
const Temp = "./temp"

var TaretAssets = []string{"blocks", "entity", "particle", "items"}

func CheckForOverride(file string) (error, bool) {
	stringSlice := strings.Split(file, "/")
	items, _ := os.ReadDir(Overrides)
	for _, item := range items {
		if stringSlice[len(stringSlice)-1] == item.Name() {
			e := CopyF(Overrides+item.Name(), file)
			if e != nil {
				return e, false
			}
			return nil, true
		}
	}
	return nil, false
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

func CopyD(in string, out string) {
	cp.Copy(in, out)
}

func TgaPng(in string, out string) error {
	if err, b := CheckForOverride(out); err != nil || b {
		return err
	}

	out = strings.ReplaceAll(out, ".tga", ".jpg")
	c1 := exec.Command("convert", "-auto-orient", in, out)
	err := c1.Run()

	if err != nil {
		return err
	}

	pngOut := strings.ReplaceAll(out, ".jpg", ".png")
	c2 := exec.Command("convert", out, pngOut)
	err = c2.Run()

	if err != nil {
		return err
	}

	return nil
}

func PsdPng(in string, out string) {
	c := exec.Command("convert", in+"[0]", out)
	c.Run()
}

func AdjustColor(in string) error {
	if err, b := CheckForOverride(in); err != nil || b {
		return nil
	}

	c2 := exec.Command("convert", in, "-modulate", "101,99,99", in)
	e := c2.Run()

	if e != nil {
		return e
	}

	c1 := exec.Command("convert", in, "-colorspace", "sRGB", "-type", "truecolor", "png32:"+in)
	e = c1.Run()

	return e
}

func CreateHeightMap(in string, out string) {
	command := exec.Command("convert", in, "-set", "colorspace", "Gray", "-separate", "-average", "-channel", "RGB", out)
	command.Run()
}
