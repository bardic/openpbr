package ui

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bardic/openpbr/cmd"
	"github.com/bardic/openpbr/store"
	"github.com/bardic/openpbr/utils"
	"github.com/bardic/openpbr/vo"
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

		for _, v := range store.Tabs {

			if v.TabName == "Config" || v.TabName == "Build Package" {
				continue
			}

			templateFile := path.Join("templates", filepath.Base(v.TemplatePath))

			output := path.Join(store.Output, "settings", v.Output)
			//err := os.MkdirAll(filepath.Base(fp), os.ModeDir)

			os.MkdirAll(path.Join(store.Output, utils.SettingDir, filepath.Dir(v.Output)), os.ModePerm)

			// if err != nil {
			// 	return
			// }

			out, err := os.Create(output)

			if err != nil {
				fmt.Println(err)
				return
			}
			defer out.Close()

			b, err := pack.templates.ReadFile(templateFile)

			if err != nil {
				fmt.Println(err)
				return
			}

			t, err := template.New("a").Parse(string(b))

			b, err = os.ReadFile(path.Join(utils.Basedir, strings.Split(v.TemplatePath, ".")[0]+".dat"))

			if err != nil {
				fmt.Println(err)
			}

			switch v.TabName {
			case "Atmospheric":
				var vo *vo.Atmospherics
				json.Unmarshal(b, &vo)
				err = t.Execute(out, vo)
			case "PBR":
				var vo *vo.PBR
				json.Unmarshal(b, &vo)
				err = t.Execute(out, vo)
			case "Fog":
				var vo *vo.Fog
				json.Unmarshal(b, &vo)
				err = t.Execute(out, vo)
			case "Lighting":
				var vo *vo.Lighting
				json.Unmarshal(b, &vo)
				err = t.Execute(out, vo)
			case "Color Grading":
				var vo *vo.ColorGrading
				json.Unmarshal(b, &vo)
				err = t.Execute(out, vo)
			case "Water":
				var vo *vo.Water
				json.Unmarshal(b, &vo)
				err = t.Execute(out, vo)
			}

			if err != nil {
				fmt.Println(err)
				return
			}
		}

		p.Canvas().Content().Refresh()
		err := (&cmd.Build{
			Templates:  pack.templates,
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

	tempDir := store.PackageStore

	for _, v := range dir {

		filePath := path.Join(tempDir, v.Name())
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
