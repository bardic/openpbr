package utils

import (
	"errors"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

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

func RunCmd(cmd *exec.Cmd) error {
	cmd.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x08000000} // CREATE_NO_WINDOW
	go cmd.Start()
	return nil
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
		Steps: make([]vo.EntryViewHolder, 0),
	}

	steps := make([]vo.EntryViewHolder, 0)

	vBox := container.NewVBox()
	label := widget.NewLabel(title)
	vStepBox := container.NewVBox()

	addBtn := widget.NewButton("Add", func() {

		entryViewHolder := &vo.EntryViewHolder{}
		entryViewHolder.Id = id
		entryViewHolder.HBox = container.NewHBox()
		entryViewHolder.KeyEntry = widget.NewEntry()
		entryViewHolder.ValueEntry = widget.NewEntry()
		entryViewHolder.DeleteButton = widget.NewButton("Delete", func() {
			for i, v := range vStepBox.Objects {
				if v == entryViewHolder.HBox {
					vStepBox.Remove(v)
					steps = append(steps[:i], steps[i+1:]...)
					break
				}
			}

			//refresh()
		})

		hbox := entryViewHolder.HBox
		hbox.Add(widget.NewLabel("Key"))
		hbox.Add(entryViewHolder.KeyEntry)
		hbox.Add(widget.NewLabel("Value"))
		hbox.Add(entryViewHolder.ValueEntry)
		hbox.Add(entryViewHolder.DeleteButton)
		steps = append(steps, *entryViewHolder)
		vStepBox.Add(steps[len(steps)-1].HBox)
		entryView.Steps = steps
		//refresh()
		id++
	})

	vBox.Add(label)
	vBox.Add(vStepBox)
	vBox.Add(addBtn)

	entryView.Steps = steps
	entryView.C = vBox

	return entryView
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

func ToFloat64(entry *widget.Entry) float64 {
	f, err := strconv.ParseFloat(entry.Text, 64)

	if err != nil {
		return 0.0
	}

	return f
}

func StepsToVO(steps []vo.EntryViewHolder) []vo.EntryViewVO {
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
