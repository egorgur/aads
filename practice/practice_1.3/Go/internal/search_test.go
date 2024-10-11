package search

// Tests for search modules

import (
	"reflect"
	"testing"
)

type DataSet struct {
	inputString     string
	subString       []string
	caseSensitivity bool
	method          string
	count           int
	output          map[string][]int
}

var TEST_SEARCH_ONE_SYMBOL = []*DataSet{ // Dataset for testing
	{inputString: "", subString: []string{"a"}, caseSensitivity: false, method: "first", count: 1, output: map[string][]int{"a": {}}},
	{inputString: "", subString: []string{"a"}, caseSensitivity: true, method: "first", count: 1, output: map[string][]int{"a": {}}},
	{inputString: "", subString: []string{"a"}, caseSensitivity: false, method: "last", count: 1, output: map[string][]int{"a": {}}},
	{inputString: "", subString: []string{"a"}, caseSensitivity: true, method: "last", count: 1, output: map[string][]int{"a": {}}},

	{inputString: "a", subString: []string{"a"}, caseSensitivity: false, method: "first", count: 1, output: map[string][]int{"a": {0}}},
	{inputString: "a", subString: []string{"a"}, caseSensitivity: true, method: "first", count: 1, output: map[string][]int{"a": {0}}},
	{inputString: "a", subString: []string{"a"}, caseSensitivity: false, method: "last", count: 1, output: map[string][]int{"a": {0}}},
	{inputString: "a", subString: []string{"a"}, caseSensitivity: true, method: "last", count: 1, output: map[string][]int{"a": {0}}},

	{inputString: "aaa", subString: []string{"a"}, caseSensitivity: false, method: "first", count: 3, output: map[string][]int{"a": {0, 1, 2}}},
	{inputString: "aaa", subString: []string{"a"}, caseSensitivity: true, method: "first", count: 3, output: map[string][]int{"a": {0, 1, 2}}},
	{inputString: "aaa", subString: []string{"a"}, caseSensitivity: false, method: "last", count: 3, output: map[string][]int{"a": {2, 1, 0}}},
	{inputString: "aaa", subString: []string{"a"}, caseSensitivity: true, method: "last", count: 3, output: map[string][]int{"a": {2, 1, 0}}},

	{inputString: "bca", subString: []string{"c"}, caseSensitivity: false, method: "first", count: 1, output: map[string][]int{"c": {1}}},
	{inputString: "bca", subString: []string{"c"}, caseSensitivity: true, method: "first", count: 1, output: map[string][]int{"c": {1}}},
	{inputString: "bca", subString: []string{"c"}, caseSensitivity: false, method: "last", count: 1, output: map[string][]int{"c": {1}}},
	{inputString: "bca", subString: []string{"c"}, caseSensitivity: true, method: "last", count: 1, output: map[string][]int{"c": {1}}},

	{inputString: "aaa", subString: []string{"a"}, caseSensitivity: false, method: "first", count: 2, output: map[string][]int{"a": {0, 1}}},
	{inputString: "aaa", subString: []string{"a"}, caseSensitivity: true, method: "first", count: 1, output: map[string][]int{"a": {0}}},
	{inputString: "aaa", subString: []string{"a"}, caseSensitivity: false, method: "last", count: 2, output: map[string][]int{"a": {2, 1}}},
	{inputString: "aaa", subString: []string{"a"}, caseSensitivity: true, method: "last", count: 10, output: map[string][]int{"a": {2, 1, 0}}},
}

