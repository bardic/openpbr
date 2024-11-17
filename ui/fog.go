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

type Fog struct {
	// RGBs map[string]vo.RGB

	waterMaxDensityEntry      *widget.Entry
	waterUniformDensityEntry  *widget.Check
	airMaxDensityEntry        *widget.Entry
	airZeroDensityHeightEntry *widget.Entry
	airMaxDensityHeightEntry  *widget.Entry
	waterScatteringRGB        *vo.RGB
	waterAbsorptionRGB        *vo.RGB
	airScatteringRGB          *vo.RGB
	airAbsorptionRGB          *vo.RGB
}

func (v *Fog) BuildLightingView(refresh func(), popupErr func(error)) *fyne.Container {
	//
	// Water Max Density
	//

	waterMaxDensityLabel := widget.NewLabel("Water Max Density")
	v.waterMaxDensityEntry = widget.NewEntry()

	//
	// Water Uniform Density
	//

	waterUniformDensityLabel := widget.NewLabel("Water Uniform Density")
	v.waterUniformDensityEntry = widget.NewCheck("", func(bool) {})

	//
	// Air Max Density
	//

	airMaxDensityLabel := widget.NewLabel("Air Max Density")
	v.airMaxDensityEntry = widget.NewEntry()

	//
	// Air Zero Density Height
	//

	airZeroDensityHeightLabel := widget.NewLabel("Air Zero Density Height")
	v.airZeroDensityHeightEntry = widget.NewEntry()

	//
	// Air Max Density Height
	//

	airMaxDensityHeightLabel := widget.NewLabel("Air Max Density Height")
	v.airMaxDensityHeightEntry = widget.NewEntry()

	//
	// Water Scattering

	waterScatteringRGBLabel := widget.NewLabel("Water Scattering")
	v.waterScatteringRGB = utils.CreateRGBEntry()

	//
	// Water Absorption

	waterAbsorptionRGBLabel := widget.NewLabel("Water Absorption")
	v.waterAbsorptionRGB = utils.CreateRGBEntry()

	//
	// Air Scattering
	//

	airScatteringRGBLabel := widget.NewLabel("Water Scattering")
	v.airScatteringRGB = utils.CreateRGBEntry()

	//
	// Air Absorption
	//

	airAbsorptionRGBLabel := widget.NewLabel("Water Scattering")
	v.airAbsorptionRGB = utils.CreateRGBEntry()

	save := widget.NewButton("Save", func() {
		fog := export.Fog{
			Out: "./example/settings/shared/fogs/default_fog_settings.json",
			Fog: vo.Fog{
				WaterMaxDensity:      utils.ToFloat64(v.waterMaxDensityEntry),
				WaterUniformDensity:  v.waterUniformDensityEntry.Checked,
				AirMaxDensity:        utils.ToFloat64(v.airMaxDensityEntry),
				AirZeroDensityHeight: utils.ToFloat64(v.airZeroDensityHeightEntry),
				AirMaxDensityHeight:  utils.ToFloat64(v.airMaxDensityHeightEntry),
				WaterScatteringR:     utils.ToFloat64(v.waterScatteringRGB.R),
				WaterScatteringG:     utils.ToFloat64(v.waterScatteringRGB.G),
				WaterScatteringB:     utils.ToFloat64(v.waterScatteringRGB.B),
				WaterAbsorptionR:     utils.ToFloat64(v.waterAbsorptionRGB.R),
				WaterAbsorptionG:     utils.ToFloat64(v.waterAbsorptionRGB.G),
				WaterAbsorptionB:     utils.ToFloat64(v.waterAbsorptionRGB.B),
				AirScatteringR:       utils.ToFloat64(v.airScatteringRGB.R),
				AirScatteringG:       utils.ToFloat64(v.airScatteringRGB.G),
				AirScatteringB:       utils.ToFloat64(v.airScatteringRGB.B),
				AirAbsorptionR:       utils.ToFloat64(v.airAbsorptionRGB.R),
				AirAbsorptionG:       utils.ToFloat64(v.airAbsorptionRGB.G),
				AirAbsorptionB:       utils.ToFloat64(v.airAbsorptionRGB.B),
			},
		}

		fog.Perform()
	})

	c := container.New(
		layout.NewFormLayout(),
		waterMaxDensityLabel, v.waterMaxDensityEntry,
		waterUniformDensityLabel, v.waterUniformDensityEntry,
		airMaxDensityLabel, v.airMaxDensityEntry,
		airZeroDensityHeightLabel, v.airZeroDensityHeightEntry,
		airMaxDensityHeightLabel, v.airMaxDensityHeightEntry,
		waterScatteringRGBLabel, v.waterScatteringRGB.C,
		waterAbsorptionRGBLabel, v.waterAbsorptionRGB.C,
		airScatteringRGBLabel, v.airScatteringRGB.C,
		airAbsorptionRGBLabel, v.airAbsorptionRGB.C,
		save, layout.NewSpacer(),
	)
	return c
}

func (v *Fog) Defaults(vo *vo.Fog) {
	v.waterMaxDensityEntry.Text = utils.FloatToString(vo.WaterMaxDensity)
	v.waterUniformDensityEntry.Checked = vo.WaterUniformDensity
	v.airMaxDensityEntry.Text = utils.FloatToString(vo.AirMaxDensity)
	v.airZeroDensityHeightEntry.Text = utils.FloatToString(vo.AirZeroDensityHeight)
	v.airMaxDensityHeightEntry.Text = utils.FloatToString(vo.AirMaxDensityHeight)
	v.waterScatteringRGB.R.Text = utils.FloatToString(vo.WaterScatteringR)
	v.waterScatteringRGB.G.Text = utils.FloatToString(vo.WaterScatteringG)
	v.waterScatteringRGB.B.Text = utils.FloatToString(vo.WaterScatteringB)
	v.waterAbsorptionRGB.R.Text = utils.FloatToString(vo.WaterAbsorptionR)
	v.waterAbsorptionRGB.G.Text = utils.FloatToString(vo.WaterAbsorptionG)
	v.waterAbsorptionRGB.B.Text = utils.FloatToString(vo.WaterAbsorptionB)
	v.airScatteringRGB.R.Text = utils.FloatToString(vo.AirScatteringR)
	v.airScatteringRGB.G.Text = utils.FloatToString(vo.AirScatteringG)
	v.airScatteringRGB.B.Text = utils.FloatToString(vo.AirScatteringB)
	v.airAbsorptionRGB.R.Text = utils.FloatToString(vo.AirAbsorptionR)
	v.airAbsorptionRGB.G.Text = utils.FloatToString(vo.AirAbsorptionG)
	v.airAbsorptionRGB.B.Text = utils.FloatToString(vo.AirAbsorptionB)
}
