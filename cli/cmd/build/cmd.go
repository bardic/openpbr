package build

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bardic/openpbr/cmd/clean"
	"github.com/bardic/openpbr/cmd/download"
	"github.com/bardic/openpbr/cmd/gen"
	"github.com/bardic/openpbr/cmd/img"
	"github.com/bardic/openpbr/utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "build",
	Short: "build pack",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(time.Now().String())

		fmt.Println("--- Cleaning workspace")

		clean.Cmd.Run(cmd, nil)

		fmt.Println("--- Prcoess Glowables")

		gen.GlowablesCmd.Run(cmd, []string{"./glowables/blocks"})

		fmt.Println("--- Download latest base assets")

		download.Cmd.Run(cmd, nil)

		fmt.Println("--- Copy custom configs")

		utils.CopyD(utils.ConfigDIr, utils.BuildDir)

		entries, _ := os.ReadDir(utils.BaseAssets)
		f := entries[0]

		for _, s := range utils.TaretAssets {
			fmt.Println("--- Create json, mer and height files for " + s)

			build(cmd, s, utils.BaseAssets+"/"+f.Name()+"/resource_pack/textures/"+s)
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}
		}
	},
}

func build(cmd *cobra.Command, target string, imgPath string) {
	fmt.Println(imgPath)
	items, _ := os.ReadDir(imgPath)

	fmt.Println(len(items))

	for _, item := range items {
		if item.IsDir() {
			outputPath := strings.ReplaceAll(imgPath, utils.BaseAssets, utils.BuildDir)
			if err := os.MkdirAll(outputPath+"/"+item.Name(), 0770); err != nil {
				return
			}
			build(cmd, target, imgPath+"/"+item.Name())
		} else {
			fmt.Println(item.Name())
			if !strings.Contains(item.Name(), ".tga") && !strings.Contains(item.Name(), ".png") {
				continue
			}

			fmt.Println("Is right type" + target)

			in := imgPath + "/" + item.Name()

			out := utils.BuildDir + "/textures/" + target + "/" + item.Name()

			fmt.Println("New in" + in)
			fmt.Println("New out" + out)

			if strings.Contains(item.Name(), ".tga") {
				out = strings.ReplaceAll(out, ".tga", ".png")
				img.TgaPngCmd.SetArgs([]string{""})
				img.TgaPngCmd.Execute()
			} else {
				err := utils.CopyF(in, out)
				if err != nil {
					return
				}
			}
			gen.HeightCmd.Run(cmd, []string{out, strings.ReplaceAll(out, ".png", "_height.png")})

			_, b := utils.CheckForOverride(strings.ReplaceAll(out, ".png", "_mer.png"))

			fn := strings.TrimSuffix(item.Name(), filepath.Ext(item.Name()))

			MerType := 1
			if b {
				MerType = 0
			}

			gen.JsonCmd.Run(cmd, []string{
				strings.ReplaceAll(out, ".png", ".texture_set.json"),
				fn,
				"[0, 0, 255]",
				fn + "_mer",
				fn + "_height",
				fmt.Sprint(MerType),
			})
		}
	}
}
