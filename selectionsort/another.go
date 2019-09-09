package main

import "fmt"

func main() {
	arr := []int{64, 25, 12, 22, 11}
	selectionSort(arr)
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
}

func selectionSort(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		largest := i
		for j := i - 1; j >= 0; j-- {
			if arr[largest] < arr[j] {
				largest = j
			}
		}
		if largest != i {
			tmp := arr[i]
			arr[i] = arr[largest]
			arr[largest] = tmp
		}
	}
}
