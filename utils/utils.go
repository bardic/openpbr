package utils

import (
	"errors"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"fyne.io/fyne/v2/widget"
)

const BaseAssets = "input"
const OutDir = "openpbr_out"
const Overrides = "overrides"
const SettingDir = "settings"
const Psds = "psds"

var ImCmd = "magick"

var Basedir string
var TargetAssets = []string{"blocks", "entity", "particle"}

var HeightMapNameSuffix = "_height"
var MerMapNameSuffix = "_mer"

var LoadStdOut *widget.RichText

func LocalPath(partialPath string) string {
	return Basedir + string(os.PathSeparator) + partialPath
}

func RunCmd(cmd *exec.Cmd) error {
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
	err := cmd.Start()
	return err
}

func StartUpCheck() error {
	if _, err := exec.LookPath(ImCmd); err != nil {
		return errors.New("imagemagick not found")
	}

	return nil
}

func AppendLoadOut(s string) {
	LoadStdOut.AppendMarkdown(s)
	LoadStdOut.Refresh()
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
