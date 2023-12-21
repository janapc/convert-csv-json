package cmd

import (
	"encoding/json"
	"os"
	"slices"
	"strings"

	"github.com/janapc/convert-csv-json/utils"
	"github.com/spf13/cobra"
)

func getHeaders(data []map[string]interface{}) []string {
	headers := []string{}
	for _, d := range data {
		for key, _ := range d {
			containsKey := slices.Contains(headers, key)
			if !containsKey {
				headers = append(headers, key)
			}
		}
	}
	return headers
}
func formatCsv(headers []string, data []map[string]interface{}) string {
	csv := strings.Join(headers[:], ";") + "\n"
	for _, d := range data {
		line := []string{}
		for _, k := range headers {
			if str, ok := d[k].(string); ok {
				line = append(line, str)
			} else {
				line = append(line, "")
			}
		}
		csv += strings.Join(line[:], ";") + "\n"
	}
	return csv
}

func convertData(file []byte) ([]map[string]interface{}, error) {
	var data []map[string]interface{}
	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// csvCmd represents the csv command
var csvCmd = &cobra.Command{
	Use:   "convert-to-csv",
	Short: "Convert json to csv",
	Long:  "Convert json to csv",
	RunE: func(cmd *cobra.Command, args []string) error {
		pathname, _ := cmd.Flags().GetString("pathname")
		destination, _ := cmd.Flags().GetString("destination")
		m := map[string]string{
			"json": pathname,
			"csv":  destination,
		}
		err := utils.ValidPaths(m)
		if err != nil {
			return err
		}
		utils.LogInfo("Processing data...")
		file, err := os.ReadFile(pathname)
		if err != nil {
			return err
		}
		data, err := convertData(file)
		if err != nil {
			return err
		}
		headers := getHeaders(data)
		csv := formatCsv(headers, data)
		f, err := os.Create(destination)
		if err != nil {
			return err
		}
		defer f.Close()
		utils.LogInfo("Created file...")
		_, err = f.WriteString(csv)
		if err != nil {
			return err
		}
		utils.LogInfo("File created")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(csvCmd)
	csvCmd.Flags().StringP("pathname", "p", "", "pathname ex: 'data.json'")
	csvCmd.Flags().StringP("destination", "d", "", "destination of file ex:'./data.csv'")
	csvCmd.MarkFlagsRequiredTogether("pathname", "destination")
}
