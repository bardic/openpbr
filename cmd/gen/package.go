package gen

import (
	"archive/zip"
	"fmt"
	"io"
	"os"

	"github.com/bardic/openpbr/cmd/utils"
	"github.com/spf13/cobra"
)

var buildDir string

var PackageCmd = &cobra.Command{
	Use:   "package",
	Short: "package project",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		buildDir = args[0]

		utils.AppendLoadOut("--- Creating zip archive...")
		archive, err := os.Create(utils.LocalPath("openpbr.mcpack"))
		if err != nil {
			return err
		}
		defer archive.Close()
		zipWriter := zip.NewWriter(archive)
		subpath, _ := utils.GetTextureSubpath(buildDir, "openpbr")
		addFileToZip(zipWriter, utils.LocalPath(subpath))
		zipWriter.Close()

		return nil
	},
}

func addFileToZip(zipWriter *zip.Writer, filePath string) error {
	files, err := os.ReadDir(filePath)

	if err != nil {
		fmt.Println(err)
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

		subpath, _ := utils.GetTextureSubpath(filePath+string(os.PathSeparator)+item.Name(), "openpbr")
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
