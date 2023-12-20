package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)

	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		// Move elements of arr[0..i-1] that are greater than key to one position ahead of their current position
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	// Example usage
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Unsorted array:", arr)

	insertionSort(arr)

	fmt.Println("Sorted array:", arr)
}