var TEST_SEARCH_MANY_SYMBOL = []*DataSet{ // Dataset for testing
	{inputString: "", subString: []string{"abc"}, caseSensitivity: false, method: "first", count: 1, output: map[string][]int{"abc": {}}},
	{inputString: "", subString: []string{"abc"}, caseSensitivity: true, method: "first", count: 1, output: map[string][]int{"abc": {}}},
	{inputString: "", subString: []string{"abc"}, caseSensitivity: false, method: "last", count: 1, output: map[string][]int{"abc": {}}},
	{inputString: "", subString: []string{"abc"}, caseSensitivity: true, method: "last", count: 1, output: map[string][]int{"abc": {}}},

	{inputString: "a", subString: []string{"abc"}, caseSensitivity: false, method: "first", count: 1, output: map[string][]int{"abc": {}}},
	{inputString: "a", subString: []string{"abc"}, caseSensitivity: true, method: "first", count: 1, output: map[string][]int{"abc": {}}},
	{inputString: "a", subString: []string{"abc"}, caseSensitivity: false, method: "last", count: 1, output: map[string][]int{"abc": {}}},
	{inputString: "a", subString: []string{"abc"}, caseSensitivity: true, method: "last", count: 1, output: map[string][]int{"abc": {}}},

	{inputString: "abc", subString: []string{"abc"}, caseSensitivity: false, method: "first", count: 1, output: map[string][]int{"abc": {0}}},
	{inputString: "abc", subString: []string{"abc"}, caseSensitivity: true, method: "first", count: 1, output: map[string][]int{"abc": {0}}},
	{inputString: "abc", subString: []string{"abc"}, caseSensitivity: false, method: "last", count: 1, output: map[string][]int{"abc": {0}}},
	{inputString: "abc", subString: []string{"abc"}, caseSensitivity: true, method: "last", count: 1, output: map[string][]int{"abc": {0}}},

	{inputString: "abcabc", subString: []string{"abc"}, caseSensitivity: false, method: "first", count: 2, output: map[string][]int{"abc": {0, 3}}},
	{inputString: "abcabc", subString: []string{"abc"}, caseSensitivity: true, method: "first", count: 2, output: map[string][]int{"abc": {0, 3}}},
	{inputString: "abcabc", subString: []string{"abc"}, caseSensitivity: false, method: "last", count: 2, output: map[string][]int{"abc": {3, 0}}},
	{inputString: "abcabc", subString: []string{"abc"}, caseSensitivity: true, method: "last", count: 2, output: map[string][]int{"abc": {3, 0}}},

	{inputString: "aabcbccaabc", subString: []string{"abc"}, caseSensitivity: false, method: "first", count: 2, output: map[string][]int{"abc": {1, 8}}},
	{inputString: "aabcbccaabc", subString: []string{"abc"}, caseSensitivity: true, method: "first", count: 2, output: map[string][]int{"abc": {1, 8}}},
	{inputString: "aAbCbccaabc", subString: []string{"AbC"}, caseSensitivity: false, method: "first", count: 2, output: map[string][]int{"AbC": {1, 8}}},
	{inputString: "aabcbccaAbC", subString: []string{"AbC"}, caseSensitivity: true, method: "first", count: 1, output: map[string][]int{"AbC": {8}}},
	{inputString: "aabcbccaabc", subString: []string{"abc"}, caseSensitivity: false, method: "last", count: 2, output: map[string][]int{"abc": {8, 1}}},
	{inputString: "aabcbccaabc", subString: []string{"abc"}, caseSensitivity: true, method: "last", count: 2, output: map[string][]int{"abc": {8, 1}}},
	{inputString: "aAbCbccaabc", subString: []string{"AbC"}, caseSensitivity: false, method: "last", count: 2, output: map[string][]int{"AbC": {8, 1}}},
	{inputString: "aabcbccaAbC", subString: []string{"AbC"}, caseSensitivity: true, method: "last", count: 1, output: map[string][]int{"AbC": {8}}},

	{inputString: "abcabc", subString: []string{"abc"}, caseSensitivity: false, method: "first", count: 1, output: map[string][]int{"abc": {0}}},
	{inputString: "abcabcabc", subString: []string{"abc"}, caseSensitivity: true, method: "first", count: 2, output: map[string][]int{"abc": {0, 3}}},
	{inputString: "abcabc", subString: []string{"abc"}, caseSensitivity: false, method: "last", count: 9, output: map[string][]int{"abc": {3, 0}}},
	{inputString: "abcabc", subString: []string{"abc"}, caseSensitivity: true, method: "last", count: 2, output: map[string][]int{"abc": {3, 0}}},

}

