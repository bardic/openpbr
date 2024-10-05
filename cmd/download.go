package cmd

import (
	"archive/zip"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/bardic/openpbr/utils"
	"github.com/spf13/cobra"
)

var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Downloads latest release",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		r := &GithubRelease{}
		if e := getJson("https://api.github.com/repos/Mojang/bedrock-samples/releases/latest", r); e != nil {
			return e
		}
		if e := downloadRelease(r.Zipball_url); e != nil {
			return e
		}

		err := extract()

		if err != nil {
			return err
		}

		return nil
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

func downloadRelease(r string) error {
	resp, err := http.Get(r)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return err
	}

	// Create the file
	out, err := os.Create(utils.LocalPath("latest.zip"))
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return err
	}

	return nil
}

func extract() error {
	archive, err := zip.OpenReader(utils.LocalPath("latest.zip"))
	if err != nil {
		return err
	}
	defer archive.Close()

	utils.AppendLoadOut("--- --- Extracting base assets")
	for _, f := range archive.File {
		filePath := filepath.Join(utils.LocalPath(utils.BaseAssets), f.Name)
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				return err
			}
			continue
		}
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		srcFile, err := f.Open()
		if err != nil {
			return err
		}
		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return err
		}

		dstFile.Close()
		srcFile.Close()
	}

	return nil
}
