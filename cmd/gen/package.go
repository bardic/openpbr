package gen

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var buildDir string

var PackageCmd = &cobra.Command{
	Use:   "package",
	Short: "package project",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		buildDir = args[0]

		fmt.Println("--- Creating zip archive...")
		archive, err := os.Create("openpbr.mcpack")
		if err != nil {
			return err
		}
		defer archive.Close()
		zipWriter := zip.NewWriter(archive)
		addFileToZip(zipWriter, buildDir)
		zipWriter.Close()

		return nil
	},
}

func addFileToZip(zipWriter *zip.Writer, filePath string) error {
	files, _ := os.ReadDir(filePath)
	for _, item := range files {
		if item.IsDir() {
			addFileToZip(zipWriter, filePath+"/"+item.Name())
			continue
		}

		f1, err := os.Open(filePath + "/" + item.Name())
		if err != nil {
			return err
		}
		defer f1.Close()

		packagePath := strings.Replace(filePath, "./output/", "", 1)
		w1, err := zipWriter.Create(packagePath + "/" + item.Name())
		if err != nil {
			return err
		}
		if _, err := io.Copy(w1, f1); err != nil {
			return err
		}
	}

	return nil
}
