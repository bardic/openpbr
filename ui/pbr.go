package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/bardic/openpbr/cmd/export"
	"github.com/bardic/openpbr/utils"
	"github.com/bardic/openpbr/vo"
)

type PBR struct {
}

func (v *PBR) BuildLightingView(refresh func(), popupErr func(error)) *fyne.Container {

	//
	// Global Block MERS
	//

	globalBlockMERSLabel := widget.NewLabel("Global Block MERS")
	globalBlockMERSEntry := utils.CreateRGBAEntry()

	//
	// Global Actor MERS
	//

	globalActorMERSLabel := widget.NewLabel("Global Actor MERS")
	globalActorMERSEntry := utils.CreateRGBAEntry()

	//
	// Global Particle MERS
	//

	globalParticleMERSLabel := widget.NewLabel("Global Particle MERS")
	globalParticleMERSEntry := utils.CreateRGBAEntry()

	//
	// Global Item MERS
	//

	globalItemMERSLabel := widget.NewLabel("Global Item MERS")
	globalItemMERSEntry := utils.CreateRGBAEntry()

	save := widget.NewButton("Save", func() {
		// Save the lighting settings
		pbr := export.GlobalPBR{
			Out: "./example/settings/pbr/pbr/global.json",
			PBR: vo.PBR{
				GlobalBlockMERS:    "1",
				GlobalActorMERS:    "1",
				GlobalParticleMERS: "1",
				GlobalItemMERS:     "1",
			},
		}

		pbr.Perform()
	})

	c := container.New(layout.NewFormLayout(), globalBlockMERSLabel, globalBlockMERSEntry.C, globalActorMERSLabel, globalActorMERSEntry.C, globalParticleMERSLabel, globalParticleMERSEntry.C, globalItemMERSLabel, globalItemMERSEntry.C, save, layout.NewSpacer())
	return c
}
