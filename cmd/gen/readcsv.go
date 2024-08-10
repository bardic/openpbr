package gen

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var ReadCSVCmd = &cobra.Command{
	Use:   "readcsv",
	Short: "create a base csv to modify",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		ReadCSV()
		return nil
	},
}

func ReadCSV() {
	f, err := os.Open("mer.csv")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
}
