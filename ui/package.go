package ui

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/bardic/openpbr/cmd"
	"github.com/bardic/openpbr/cmd/export"
	"github.com/bardic/openpbr/utils"
)

type Pack struct {
	templates embed.FS
	window    fyne.Window
}

func (p *Pack) BuildPackageView(refresh func(), popupSave func(*export.Manifest, error), popupErr func(error)) *fyne.Container {

	pb := widget.NewProgressBarInfinite()
	pb.Hide()
	utils.LoadStdOut = widget.NewRichText()
	utils.LoadStdOut.Scroll = widget.NewEntry().Scroll
	utils.LoadStdOut.Resize(fyne.NewSize(300, 600))

	loadBtn := widget.NewButton("Load Config", func() {
		dialog.ShowFileOpen(func(f fyne.URIReadCloser, err error) {
			if err != nil {
				popupErr(err)
				return
			}
			if f == nil {
				return
			}
			var saveFile = f.URI().Path()
			utils.Basedir = filepath.Dir(saveFile)
			pb.Show()
			dir, err := p.templates.ReadDir("templates")

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

				b, err := p.templates.ReadFile("templates/" + v.Name())

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

			refresh()
			err = (&cmd.Build{
				ConfigPath: saveFile,
			}).Perform()

			if err != nil {
				pb.Theme().Color(fyne.ThemeColorName("red"), fyne.ThemeVariant(1))
				return
			}

			pb.Theme().Color(fyne.ThemeColorName("green"), fyne.ThemeVariant(1))
			pb.Stop()
			refresh()

		}, p.window)
	})
	loadBtn.Resize(fyne.NewSize(25, 25))

	loadConfigContainer := container.NewBorder(loadBtn, pb, nil, nil,
		container.NewGridWithRows(1, utils.LoadStdOut))

	utils.LoadStdOut.Refresh()

	return loadConfigContainer
}
