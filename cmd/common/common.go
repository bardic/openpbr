package common

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/bardic/openpbr/cmd/gen"
	"github.com/bardic/openpbr/cmd/img"
	"github.com/bardic/openpbr/cmd/utils"
	"github.com/spf13/cobra"
)

// Build is a recursive function that processes the images and generates json, mer, and height files.
func Build(cmd *cobra.Command, target string, imgPath string) error {
	subPaths := strings.Split(imgPath, string(os.PathSeparator))
	items, _ := os.ReadDir(imgPath)

	for _, item := range items {
		outPath := utils.OutDir + string(os.PathSeparator) + strings.Join(subPaths[3:], string(os.PathSeparator)) + string(os.PathSeparator) + item.Name()
		if item.IsDir() {
			if err := os.MkdirAll(outPath, os.ModePerm); err != nil {
				return err
			}
			p := filepath.Join(imgPath, item.Name())
			Build(cmd, target, p)
		} else {

			in := filepath.Join(imgPath, item.Name())

			if strings.Contains(outPath, ".tga") {
				img.TgaPngCmd.RunE(cmd, []string{in, strings.ReplaceAll(in, ".tga", ".png")})
				in = strings.ReplaceAll(in, ".tga", ".png")
				outPath = strings.ReplaceAll(outPath, ".tga", ".png")
			}

			if !strings.Contains(outPath, ".png") {
				continue
			}

			err := utils.CopyF(in, outPath)
			if err != nil {
				return err
			}

			img.AdjustColorCmd.Run(cmd, []string{outPath})

			if utils.NormalMaps {
				err := gen.NormalCmd.RunE(cmd, []string{outPath, strings.ReplaceAll(outPath, ".png", "_normal.png")})
				if err != nil {
					return err
				}
			} else {
				err := gen.HeightCmd.RunE(cmd, []string{outPath, strings.ReplaceAll(outPath, ".png", "_height.png")})
				if err != nil {
					fmt.Println(err)
					return err
				}
			}

		}
	}

	return nil
}

func CreateMers(cmd *cobra.Command, inputPath string) error {

	if filepath.Ext(inputPath) == ".tga" {
		return nil
	}

	subPaths := strings.Split(inputPath, string(os.PathSeparator))
	items, _ := os.ReadDir(inputPath)

	for _, item := range items {
		if item.IsDir() {
			CreateMers(cmd, inputPath+string(os.PathSeparator)+item.Name())
		} else {
			outPath := utils.OutDir + string(os.PathSeparator) + strings.Join(subPaths[3:], string(os.PathSeparator)) + string(os.PathSeparator) + item.Name()
			outPath = strings.Replace(outPath, ".tga", ".png", 1)
			merPath := strings.ReplaceAll(outPath, ".png", "_mer.png")

			useMerFile := true
			if _, err := os.Stat(merPath); err != nil {
				useMerFile = false
			}

			fn := strings.TrimSuffix(item.Name(), filepath.Ext(item.Name()))

			merArr, err := merLookup(strings.Join(subPaths[3:], string(os.PathSeparator)) + string(os.PathSeparator) + item.Name())

			if err != nil {
				fmt.Println(err)
				return err
			}

			merFile := fn + "_mer"

			heightNormalFile := fn + "_height"
			if utils.NormalMaps {
				heightNormalFile = fn + "_normal"
			}

			err = gen.JsonCmd.RunE(cmd, []string{
				strings.ReplaceAll(outPath, ".png", ".texture_set.json"),
				fn,
				merArr,
				merFile,
				heightNormalFile,
				strconv.FormatBool(useMerFile),
				utils.TexturesetVersion,
			})

			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	return nil
}

func merLookup(asset string) (string, error) {
	f, err := os.Open("mer.csv")
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, record := range records {
		if record[0] == asset {
			if utils.TexturesetVersion != "1.21.30" {
				return "[" + record[1] + "," + record[2] + "," + record[3] + "]", nil
			}

			return "[" + record[1] + "," + record[2] + "," + record[3] + "," + record[4] + "]", nil
		}
	}

	if utils.TexturesetVersion != "1.21.30" {
		return "[0,0,255]", nil
	}

	return "[0,0,255,255]", nil
}
