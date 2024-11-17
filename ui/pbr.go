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
	globalBlockMERSEntry    *vo.RGBA
	globalActorMERSEntry    *vo.RGBA
	globalParticleMERSEntry *vo.RGBA
	globalItemMERSEntry     *vo.RGBA
}

func (v *PBR) BuildLightingView(refresh func(), popupErr func(error)) *fyne.Container {

	//
	// Global Block MERS
	//

	globalBlockMERSLabel := widget.NewLabel("Global Block MERS")
	v.globalBlockMERSEntry = utils.CreateRGBAEntry()

	//
	// Global Actor MERS
	//

	globalActorMERSLabel := widget.NewLabel("Global Actor MERS")
	v.globalActorMERSEntry = utils.CreateRGBAEntry()

	//
	// Global Particle MERS
	//

	globalParticleMERSLabel := widget.NewLabel("Global Particle MERS")
	v.globalParticleMERSEntry = utils.CreateRGBAEntry()

	//
	// Global Item MERS
	//

	globalItemMERSLabel := widget.NewLabel("Global Item MERS")
	v.globalItemMERSEntry = utils.CreateRGBAEntry()

	save := widget.NewButton("Save", func() {
		pbr := export.GlobalPBR{
			Out: "./example/settings/pbr/pbr/global.json",
			PBR: vo.PBR{
				GlobalBlockR:    utils.ToFloat64(v.globalBlockMERSEntry.R),
				GlobalBlockG:    utils.ToFloat64(v.globalBlockMERSEntry.G),
				GlobalBlockB:    utils.ToFloat64(v.globalBlockMERSEntry.B),
				GlobalBlockA:    utils.ToFloat64(v.globalBlockMERSEntry.A),
				GlobalActorR:    utils.ToFloat64(v.globalActorMERSEntry.R),
				GlobalActorG:    utils.ToFloat64(v.globalActorMERSEntry.G),
				GlobalActorB:    utils.ToFloat64(v.globalActorMERSEntry.B),
				GlobalActorA:    utils.ToFloat64(v.globalActorMERSEntry.A),
				GlobalParticleR: utils.ToFloat64(v.globalParticleMERSEntry.R),
				GlobalParticleG: utils.ToFloat64(v.globalParticleMERSEntry.G),
				GlobalParticleB: utils.ToFloat64(v.globalParticleMERSEntry.B),
				GlobalParticleA: utils.ToFloat64(v.globalParticleMERSEntry.A),
				GlobalItemR:     utils.ToFloat64(v.globalItemMERSEntry.R),
				GlobalItemG:     utils.ToFloat64(v.globalItemMERSEntry.G),
				GlobalItemB:     utils.ToFloat64(v.globalItemMERSEntry.B),
				GlobalItemA:     utils.ToFloat64(v.globalItemMERSEntry.A),
			},
		}

		pbr.Perform()
	})

	c := container.New(
		layout.NewFormLayout(),
		globalBlockMERSLabel,
		v.globalBlockMERSEntry.C,
		globalActorMERSLabel,
		v.globalActorMERSEntry.C,
		globalParticleMERSLabel,
		v.globalParticleMERSEntry.C,
		globalItemMERSLabel,
		v.globalItemMERSEntry.C,
		save,
		layout.NewSpacer())
	return c
}

func (v *PBR) Defaults(vo *vo.PBR) {
	v.globalBlockMERSEntry.R.Text = utils.FloatToString(vo.GlobalBlockR)
	v.globalBlockMERSEntry.G.Text = utils.FloatToString(vo.GlobalBlockG)
	v.globalBlockMERSEntry.B.Text = utils.FloatToString(vo.GlobalBlockB)
	v.globalBlockMERSEntry.A.Text = utils.FloatToString(vo.GlobalBlockA)
	v.globalActorMERSEntry.R.Text = utils.FloatToString(vo.GlobalActorR)
	v.globalActorMERSEntry.G.Text = utils.FloatToString(vo.GlobalActorG)
	v.globalActorMERSEntry.B.Text = utils.FloatToString(vo.GlobalActorB)
	v.globalActorMERSEntry.A.Text = utils.FloatToString(vo.GlobalActorA)
	v.globalParticleMERSEntry.R.Text = utils.FloatToString(vo.GlobalParticleR)
	v.globalParticleMERSEntry.G.Text = utils.FloatToString(vo.GlobalParticleG)
	v.globalParticleMERSEntry.B.Text = utils.FloatToString(vo.GlobalParticleB)
	v.globalParticleMERSEntry.A.Text = utils.FloatToString(vo.GlobalParticleA)
	v.globalItemMERSEntry.R.Text = utils.FloatToString(vo.GlobalItemR)
	v.globalItemMERSEntry.G.Text = utils.FloatToString(vo.GlobalItemG)
	v.globalItemMERSEntry.B.Text = utils.FloatToString(vo.GlobalItemB)
	v.globalItemMERSEntry.A.Text = utils.FloatToString(vo.GlobalItemA)
}

func (a *PBR) Save() {
}
