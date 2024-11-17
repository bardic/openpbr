package utils

import (
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/bardic/openpbr/vo"
)

const BaseAssets = "input"
const OutDir = "openpbr_out"
const Overrides = "overrides"
const SettingDir = "settings"
const Psds = "psds"

var ImCmd = "magick"

var Basedir string
var TargetAssets = []string{"blocks", "entity", "particle"}

var HeightMapNameSuffix = "_height"
var MerMapNameSuffix = "_mer"

var LoadStdOut *widget.RichText

func LocalPath(partialPath string) string {
	return Basedir + string(os.PathSeparator) + partialPath
}

func StartUpCheck() error {
	if _, err := exec.LookPath(ImCmd); err != nil {
		return errors.New("imagemagick not found")
	}

	return nil
}

func AppendLoadOut(s string) {
	LoadStdOut.AppendMarkdown(s)
	LoadStdOut.Refresh()
}

func GetTextureSubpath(p string, key string) (string, error) {
	subpaths := strings.Split(p, string(os.PathSeparator))
	for i, subpath := range subpaths {
		if subpath == key {
			sub := strings.Join(subpaths[i:], string(os.PathSeparator))
			return sub, nil
		}
	}

	return "", errors.New("")
}

func CreateEntryView(title string, id int) *vo.EntryView {

	entryView := &vo.EntryView{
		Steps: make([]*vo.EntryViewHolder, 0),
		C:     container.NewVBox(),
	}

	label := widget.NewLabel(title)
	addBtn := widget.NewButton("Add", func() {
		entryViewHolder := CreateEntryViewHolder()
		entryView.Steps = append(entryView.Steps, entryViewHolder)
		entryView.C.Add(entryView.Steps[len(entryView.Steps)-1].HBox)
		del := func() {
			entryView.C.Remove(entryViewHolder.HBox)
		}
		entryViewHolder.DeleteButton.OnTapped = del

		id++
	})

	entryView.C.Add(label)
	entryView.C.Add(addBtn)

	return entryView
}

func CreateEntryViewHolder() *vo.EntryViewHolder {
	keyEntry := widget.NewEntry()
	valueEntry := widget.NewEntry()
	deleteButton := widget.NewButton("Delete", nil)

	hBox := container.NewHBox()
	hBox.Add(widget.NewLabel("Key"))
	hBox.Add(keyEntry)
	hBox.Add(widget.NewLabel("Value"))
	hBox.Add(valueEntry)
	hBox.Add(deleteButton)

	return &vo.EntryViewHolder{
		KeyEntry:     keyEntry,
		ValueEntry:   valueEntry,
		DeleteButton: deleteButton,
		HBox:         hBox,
	}
}

func CreateRGBEntry() *vo.RGB {

	highlightsContrastRLabel := widget.NewLabel("R")
	highlightsContrastREntry := widget.NewEntry()
	highlightsContrastGLabel := widget.NewLabel("G")
	highlightsContrastGEntry := widget.NewEntry()
	highlightsContrastBLabel := widget.NewLabel("B")
	highlightsContrastBEntry := widget.NewEntry()

	highlightsContrastHBox := container.NewAdaptiveGrid(7, highlightsContrastRLabel, highlightsContrastREntry, highlightsContrastGLabel, highlightsContrastGEntry, highlightsContrastBLabel, highlightsContrastBEntry)
	highlightsContrastVBox := container.NewHBox(highlightsContrastHBox)

	return &vo.RGB{
		R: highlightsContrastREntry,
		G: highlightsContrastGEntry,
		B: highlightsContrastBEntry,
		C: highlightsContrastVBox,
	}
}

func CreateRGBAEntry() *vo.RGBA {

	rgba := CreateRGBEntry()

	highlightsContrastRLabel := widget.NewLabel("R")
	highlightsContrastGLabel := widget.NewLabel("G")
	highlightsContrastBLabel := widget.NewLabel("B")
	highlightsContrastALabel := widget.NewLabel("A")
	highlightsContrastAEntry := widget.NewEntry()

	highlightsContrastHBox := container.NewAdaptiveGrid(9, highlightsContrastRLabel, rgba.R, highlightsContrastGLabel, rgba.G, highlightsContrastBLabel, rgba.B, highlightsContrastALabel, highlightsContrastAEntry)
	highlightsContrastVBox := container.NewHBox(highlightsContrastHBox)

	rgba.C = highlightsContrastVBox

	return &vo.RGBA{
		RGB: *rgba,
		A:   highlightsContrastAEntry,
	}
}

func ToFloat64(entry *widget.Entry) float64 {
	f, err := strconv.ParseFloat(entry.Text, 64)

	if err != nil {
		return 0.0
	}

	return f
}

func StepsToVO(steps []*vo.EntryViewHolder) []vo.EntryViewVO {
	var voSteps []vo.EntryViewVO
	for _, step := range steps {
		voSteps = append(voSteps, vo.EntryViewVO{
			Key:   step.KeyEntry.Text,
			Value: ToFloat64(step.ValueEntry),
		})
	}

	if len(voSteps) > 0 {
		voSteps[len(voSteps)-1].Last = true
	}

	return voSteps
}

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
