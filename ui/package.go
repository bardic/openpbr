package ui

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bardic/openpbr/cmd"
	"github.com/bardic/openpbr/store"
	"github.com/bardic/openpbr/utils"
)

type Pack struct {
	templates embed.FS
	window    fyne.Window
}

func (pack *Pack) Save() {
}

func (pack *Pack) Build(p fyne.Window) *fyne.Container {

	pb := widget.NewProgressBarInfinite()
	pb.Hide()
	utils.LoadStdOut = widget.NewRichText()
	utils.LoadStdOut.Scroll = widget.NewEntry().Scroll
	utils.LoadStdOut.Resize(fyne.NewSize(300, 600))

	loadBtn := widget.NewButton("Load Config", func() {

		var saveFile = path.Join(store.PackageStore, "conf.json")
		utils.Basedir = filepath.Dir(saveFile)
		pb.Show()
		dir, err := pack.templates.ReadDir("templates")

		if err != nil {
			fmt.Println(err)
			return
		}

		tempDir := utils.LocalPath("templates")
		os.MkdirAll(tempDir, os.ModePerm)
		os.MkdirAll(utils.LocalPath(utils.OutDir), os.ModePerm)
		os.MkdirAll(utils.LocalPath(utils.Psds), os.ModePerm)
		os.MkdirAll(utils.LocalPath(utils.Overrides), os.ModePerm)
		os.MkdirAll(utils.LocalPath(utils.SettingDir), os.ModePerm)

		for _, v := range dir {

			filePath := tempDir + string(os.PathSeparator) + v.Name()
			out, err := os.Create(filePath)

			if err != nil {
				fmt.Println(err)
				return
			}
			defer out.Close()

			b, err := pack.templates.ReadFile("templates/" + v.Name())

			if err != nil {
				fmt.Println(err)
				return
			}

			_, err = io.Writer.Write(out, b)

			if err != nil {
				fmt.Println(err)
				return
			}
		}

		p.Canvas().Content().Refresh()
		err = (&cmd.Build{
			ConfigPath: saveFile,
		}).Perform()

		if err != nil {
			pb.Theme().Color(fyne.ThemeColorName("red"), fyne.ThemeVariant(1))
			return
		}

		pb.Theme().Color(fyne.ThemeColorName("green"), fyne.ThemeVariant(1))
		pb.Stop()
		p.Canvas().Content().Refresh()
	})
	loadBtn.Resize(fyne.NewSize(25, 25))

	loadConfigContainer := container.NewBorder(loadBtn, pb, nil, nil,
		container.NewGridWithRows(1, utils.LoadStdOut))

	utils.LoadStdOut.Refresh()

	return loadConfigContainer
}

func (p *Pack) Defaults(b []byte) {
	dir, err := p.templates.ReadDir("defaults")

	if err != nil {
		fmt.Println(err)
		return
	}

	tempDir := utils.LocalPath("defaults")
	os.MkdirAll(tempDir, os.ModePerm)

	for _, v := range dir {

		filePath := tempDir + string(os.PathSeparator) + v.Name()
		out, err := os.Create(filePath)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer out.Close()

		b, err := p.templates.ReadFile("defaults/" + v.Name())

		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = io.Writer.Write(out, b)

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
