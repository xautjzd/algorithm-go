package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	bubbleSort(arr)
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
}

func bubbleSort(arr []int) {
	n := len(arr)
	swapped := false
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}
