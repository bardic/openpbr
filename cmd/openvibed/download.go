package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type downloadCmd struct {
	appfs    afero.Fs
	url      string
	base     string
	download string
}

func NewDownloadCmd(fs afero.Fs) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download",
		Short: "downloads latest zip of assets from microsoft's github",
		RunE: func(cmd *cobra.Command, args []string) error {
			downloadDir := "./downloads"
			baseDir := "./base"
			downloadURL := viper.GetString("url")

			dlCmd := downloadCmd{
				appfs:    fs,
				url:      downloadURL,
				base:     baseDir,
				download: downloadDir,
			}

			if err := dlCmd.prepDisk(); err != nil {
				return err
			}

			if err := dlCmd.DownloadRelease(); err != nil {
				return err
			}

			if err := dlCmd.extract(); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().String("url", "", "-")
	viper.BindPFlag("url", cmd.Flags().Lookup("url"))
	return cmd

}

func (c *downloadCmd) prepDisk() error {
	err := c.appfs.MkdirAll(c.download, os.ModePerm)

	if err != nil {
		return err
	}

	err = c.appfs.MkdirAll(c.base, os.ModePerm)

	if err != nil {
		return err
	}

	return nil
}

// r.ZipballUrl, latestZipPath
func (c *downloadCmd) DownloadRelease() error {

	r, err := URLToJson[GithubRelease](c.url)

	if err != nil {
		fmt.Println("meowmeow")
		return err
	}

	fmt.Println("dowload url")
	fmt.Println(r)

	resp, err := http.Get(c.url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return err
	}

	// Create the file
	out, err := c.appfs.Create(c.download)
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

func (c *downloadCmd) extract() error {
	latestZipPath := filepath.Join(c.download, "latest.zip")
	archive, err := zip.OpenReader(latestZipPath)
	if err != nil {
		return err
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(c.base, f.Name)
		if f.FileInfo().IsDir() {
			if err := c.appfs.MkdirAll(filePath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		if err := c.appfs.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := c.appfs.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
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
