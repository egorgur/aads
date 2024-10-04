package search

import "strings"

type Method uint8 // Method enum

const (
	First Method = iota // Method enum
	Last                // Method enum
)

type searchParameters struct {
	// Search func parameters
	inputString     string
	subString       []string
	caseSensitivity bool
	method          Method // A Method to search the sub strings
	count           uint   // Optional
}

func WithCaseSensitivity(case_sensitivity bool) func(*searchParameters) {
	// case_sensitivity parameter setter
	return func(params *searchParameters) {
		params.caseSensitivity = case_sensitivity
	}
}

func WithMethod(method Method) func(*searchParameters) {
	// method parameter setter
	return func(params *searchParameters) {
		params.method = method
	}
}

func WithCount(count uint) func(*searchParameters) {
	// count parameter setter
	return func(params *searchParameters) {
		params.count = count
	}
}

type SearchResult struct {
	value []uint
}

func Search(inputString string, subString []string, options ...func(*searchParameters)) (result map[string]*SearchResult) {
	// Return an array of indexes of the found sub stings
	params := &searchParameters{
		inputString:     inputString,
		subString:       subString,
		caseSensitivity: false,
		method:          First,
		count:           1,
	}
	// Set optional parameters
	for _, option := range options {
		option(params)
	}

	length := len(inputString)

	var indexes = make(map[string]*SearchResult)

	for _, subStr := range params.subString {

		indexes[subStr] = nil

		patternTable := patternTable(subStr, params.caseSensitivity)

		subLength := len(subStr)
		strIndex, subStrIndex := 0, 0

		for strIndex < length {
			if params.caseSensitivity {
				if subStr[subStrIndex] == inputString[strIndex] {
					subStrIndex++
					strIndex++
				}
			} else {
				if strings.EqualFold(string(subStr[subStrIndex]), string(inputString[strIndex])) {
					subStrIndex++
					strIndex++
				}
			}

			if subStrIndex == subLength {
				if indexes[subStr] == nil {
					indexes[subStr] = &SearchResult{
						value: []uint{},
					}
				}
				indexes[subStr].value = append(indexes[subStr].value, uint(strIndex-subLength))
				subStrIndex = patternTable[subStrIndex-1]
			} else if strIndex < length && subStr[subStrIndex] != inputString[strIndex] {
				if subStrIndex != 0 {
					subStrIndex = patternTable[subStrIndex-1]
				} else {
					strIndex++
				}
			}
		}
	}

	return indexes
}

func patternTable(subString string, caseSensitivity bool) []int {
	// Create a pattern indexes table for Knuth–Morris–Pratt algorithm
	length := len(subString)
	patternTable := make([]int, length)
	prevIndex := 0
	index := 1

	for index < length {
		if caseSensitivity {
			if strings.EqualFold(string(subString[index]), string(subString[prevIndex])) {
				prevIndex++
				patternTable[index] = prevIndex
				index++
			} else {
				if prevIndex != 0 {
					prevIndex = patternTable[prevIndex-1]
				} else {
					patternTable[index] = 0
					index++
				}
			}
		} else {
			if subString[index] == subString[prevIndex] {
				prevIndex++
				patternTable[index] = prevIndex
				index++
			} else {
				if prevIndex != 0 {
					prevIndex = patternTable[prevIndex-1]
				} else {
					patternTable[index] = 0
					index++
				}
			}
		}
	}

	return patternTable
}
