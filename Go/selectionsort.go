package main

import "fmt"

func selectionSort(arr []int) {
    n := len(arr)

    for i := 0; i < n-1; i++ {
        // Find the minimum element in the unsorted part of the array
        minIndex := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }

        // Swap the found minimum element with the first element
        arr[i], arr[minIndex] = arr[minIndex], arr[i]
    }
}

func main() {
    // Example usage
    arr := []int{64, 25, 12, 22, 11}
    fmt.Println("Unsorted array:", arr)

    selectionSort(arr)

    fmt.Println("Sorted array:", arr)
}
