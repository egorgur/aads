/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"pr1_4/mysort"
	"strings"

	"github.com/spf13/cobra"
)

// sortCmd represents the sort command
var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "Sort command",
	Long: `Sorting slices:
Sort cli command. Sorts array of strings of integers.
Use --file flag to read an array from file and sort it. The sorted array will be written in the same file.`,
	Run: func(cmd *cobra.Command, args []string) {
		strData, err := cmd.Flags().GetString("data")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if strData == "" && file == "" {
			fmt.Println("Error: either --data or --file must be provided")
			os.Exit(1)
		}
		if strData != "" && file != "" {
			fmt.Println("Error: both --data and --file cannot be used simultaneously")
			os.Exit(1)
		}

		if file != "" {
			var err error
			strData, err = readFile(file)
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
				os.Exit(1)
			}
		}

		reverse, err := cmd.Flags().GetBool("reverse")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		visualize, err := cmd.Flags().GetBool("no-visualize")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		
		data := strings.Split(strData, ",")

		sortedData, err := mysort.Sort(convertArray(data), reverse, visualize)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		
		result := arrayToString(sortedData)

		if file != "" {
			writeFile(file, result)
		}

		fmt.Printf("Result: %v", result)
	},
}

func readFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath) // use os package instead of io/utils
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func writeFile(filePath string, data string) (error) {
	return os.WriteFile(filePath, []byte(data), os.ModePerm)
}

func convertArray(data []string) []interface{} {
	nData := make([]interface{}, len(data))
	for i := range data {
		nData[i] = data[i]
	}
	return nData
}

func arrayToString(data []interface{}) string {
	strs := make([]string, len(data))
	for i:= range data {
		strs[i] = data[i].(string)
	}
	return strings.Join(strs, ",")
}

func init() {
	rootCmd.AddCommand(sortCmd)

	// Here you will define your flags and configuration settings.
	sortCmd.Flags().StringP("data", "d", "", "Data array to sort. Use a comma to differ the elements.")
	sortCmd.Flags().StringP("file", "f", "", "Path to the text file to sort data in. Use a comma to differ the elements.")
	sortCmd.Flags().BoolP("reverse", "r", false, "Reverse sort.")
	sortCmd.Flags().BoolP("no-visualize", "", true, "Do not show sort visualization.")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sortCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sortCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

