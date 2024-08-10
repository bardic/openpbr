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

var CreateCSVCmd = &cobra.Command{
	Use:   "createcsv",
	Short: "create a base csv to modify",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("cats")
		for _, target := range utils.TargetAssets {
			CreateBaseCSV(args[0] + string(os.PathSeparator) + target)
		}

		return nil
	},
}

var records = [][]string{
	{"path", "metalness", "emissive", "roughness", "subsurface"},
}

func CreateBaseCSV(in string) {
	fmt.Println(in)

	f, err := os.Create("mer.csv")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer f.Close()

	scan(in)

	fmt.Println(records)
	w := csv.NewWriter(f)
	w.WriteAll(records)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

func scan(in string) {
	subPaths := strings.Split(in, string(os.PathSeparator))
	items, _ := os.ReadDir(in)
	for _, item := range items {

		if item.IsDir() {
			scan(filepath.Join(in, item.Name()))
			continue
		}

		records = append(records, []string{
			strings.Join(subPaths[3:], string(os.PathSeparator)) + string(os.PathSeparator) + item.Name(),
			"0",
			"0",
			"255",
			"255",
		})

	}
}
