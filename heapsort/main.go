package main

import "fmt"

func heaptify(arr []int, n, root int) {
	largest := root
	left := 2*root + 1
	right := 2*root + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != root {
		swap := arr[root]
		arr[root] = arr[largest]
		arr[largest] = swap

		heaptify(arr, n, largest)
	}
}

func heapsort(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heaptify(arr, len(arr), i)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		swap := arr[i]
		arr[i] = arr[0]
		arr[0] = swap
		heaptify(arr, i, 0)
	}
}

func main() {
	arr := []int{15, 11, 13, 5, 3, 2, 8, 7, 6, 20}
	heapsort(arr)
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
}
