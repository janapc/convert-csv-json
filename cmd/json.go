/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/janapc/convert-csv-json/utils"
	"github.com/spf13/cobra"
)

func getFirstLine(file string) string {
	regexFisrtLine := regexp.MustCompile("(?m)^(.*)$")
	firstLine := regexFisrtLine.FindString(file)
	return firstLine
}

func getKeys(file string, separator string) []string {
	firstLine := getFirstLine(file)
	headers := strings.Split(firstLine, separator)
	return headers
}

func getValues(file string) []string {
	firstLine := getFirstLine(file)
	values := strings.Replace(file, firstLine, "", -1)
	lines := strings.Split(values, "\n")
	return lines
}

func formatMessage(values []string, keys []string, separator string) string {
	var msg string
	for index, value := range values {
		if index == 0 {
			continue
		}
		if value == "" && len(values)-1 != index {
			continue
		}
		if value == "" && len(values)-1 == index {
			msg = msg[:len(msg)-1]
			msg += "]"
			continue
		}
		v := strings.Replace(value, "\r", "", -1)
		valuesFormatted := strings.Split(v, separator)
		if index == 1 {
			msg += "[{"
		} else {
			msg += "{"
		}
		for i, key := range keys {
			k := strings.Replace(key, "\r", "", -1)
			if len(keys)-1 == i {
				msg += fmt.Sprintf(`"%s" : "%s"`, k, string(valuesFormatted[i]))
			} else {
				msg += fmt.Sprintf(`"%s" : "%s" ,`, k, string(valuesFormatted[i]))
			}
		}
		if len(values)-1 == index {
			msg += "}]"
		} else {
			msg += "},"
		}
	}
	return msg
}

// convertJsonCmd represents the convertJson command
var jsonCmd = &cobra.Command{
	Use:   "convert-to-json",
	Short: "Convert csv to json",
	Long:  `Convert csv to json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pathname, _ := cmd.Flags().GetString("pathname")
		separator, _ := cmd.Flags().GetString("separator")
		destination, _ := cmd.Flags().GetString("destination")
		m := map[string]string{
			"csv":  pathname,
			"json": destination,
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
		keys := getKeys(string(file), separator)
		values := getValues(string(file))
		f, err := os.Create(destination)
		if err != nil {
			return err
		}
		defer f.Close()
		msg := formatMessage(values, keys, separator)
		utils.LogInfo("Created file...")
		_, err = f.WriteString(msg)
		if err != nil {
			return err
		}
		utils.LogInfo("File created")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)
	jsonCmd.Flags().StringP("pathname", "p", "", "pathname ex: 'data.csv'")
	jsonCmd.Flags().StringP("separator", "s", "", "delimiter of data ex:','")
	jsonCmd.Flags().StringP("destination", "d", "", "destination of file ex:'./data.json'")
	jsonCmd.MarkFlagsRequiredTogether("pathname", "separator", "destination")
}
