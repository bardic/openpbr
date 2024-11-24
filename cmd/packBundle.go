package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/bardic/openpbr/utils"
)

type PackBundle struct {
	InDir  string
	OutDir string
}

func (cmd *PackBundle) Perform() error {
	os.MkdirAll(path.Join(utils.Basedir, "openpbr"), os.ModePerm)
	utils.AppendLoadOut("--- Package Release")
	archive, err := os.Create(path.Join(utils.Basedir, "openpbr", "openpbr.mcpack"))
	if err != nil {
		return err
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)
	addFileToZip(zipWriter, path.Join(utils.Basedir, "export"))
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
			addFileToZip(
				zipWriter,
				path.Join(filePath, item.Name()),
			)
			continue
		}

		f1, err := os.Open(
			path.Join(filePath, item.Name()),
		)

		if err != nil {
			return err
		}

		defer f1.Close()

		// fmt.Println(path.Join(filePath, item.Name()))

		subpath, _ := utils.GetTextureSubpath(
			path.Join(filePath, item.Name()),
			utils.OutDir,
		)

		fmt.Println(subpath)

		w1, err := zipWriter.Create(path.Join("openpbr", subpath))

		if err != nil {
			return err
		}

		if _, err := io.Copy(w1, f1); err != nil {
			return err
		}
	}

	return nil
}
