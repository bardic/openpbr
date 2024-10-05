package gen

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bardic/openpbr/cmd/utils"
	"github.com/spf13/cobra"
)

var defaultMer = []string{"0", "0", "255", "255"}

var CreateCSVCmd = &cobra.Command{
	Use:   "createcsv",
	Short: "create a base csv to modify",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		defaultMer = strings.Split(args[1], ",")

		f, err := os.Create(utils.LocalPath("mer.csv"))

		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		defer f.Close()

		scan(args[0] + string(os.PathSeparator) + "textures")
		w := csv.NewWriter(f)
		w.WriteAll(records)

		if err := w.Error(); err != nil {
			log.Fatalln("error writing csv:", err)
		}

		return nil
	},
}

var records = [][]string{
	{"path", "metalness", "emissive", "roughness", "subsurface"},
}

func scan(in string) {
	items, err := os.ReadDir(in)

	if err != nil {
		fmt.Println(err)
	}

	for _, item := range items {

		if item.IsDir() {
			scan(filepath.Join(in, item.Name()))
			continue
		}

		ss := "255"
		if len(defaultMer) > 3 {
			ss = defaultMer[4]
		}

		p, e := utils.GetTextureSubpath(in, "textures")

		if e != nil {
			return
		}

		records = append(records, []string{
			p + string(os.PathSeparator) + item.Name(),
			defaultMer[0],
			defaultMer[1],
			defaultMer[2],
			ss,
		})

	}
}
