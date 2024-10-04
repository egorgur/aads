package search

// Tests for search modules

import (
	"strings"
	"testing"
)

type DataSet struct {
	inputString     string
	subString       []string
	caseSensitivity bool
	method          Method                   // A Method to search the sub strings
	count           uint                     // Optional
	output          map[string]*SearchResult // Awaited output - substring to indexes map
}

var TEST_SEARCH_ONE_SYMBOL = []interface{}{
	[]interface{}{"", "a", false, First, 1, map[string]*SearchResult{"a": nil}},
	[]interface{}{"", "a", true, First, 1, map[string]*SearchResult{"a": nil}},
	[]interface{}{"", "a", false, Last, 1, map[string]*SearchResult{"a": nil}},
	[]interface{}{"", "a", true, Last, 1, map[string]*SearchResult{"a": nil}},
}

func TestHello(t *testing.T) {
	// for inter, dataSlice := range TEST_SEARCH_ONE_SYMBOL{
	// 	t.Run("test 1", func(t *testing.T) {
			
	// 	})
	// }
}
