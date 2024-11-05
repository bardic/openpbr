package cmd

import (
	"archive/zip"
	"io"
	"os"

	"github.com/bardic/openpbr/utils"
)

type PackBundle struct {
	InDir  string
	OutDir string
}

func (cmd *PackBundle) Perform() error {
	utils.AppendLoadOut("--- Package Release")
	archive, err := os.Create("openpbr.mcpack")
	if err != nil {
		return err
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)
	addFileToZip(zipWriter, "openpbr_out")
	zipWriter.Close()

	return nil
}

func addFileToZip(zipWriter *zip.Writer, filePath string) error {
	files, err := os.ReadDir(filePath)

	if err != nil {
		return err
	}

	for _, item := range files {
		if item.IsDir() {
			addFileToZip(zipWriter, filePath+string(os.PathSeparator)+item.Name())
			continue
		}

		f1, err := os.Open(filePath + string(os.PathSeparator) + item.Name())
		if err != nil {
			return err
		}
		defer f1.Close()

		subpath, _ := utils.GetTextureSubpath(filePath+string(os.PathSeparator)+item.Name(), utils.OutDir)
		w1, err := zipWriter.Create(subpath)
		if err != nil {
			return err
		}
		if _, err := io.Copy(w1, f1); err != nil {
			return err
		}
	}

	return nil
}
