package main

import "fmt"

func main() {
	arr := []int{10, 80, 70, 90, 40, 50, 70}
	quicksort(arr, 0, len(arr)-1)

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
}

func quicksort(arr []int, low, high int) {
	if low < high {
		base := partition(arr, low, high)
		quicksort(arr, 0, base-1)
		quicksort(arr, base+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	for low < high {
		for low < high && arr[low] <= pivot {
			low++
		}
		arr[high] = arr[low]
		for low < high && arr[high] >= pivot {
			high--
		}
		arr[low] = arr[high]
	}
	arr[high] = pivot
	return high
}