var TEST_SEARCH_FEW_SUBSTR = []*DataSet{ // Dataset for testing
	{inputString: "", subString: []string{"abc","a"}, caseSensitivity: false, method: "first", count: 1, output: map[string][]int{"abc": {},"a": {}}},
	{inputString: "", subString: []string{"abc","a"}, caseSensitivity: true, method: "first", count: 1, output: map[string][]int{"abc": {},"a": {}}},
	{inputString: "", subString: []string{"abc","a"}, caseSensitivity: false, method: "last", count: 1, output: map[string][]int{"abc": {},"a": {}}},
	{inputString: "", subString: []string{"abc","a"}, caseSensitivity: true, method: "last", count: 1, output: map[string][]int{"abc": {},"a": {}}},

	{inputString: "a", subString: []string{"abc","a"}, caseSensitivity: false, method: "first", count: 1, output: map[string][]int{"abc": {},"a": {0}}},
	{inputString: "a", subString: []string{"abc","a"}, caseSensitivity: true, method: "first", count: 1, output: map[string][]int{"abc": {},"a": {0}}},
	{inputString: "a", subString: []string{"abc","a"}, caseSensitivity: false, method: "last", count: 1, output: map[string][]int{"abc": {},"a": {0}}},
	{inputString: "a", subString: []string{"abc","a"}, caseSensitivity: true, method: "last", count: 1, output: map[string][]int{"abc": {},"a": {0}}},

	{inputString: "ababbababa", subString: []string{"aba","bba"}, caseSensitivity: false, method: "first", count: 4, output: map[string][]int{"aba": {0, 5, 7},"bba": {3}}},
	{inputString: "ababbababa", subString: []string{"aba","bba"}, caseSensitivity: true, method: "first", count: 4, output: map[string][]int{"aba": {0, 5, 7},"bba": {3}}},
	{inputString: "ababbababa", subString: []string{"aba","bba"}, caseSensitivity: false, method: "last", count: 4, output: map[string][]int{"aba": {7, 5, 0},"bba": {3}}},
	{inputString: "ababbababa", subString: []string{"aba","bba"}, caseSensitivity: true, method: "last", count: 4, output: map[string][]int{"aba": {7, 5, 0},"bba": {3}}},

	{inputString: "ababbababa", subString: []string{"aba","bba"}, caseSensitivity: false, method: "first", count: 2, output: map[string][]int{"aba": {0, 5},"bba": {3}}},
	{inputString: "ababbababa", subString: []string{"aba","bba"}, caseSensitivity: true, method: "first", count: 1, output: map[string][]int{"aba": {0},"bba": {3}}},
	{inputString: "ababbababa", subString: []string{"aba","bba"}, caseSensitivity: false, method: "last", count: 1, output: map[string][]int{"aba": {7},"bba": {3}}},
	{inputString: "ababbababa", subString: []string{"aba","bba"}, caseSensitivity: true, method: "last", count: 10, output: map[string][]int{"aba": {7, 5, 0},"bba": {3}}},
}

func TestSearch(t *testing.T) {
	for _, dataSet := range TEST_SEARCH_ONE_SYMBOL {
		t.Run("TEST_SEARCH_ONE_SYMBOL", func(t *testing.T) {
			want := dataSet.output
			got := Search(
				dataSet.inputString,
				dataSet.subString,
				dataSet.caseSensitivity,
				dataSet.method,
				dataSet.count)

			for subStr := range want {
				if !reflect.DeepEqual(got[subStr], want[subStr]) {
					t.Errorf("\ngot: %+v \nwant: %+v\n", got, want)
				}
			}
		})
	}

	for _, dataSet := range TEST_SEARCH_MANY_SYMBOL {
		t.Run("TEST_SEARCH_MANY_SYMBOL", func(t *testing.T) {
			want := dataSet.output
			got := Search(
				dataSet.inputString,
				dataSet.subString,
				dataSet.caseSensitivity,
				dataSet.method,
				dataSet.count)

			for subStr := range want {
				if !reflect.DeepEqual(got[subStr], want[subStr]) {
					t.Errorf("\ngot: %+v \nwant: %+v\n", got, want)
				}
			}
		})
	}

	for _, dataSet := range TEST_SEARCH_FEW_SUBSTR {
		t.Run("TEST_SEARCH_FEW_SUBSTR", func(t *testing.T) {
			want := dataSet.output
			got := Search(
				dataSet.inputString,
				dataSet.subString,
				dataSet.caseSensitivity,
				dataSet.method,
				dataSet.count)

			for subStr := range want {
				if !reflect.DeepEqual(got[subStr], want[subStr]) {
					t.Errorf("\ngot: %+v \nwant: %+v\n", got, want)
				}
			}
		})
	}
}
