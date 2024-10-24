package mysort

import (
	"reflect"
	"sort"
	"testing"
)

// Testing

// TEST_NUMBER = [
//     [],
//     [1],
//     [1, 2, 3, 4, 5],
//     [0, 0, 0, 55, 55, 60],
//     [9, 8, 7, 6, 5, 4, 3, 2, 1, 0],
//     [8, 0, 42, 3, 4, 8, 0, 45, 50, 9999, 7],
//     [-5, 0, 9, -999, 874, 35, -4, -5, 0],
//     [1, 1, 1],
// ]

var TEST_NUMBER = [][]interface{}{
	{},
	{1},
	{1, 2, 3, 4, 5},
	{0, 0, 0, 55, 55, 60},
	{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
	{8, 0, 42, 3, 4, 8, 0, 45, 50, 9999, 7},
	{-5, 0, 9, -999, 874, 35, -4, -5, 0},
	{1, 1, 1},
}

// TEST_STR = [
//     [],
//     ['a'],
//     ['a', 'b', 'c', 'd', 'e'],
//     ['aa', 'aa', 'aa', 'ab', 'ac', 'b'],
//     ['e', 'd', 'c', 'b', 'a'],
//     ['abc', 'a', 'foo', 'bar', 'booz', 'baz', 'spam', 'love'],
//     ['abc', 'abc', 'abc'],
//     [''],
// ]

var TEST_STR = [][]interface{}{
	{},
	{"a"},
	{"a", "b", "c", "d", "e"},
	{"aa", "aa", "aa", "ab", "ac", "b"},
	{"e", "d", "c", "b", "a"},
	{"abc", "a", "foo", "bar", "booz", "baz", "spam", "love"},
	{"abc", "abc", "abc"},
	{""},
}

func TestSearch(t *testing.T) {
	for _, dataSet := range TEST_NUMBER {
		t.Run("TEST_NUMBER_INCREASE", func(t *testing.T) {
			var got = make([]interface{}, len(dataSet))
			copy(got, dataSet)

			sortChan := make(chan []interface{})
			resultChan := make(chan []interface{})

			Sort(got, false, false, sortChan, resultChan)
			var want = make([]interface{}, len(dataSet))
			copy(want, dataSet)
			sort.Slice(want, func(i, j int) bool { return want[i].(int) < want[j].(int) })
			if !reflect.DeepEqual(got, want) {
				t.Errorf("\ngot: %+v \nwant: %+v\n", got, want)
			}
		})
	}

	for _, dataSet := range TEST_NUMBER {
		t.Run("TEST_NUMBER_DECREASE", func(t *testing.T) {
			var got = make([]interface{}, len(dataSet))
			copy(got, dataSet)

			sortChan := make(chan []interface{})
			resultChan := make(chan []interface{})

			Sort(got, true, false, sortChan, resultChan)
			var want = make([]interface{}, len(dataSet))
			copy(want, dataSet)
			sort.Slice(want, func(i, j int) bool { return want[i].(int) > want[j].(int) })
			if !reflect.DeepEqual(got, want) {
				t.Errorf("\ngot: %+v \nwant: %+v\n", got, want)
			}
		})
	}

	for _, dataSet := range TEST_STR {
		t.Run("TEST_STRING_INCREASE", func(t *testing.T) {
			var got = make([]interface{}, len(dataSet))
			copy(got, dataSet)

			sortChan := make(chan []interface{})
			resultChan := make(chan []interface{})

			Sort(got, false, false, sortChan, resultChan)
			var want = make([]interface{}, len(dataSet))
			copy(want, dataSet)
			sort.Slice(want, func(i, j int) bool { return want[i].(string) < want[j].(string) })
			if !reflect.DeepEqual(got, want) {
				t.Errorf("\ngot: %+v \nwant: %+v\n", got, want)
			}
		})
	}

	for _, dataSet := range TEST_STR {
		t.Run("TEST_STRING_DECREASE", func(t *testing.T) {
			var got = make([]interface{}, len(dataSet))
			copy(got, dataSet)

			sortChan := make(chan []interface{})
			resultChan := make(chan []interface{})

			Sort(got, true, false, sortChan, resultChan)
			var want = make([]interface{}, len(dataSet))
			copy(want, dataSet)
			sort.Slice(want, func(i, j int) bool { return want[i].(string) > want[j].(string) })
			if !reflect.DeepEqual(got, want) {
				t.Errorf("\ngot: %+v \nwant: %+v\n", got, want)
			}
		})
	}
}
