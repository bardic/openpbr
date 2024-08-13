package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

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
var TargetAssets = []string{"blocks", "entity", "particle", "items"}

var HeightMapNameSuffix = "_height"
var NormalMapNameSuffix = "_normal"
var MerMapNameSuffix = "_mer"

var LoadStdOut *widget.TextGrid

func CheckEnv() {
	fmt.Println("Check env")

	if _, err := exec.LookPath(IM_CMD); err != nil {
		if _, err := exec.LookPath("convert"); err != nil {
			fmt.Println("--- Fatal :: ImageMagick not found")
		} else {
			IM_CMD = "convert"
		}
	}

	if _, err := exec.LookPath("pngcrush"); err != nil {
		fmt.Println("--- Fatal :: PNGCrush not found")
	}

	if _, err := exec.LookPath("nvtt_export.exe"); err != nil {
		fmt.Println("--- Warning :: Nvidia Texture Tools not found")
	}
}

func LocalPath(partialPath string) string {
	return Basedir + string(os.PathSeparator) + partialPath
}

func CopyF(in string, out string) error {
	data, err := os.ReadFile(in)

	if err != nil {
		return err
	}

	// Write data to dst
	err = os.WriteFile(out, data, 0644)

	if err != nil {
		return err
	}

	return nil
}

func TgaPng(in string, out string) error {
	c := exec.Command(IM_CMD, in, "png32:"+out)
	c.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
	go c.Run()
	return nil
}

func PsdPng(in string, out string) error {
	c := exec.Command(IM_CMD, in+"[0]", "png32:"+out)
	c.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
	go c.Run()
	return nil
}

func AdjustColor(in string) error {
	c := exec.Command(IM_CMD, in, "-modulate", "95,105,105", "png32:"+in)
	c.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
	go c.Run()
	return nil
}

func CreateHeightMap(in string, out string) error {
	c := exec.Command(IM_CMD, in, "-channel", "RGB", "-negate", "-set", "colorspace", "Gray", "png32:"+out)
	c.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
	go c.Run()
	return nil
}

func Upscale(in string, out string) {
	c := exec.Command(IM_CMD, in, "-filter", "point", "-set", "option:distort:scale", "-distort", "SRT", "0", "-scale", "100%", "-unsharp", "12x6+0.5+0", "-type", "truecolor", "png32:"+out)
	c.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
	go c.Run()
}

func CreateNormalMap(in string, out string) error {
	c := exec.Command("nvtt_export.exe", in, "-p", "norm.dpf", "-o", out)
	c.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
	go c.Run()
	return nil
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
			c.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
			go c.Run()
			return
		}
	}
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
