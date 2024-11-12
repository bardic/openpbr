package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/bardic/openpbr/cmd/export"
	"github.com/bardic/openpbr/vo"
)

type PBR struct {
}

func (v *PBR) BuildLightingView(refresh func(), popupErr func(error)) *fyne.Container {

	//
	// Global Block MERS
	//

	globalBlockMERSLabel := widget.NewLabel("Global Block MERS")
	globalBlockMERSEntry := widget.NewEntry()

	//
	// Global Actor MERS
	//

	globalActorMERSLabel := widget.NewLabel("Global Actor MERS")
	globalActorMERSEntry := widget.NewEntry()

	//
	// Global Particle MERS
	//

	globalParticleMERSLabel := widget.NewLabel("Global Particle MERS")
	globalParticleMERSEntry := widget.NewEntry()

	//
	// Global Item MERS
	//

	globalItemMERSLabel := widget.NewLabel("Global Item MERS")
	globalItemMERSEntry := widget.NewEntry()

	save := widget.NewButton("Save", func() {
		// Save the lighting settings
		pbr := export.GlobalPBR{
			Out: "./openpbr_out/pbr/global.json",
			PBR: vo.PBR{
				GlobalBlockMERS:    globalBlockMERSEntry.Text,
				GlobalActorMERS:    globalActorMERSEntry.Text,
				GlobalParticleMERS: globalParticleMERSEntry.Text,
				GlobalItemMERS:     globalItemMERSEntry.Text,
			},
		}

		pbr.Perform()
	})

	c := container.New(layout.NewFormLayout(), globalBlockMERSLabel, globalBlockMERSEntry, globalActorMERSLabel, globalActorMERSEntry, globalParticleMERSLabel, globalParticleMERSEntry, globalItemMERSLabel, globalItemMERSEntry, save, layout.NewSpacer())
	return c
}
