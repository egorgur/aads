package mysort

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//Sorting functions

// Минимальный размер подмассива для сортировки вставками
const minRun = 32

// printArray выводит массив в консоль
func printArray(arr []interface{}) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

// printStep выводит текущий шаг сортировки и массив
func printStep(stepName string, arr []interface{}) {
	fmt.Printf("%s: %s\n", stepName, strings.Repeat("=", 10))
	printArray(arr)
	time.Sleep(500 * time.Millisecond) // Задержка для визуализации
}

// insertionSort выполняет сортировку вставками
func insertionSort(arr []interface{}, left, right int, less func(a, b interface{}) bool, visualize bool) {
	for i := left + 1; i <= right; i++ {
		key := arr[i]
		j := i - 1
		for j >= left && less(key, arr[j]) {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
		if visualize {
			printStep("Insertion Sort", arr)
		}
	}
}

// merge выполняет слияние двух отсортированных подмассивов
func merge(arr []interface{}, l, m, r int, less func(a, b interface{}) bool, visualize bool) {
	n1, n2 := m-l+1, r-m
	left := make([]interface{}, n1)
	right := make([]interface{}, n2)

	copy(left, arr[l:l+n1])
	copy(right, arr[m+1:m+1+n2])

	i, j, k := 0, 0, l
	for i < n1 && j < n2 {
		if less(left[i], right[j]) {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
		if visualize {
			printStep("Merge", arr) // Визуализация процесса слияния
		}
	}
	for i < n1 {
		arr[k] = left[i]
		i++
		k++
		if visualize {
			printStep("Merge", arr)
		}
	}
	for j < n2 {
		arr[k] = right[j]
		j++
		k++
		if visualize {
			printStep("Merge", arr)
		}
	}
}

// TimSort выполняет сортировку Timsort с визуализацией
func TimSort(arr []interface{}, less func(a, b interface{}) bool, visualize bool) {
	n := len(arr)

	// Сортировка подмассивов вставками
	for i := 0; i < n; i += minRun {
		end := i + minRun - 1
		if end >= n {
			end = n - 1
		}
		insertionSort(arr, i, end, less, visualize)
	}

	// Слияние отсортированных подмассивов
	for size := minRun; size < n; size *= 2 {
		for left := 0; left < n; left += 2 * size {
			mid := left + size - 1
			if mid >= n-1 {
				break
			}
			right := left + 2*size - 1
			if right >= n {
				right = n - 1
			}
			merge(arr, left, mid, right, less, visualize)
		}
	}
}

// TimSort function wrapper
func Sort(input_arr []interface{}, reverse bool, visualize bool) ([]interface{}, error) {

	var containsInts bool

	containsInts = onlyInts(input_arr) // If only integers are present

	var containsStrings bool = false

	if !containsInts {
		containsStrings = onlyStrings(input_arr) // If only strings are present
	}

	var less func(a, b interface{}) bool

	if containsInts {
		if reverse {
			less = func(a, b interface{}) bool {
				return a.(int) > b.(int)
			}
		} else {
			less = func(a, b interface{}) bool {
				return a.(int) < b.(int)
			}
		}

		TimSort(input_arr, less, visualize)
		return input_arr, nil
	}

	if containsStrings {
		if reverse {
			less = func(a, b interface{}) bool {
				return a.(string) > b.(string)
			}
		} else {
			less = func(a, b interface{}) bool {
				return a.(string) < b.(string)
			}
		}

		TimSort(input_arr, less, visualize)
		return input_arr, nil
	}

	// If there are different element's types in the array or all elements are not sortable
	return nil, errors.New("Not sortable")
}

func onlyInts(input_array []interface{}) bool {
	for _, v := range input_array {
		_, ok := v.(int)
		if !ok {
			return false
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
