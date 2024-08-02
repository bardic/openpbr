package download

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/bardic/openpbr/utils"
	"github.com/spf13/cobra"
)

type GithubRelease struct {
	Zipball_url string
}

var Cmd = &cobra.Command{
	Use:   "download",
	Short: "Downloads latest release",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		r := &GithubRelease{}
		getJson("https://api.github.com/repos/Mojang/bedrock-samples/releases/latest", r)
		donwloadRelease(r.Zipball_url)
	},
}

func getJson(url string, target interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&target)
	if err != nil {
		return err
	}

	return nil
}

func donwloadRelease(r string) {
	resp, err := http.Get(r)
	if err != nil {
		fmt.Printf("err: %s", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}

	// Create the file
	out, err := os.Create("latest.zip")
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		fmt.Printf("err: %s", err)
	}

	archive, err := zip.OpenReader("latest.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	fmt.Println("--- --- Extracting base assets")
	for _, f := range archive.File {
		filePath := filepath.Join(utils.BaseAssets, f.Name)
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				panic(err)
			}
			continue
		}
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		srcFile, err := f.Open()
		if err != nil {
			panic(err)
		}
		if _, err := io.Copy(dstFile, srcFile); err != nil {
			panic(err)
		}

		dstFile.Close()
		srcFile.Close()
	}

}
