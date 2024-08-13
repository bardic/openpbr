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
	items, err := os.ReadDir(imgPath)

	if err != nil {
		fmt.Println(err)
	}

	for _, item := range items {

		if item.IsDir() {
			p := filepath.Join(imgPath, item.Name())

			Build(cmd, target, p)
		} else {
			subpath, err := utils.GetTextureSubpath(imgPath, "textures")

			if err != nil {
				return err

			}

			in := filepath.Join(imgPath, item.Name())
			outPath := utils.LocalPath(utils.OutDir + string(os.PathSeparator) + subpath + string(os.PathSeparator) + item.Name())
			os.MkdirAll(filepath.Dir(outPath), os.ModePerm)

			if strings.Contains(outPath, ".tga") {
				err := img.TgaPngCmd.RunE(cmd, []string{in, strings.ReplaceAll(outPath, ".tga", ".png")})
				if err != nil {
					return err
				}
				in = strings.ReplaceAll(in, ".tga", ".png")
				outPath = strings.ReplaceAll(outPath, ".tga", ".png")
			}

			if !strings.Contains(outPath, ".png") {
				continue
			}

			go utils.CopyF(in, outPath)
			// if err != nil {
			// 	return err
			// }

			img.AdjustColorCmd.Run(cmd, []string{outPath})

			if utils.NormalMaps {
				err := gen.NormalCmd.RunE(cmd, []string{outPath, strings.ReplaceAll(outPath, ".png", utils.NormalMapNameSuffix+".png")})
				if err != nil {
					return err
				}
			} else {
				err := gen.HeightCmd.RunE(cmd, []string{outPath, strings.ReplaceAll(outPath, ".png", utils.HeightMapNameSuffix+".png")})
				if err != nil {
					fmt.Println(err)
					return err
				}
			}
		}
	}

	return nil
}

func CreateMers(cmd *cobra.Command, out string) error {
	items, _ := os.ReadDir(out)

	for _, item := range items {
		if item.IsDir() {
			CreateMers(cmd, out+string(os.PathSeparator)+item.Name())
		} else {
			q, e := utils.GetTextureSubpath(out, "textures")

			if e != nil {
				return e
			}

			outPath := utils.LocalPath(utils.OutDir + string(os.PathSeparator) + q + string(os.PathSeparator) + item.Name())
			outPath = strings.Replace(outPath, ".tga", ".png", 1)
			merPath := strings.ReplaceAll(outPath, ".png", utils.MerMapNameSuffix+".png")

			useMerFile := true
			if _, err := os.Stat(merPath); err != nil {
				useMerFile = false
			}

			fn := strings.TrimSuffix(item.Name(), filepath.Ext(item.Name()))

			merArr, err := merLookup(q + string(os.PathSeparator) + item.Name())

			if err != nil {
				fmt.Println(err)
				return err
			}

			heightNormalFile := fn + utils.HeightMapNameSuffix
			if utils.NormalMaps {
				heightNormalFile = fn + utils.NormalMapNameSuffix
			}

			err = gen.JsonCmd.RunE(cmd, []string{
				strings.ReplaceAll(outPath, ".png", ".texture_set.json"),
				fn,
				merArr,
				fn + utils.MerMapNameSuffix,
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
	f, err := os.Open(utils.LocalPath("mer.csv"))
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
