/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"pr3/internal"
	"strings"

	"github.com/spf13/cobra"
)

type searchParameters struct {
	text            string
	pattern         []string
	caseSensitivity bool
	method          string
	count           uint
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search all patterns in a given string or text file",
	Long: `Search all patterns in a given string or text file
	
	Flags:
	    --text		give a text in which the search will find the patterns
	-f, --file		set the path to a text file to open and read
	-p, --pattern		(can be multiple) set the pattern to find in a given text
	-c, --case		case sensitivity (default "false")
	-r, --reverse		find the patterns in reversed text (the indexes will correspond to the unreversed text)
	    --count		set the maximum number of possible found patterns
	`,
	Run: func(cmd *cobra.Command, args []string) {
		parameters := &searchParameters{
			text:            "",
			pattern:         []string{""},
			caseSensitivity: false,
			method:          "first",
			count:           99999,
		}

		text, err := cmd.Flags().GetString("text")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if text == "" && file == "" {
			fmt.Println("Error: either --text or --file must be provided")
			os.Exit(1)
		}
		if text != "" && file != "" {
			fmt.Println("Error: both --text and --file cannot be used simultaneously")
			os.Exit(1)
		}

		if file != "" {
			var err error
			text, err = readFile(file)
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
				os.Exit(1)
			}
		}

		parameters.text = text

		substrings, err := cmd.Flags().GetString("pattern")
		if err == nil {
			parameters.pattern = strings.Split(substrings, ",")
		}

		caseSensitive, err := cmd.Flags().GetBool("case-sensitive")

		if err == nil {
			parameters.caseSensitivity = bool(caseSensitive)
		}

		method, err := cmd.Flags().GetString("method")
		if err == nil {
			parameters.method = method
		}

		count, err := cmd.Flags().GetInt("count")
		if err == nil {
			if count != 0 {
				parameters.count = uint(count)
			}
		}

		results := search.Search(
			parameters.text,
			parameters.pattern,
			parameters.caseSensitivity,
			parameters.method,
			int(parameters.count))

		for subString, positions := range results {
			fmt.Println(subString,positions)
			printResult(parameters.text, subString, positions)
		}
	},
}

func printResult(text string, subString string, positions []int) {
	lettersList := strings.Split(text, "")
	highlightColor := "\033[31m"
	resetColor := "\033[0m"
	for _ ,pos := range positions {
		lettersList[pos] = highlightColor + lettersList[pos]
		lettersList[pos + len(subString) - 1] = lettersList[pos + len(subString) - 1] + resetColor
	}
	fmt.Println(subString+" "+strings.Join(lettersList,""))
}

func readFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	searchCmd.Flags().StringP("text", "", "", "Text to search in")
	searchCmd.Flags().StringP("file", "f", "", "Path to the text file to search in")
	searchCmd.Flags().StringP("pattern", "s", "", "Comma-separated list of substrings to search for")
	searchCmd.Flags().BoolP("case-sensitive", "c", false, "Enable case-sensitive search")
	searchCmd.Flags().StringP("method", "m", "first", "Search method: 'first' for forward search, 'last' for reverse search")
	searchCmd.Flags().IntP("count", "n", 0, "Number of matches to find (0 means find all matches)")

	searchCmd.MarkFlagRequired("substrings")
}

