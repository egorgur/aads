/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"pr1_4/mysort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/charmbracelet/lipgloss"
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

		sortChan := make(chan []interface{})
		resultChan := make(chan []interface{})

		go func() {
			mysort.Sort(convertArray(data), reverse, visualize, sortChan, resultChan)
		}()

		for arr := range sortChan {

			renderChart(arr)

			fmt.Printf("sort step %#+v\n", arr)
		}

		sortedData := <-resultChan

		result := arrayToString(sortedData)

		if file != "" {
			writeFile(file, result)
		}

		fmt.Printf("Result: %v\n", result)

		renderChart(sortedData)
	},
}

func readFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath) // use os package instead of io/utils
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func writeFile(filePath string, data string) error {
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
	for i := range data {
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

// VERTICAL BAR CHART

var defaultStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("63")) // purple

var axisStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("3")) // yellow

var labelStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("63")) // purple

var blockStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("9")) // red

var blockStyle2 = lipgloss.NewStyle().
	Foreground(lipgloss.Color("2")) // green

var blockStyle3 = lipgloss.NewStyle().
	Foreground(lipgloss.Color("6")) // cyan

var blockStyle4 = lipgloss.NewStyle().
	Foreground(lipgloss.Color("3")) // yellow

func renderChart(arr []interface{}) {

	var containsInts bool

	containsInts = onlyInts(arr) // If only integers are present

	var containsStrings bool = false

	if !containsInts {
		containsStrings = onlyStrings(arr) // If only strings are present
	}

	if containsInts {
		maxLength := 0
		for _, element := range arr {
			if value, _ := strconv.Atoi(element.(string)); value > maxLength {
				maxLength = value
			}
		}

		// Выводим диаграмму
		for _, element := range arr {
			value, _ := strconv.Atoi(element.(string))
			bar := ""
			for i := 0; i < value; i++ {
				bar += "█" // Символ для столбика
			}
			fmt.Printf("%-10s | %s (%d)\n", element, bar, value)
		}
		return
	}

	if containsStrings {
		// max length for strings
		maxLength := 0
		for _, element := range arr {
			if len(element.(string)) > maxLength {
				maxLength = len(element.(string))
			}
		}

		// Выводим диаграмму
		for _, element := range arr {
			bar := ""
			for i := 0; i < len(element.(string)); i++ {
				bar += "█" // Символ для столбика
			}
			fmt.Printf("%-10s | %s (%d)\n", element, bar, len(element.(string)))
		}
		return
	}

}

func onlyInts(input_array []interface{}) bool {
	for _, v := range input_array {
		_, ok := v.(int)
		if !ok {
			_, err := strconv.Atoi(v.(string))
			if err != nil {
				return false
			}
		}
	}
	return true
}

func onlyStrings(input_array []interface{}) bool {
	for _, v := range input_array {
		_, ok := v.(string)
		if !ok {
			return false
		}
	}
	return true
}
