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

type Water struct {
}

func (v *Water) BuildLightingView(refresh func(), popupErr func(error)) *fyne.Container {
	//
	// Chlorophyll
	//

	chlorophyllLabel := widget.NewLabel("Chlorophyll")
	chlorophyllEntry := widget.NewEntry()

	//
	// Suspended Sediment
	//

	suspendedSedimentLabel := widget.NewLabel("Suspended Sediment")
	suspendedSedimentEntry := widget.NewEntry()

	//
	// CDOM
	//

	cdomLabel := widget.NewLabel("CDOM")
	cdomEntry := widget.NewEntry()

	//
	// Waves Enabled
	//

	wavesEnabledLabel := widget.NewLabel("Waves Enabled")
	wavesEnabledEntry := widget.NewCheck("Waves Enabled", func(bool) {})

	//
	// Waves Depth
	//

	wavesDepthLabel := widget.NewLabel("Waves Depth")
	wavesDepthEntry := widget.NewEntry()

	//
	// Waves Frequency
	//

	wavesFrequencyLabel := widget.NewLabel("Waves Frequency")
	wavesFrequencyEntry := widget.NewEntry()

	//
	// Waves Frequency Scaling
	//

	wavesFrequencyScalingLabel := widget.NewLabel("Waves Frequency Scaling")
	wavesFrequencyScalingEntry := widget.NewEntry()

	//
	// Waves Mix
	//

	wavesMixLabel := widget.NewLabel("Waves Mix")
	wavesMixEntry := widget.NewEntry()

	//
	// Waves Octaves
	//

	wavesOctavesLabel := widget.NewLabel("Waves Octaves")
	wavesOctavesEntry := widget.NewEntry()

	//
	// Waves Pull
	//

	wavesPullLabel := widget.NewLabel("Waves Pull")
	wavesPullEntry := widget.NewEntry()

	//
	// Waves Sample Width
	//

	wavesSampleWidthLabel := widget.NewLabel("Waves Sample Width")
	wavesSampleWidthEntry := widget.NewEntry()

	//
	// Waves Shape
	//

	wavesShapeLabel := widget.NewLabel("Waves Shape")
	wavesShapeEntry := widget.NewEntry()

	//
	// Waves Speed
	//

	wavesSpeedLabel := widget.NewLabel("Waves Speed")
	wavesSpeedEntry := widget.NewEntry()

	//
	// Waves Speed Scaling
	//

	wavesSpeedScalingLabel := widget.NewLabel("Waves Speed Scaling")
	wavesSpeedScalingEntry := widget.NewEntry()

	save := widget.NewButton("Save", func() {
		water := export.Water{
			Out: "./openpbr_out/water/water.json",
			Water: vo.Water{
				Chlorophyll:           utils.ToFloat64(chlorophyllEntry),
				SuspendedSediment:     utils.ToFloat64(suspendedSedimentEntry),
				CDOM:                  utils.ToFloat64(cdomEntry),
				WavesEnabled:          wavesEnabledEntry.Checked,
				WavesDepth:            utils.ToFloat64(wavesDepthEntry),
				WavesFrequency:        utils.ToFloat64(wavesFrequencyEntry),
				WavesFrequencyScaling: utils.ToFloat64(wavesFrequencyScalingEntry),
				WavesMix:              utils.ToFloat64(wavesMixEntry),
				WavesOctaves:          utils.ToFloat64(wavesOctavesEntry),
				WavesPull:             utils.ToFloat64(wavesPullEntry),
				WavesSampleWidth:      utils.ToFloat64(wavesSampleWidthEntry),
				WavesShape:            utils.ToFloat64(wavesShapeEntry),
				WavesSpeed:            utils.ToFloat64(wavesSpeedEntry),
				WavesSpeedScaling:     utils.ToFloat64(wavesSpeedScalingEntry),
			},
		}

		water.Perform()

	})

	c := container.New(layout.NewFormLayout(),
		chlorophyllLabel, chlorophyllEntry,
		suspendedSedimentLabel, suspendedSedimentEntry,
		cdomLabel, cdomEntry,
		wavesEnabledLabel, wavesEnabledEntry,
		wavesDepthLabel, wavesDepthEntry,
		wavesFrequencyLabel, wavesFrequencyEntry,
		wavesFrequencyScalingLabel, wavesFrequencyScalingEntry,
		wavesMixLabel, wavesMixEntry,
		wavesOctavesLabel, wavesOctavesEntry,
		wavesPullLabel, wavesPullEntry,
		wavesSampleWidthLabel, wavesSampleWidthEntry,
		wavesShapeLabel, wavesShapeEntry,
		wavesSpeedLabel, wavesSpeedEntry,
		wavesSpeedScalingLabel, wavesSpeedScalingEntry,
		save, layout.NewSpacer())
	return c
}
