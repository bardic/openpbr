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
	RGBs map[string]vo.RGB
}

func (v *Fog) BuildLightingView(refresh func(), popupErr func(error)) *fyne.Container {
	//
	// Water Max Density
	//

	waterMaxDensityLabel := widget.NewLabel("Water Max Density")
	waterMaxDensityEntry := widget.NewEntry()

	//
	// Water Uniform Density
	//

	waterUniformDensityLabel := widget.NewLabel("Water Uniform Density")
	waterUniformDensityEntry := widget.NewCheck("", func(bool) {})

	//
	// Air Max Density
	//

	airMaxDensityLabel := widget.NewLabel("Air Max Density")
	airMaxDensityEntry := widget.NewEntry()

	//
	// Air Zero Density Height
	//

	airZeroDensityHeightLabel := widget.NewLabel("Air Zero Density Height")
	airZeroDensityHeightEntry := widget.NewEntry()

	//
	// Air Max Density Height
	//

	airMaxDensityHeightLabel := widget.NewLabel("Air Max Density Height")
	airMaxDensityHeightEntry := widget.NewEntry()

	//
	// Water Scattering

	waterScatteringRGBLabel := widget.NewLabel("Water Scattering")
	waterScatteringRGB := utils.CreateRGBEntry()

	//
	// Water Absorption

	waterAbsorptionRGBLabel := widget.NewLabel("Water Absorption")
	waterAbsorptionRGB := utils.CreateRGBEntry()

	//
	// Air Scattering
	//

	airScatteringRGBLabel := widget.NewLabel("Water Scattering")
	airScatteringRGB := utils.CreateRGBEntry()

	//
	// Air Absorption
	//

	airAbsorptionRGBLabel := widget.NewLabel("Water Scattering")
	airAbsorptionRGB := utils.CreateRGBEntry()

	save := widget.NewButton("Save", func() {
		fog := export.Fog{
			Out: "./openpbr_out/fogs/default_fog_settings.json",
			Fog: vo.Fog{
				WaterMaxDensity:      utils.ToFloat64(waterMaxDensityEntry),
				WaterUniformDensity:  waterUniformDensityEntry.Checked,
				AirMaxDensity:        utils.ToFloat64(airMaxDensityEntry),
				AirZeroDensityHeight: utils.ToFloat64(airZeroDensityHeightEntry),
				AirMaxDensityHeight:  utils.ToFloat64(airMaxDensityHeightEntry),
				WaterScatteringR:     utils.ToFloat64(waterScatteringRGB.R),
				WaterScatteringG:     utils.ToFloat64(waterScatteringRGB.G),
				WaterScatteringB:     utils.ToFloat64(waterScatteringRGB.B),
				WaterAbsorptionR:     utils.ToFloat64(waterAbsorptionRGB.R),
				WaterAbsorptionG:     utils.ToFloat64(waterAbsorptionRGB.G),
				WaterAbsorptionB:     utils.ToFloat64(waterAbsorptionRGB.B),
				AirScatteringR:       utils.ToFloat64(airScatteringRGB.R),
				AirScatteringG:       utils.ToFloat64(airScatteringRGB.G),
				AirScatteringB:       utils.ToFloat64(airScatteringRGB.B),
				AirAbsorptionR:       utils.ToFloat64(airAbsorptionRGB.R),
				AirAbsorptionG:       utils.ToFloat64(airAbsorptionRGB.G),
				AirAbsorptionB:       utils.ToFloat64(airAbsorptionRGB.B),
			},
		}

		fog.Perform()
	})

	c := container.New(
		layout.NewFormLayout(),
		waterMaxDensityLabel, waterMaxDensityEntry,
		waterUniformDensityLabel, waterUniformDensityEntry,
		airMaxDensityLabel, airMaxDensityEntry,
		airZeroDensityHeightLabel, airZeroDensityHeightEntry,
		airMaxDensityHeightLabel, airMaxDensityHeightEntry,
		waterScatteringRGBLabel, waterScatteringRGB.C,
		waterAbsorptionRGBLabel, waterAbsorptionRGB.C,
		airScatteringRGBLabel, airScatteringRGB.C,
		airAbsorptionRGBLabel, airAbsorptionRGB.C,
		save, layout.NewSpacer(),
	)
	return c
}
