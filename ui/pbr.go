package ui

import (
	"encoding/json"
	"path"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/bardic/openpbr/cmd/export"
	"github.com/bardic/openpbr/store"
	"github.com/bardic/openpbr/utils"
	"github.com/bardic/openpbr/vo"
)

type PBR struct {
	globalBlockMERSEntry    *vo.RGBA
	globalActorMERSEntry    *vo.RGBA
	globalParticleMERSEntry *vo.RGBA
	globalItemMERSEntry     *vo.RGBA
}

func (v *PBR) Build(p fyne.Window) *fyne.Container {

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
	)
	return c
}

func (v *PBR) Defaults(b []byte) {

	if b == nil {
		v.globalBlockMERSEntry.R.Text = ""
		v.globalBlockMERSEntry.G.Text = ""
		v.globalBlockMERSEntry.B.Text = ""
		v.globalBlockMERSEntry.A.Text = ""
		v.globalActorMERSEntry.R.Text = ""
		v.globalActorMERSEntry.G.Text = ""
		v.globalActorMERSEntry.B.Text = ""
		v.globalActorMERSEntry.A.Text = ""
		v.globalParticleMERSEntry.R.Text = ""
		v.globalParticleMERSEntry.G.Text = ""
		v.globalParticleMERSEntry.B.Text = ""
		v.globalParticleMERSEntry.A.Text = ""
		v.globalItemMERSEntry.R.Text = ""
		v.globalItemMERSEntry.G.Text = ""
		v.globalItemMERSEntry.B.Text = ""
		v.globalItemMERSEntry.A.Text = ""
		return
	}

	var vo vo.PBR
	json.Unmarshal(b, &vo)

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

func (v *PBR) Save() {
	pbr := export.GlobalPBR{
		PBR: vo.PBR{
			BaseConf: vo.BaseConf{
				Out: path.Join(store.PackageStore, "pbr_global.json"),
			},
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
}
