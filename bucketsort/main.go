package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []float64{0.897, 0.565, 0.656, 0.1234, 0.665, 0.3434}
	bucketsort(arr)
}

func bucketsort(arr []float64) {
	var result [][]float64
	for i := 0; i < len(arr); i++ {
		item := make([]float64, 0, len(arr))
		result = append(result, item)
	}

	for _, v := range arr {
		index := int(v * float64(len(arr)))
		result[index] = append(result[index], v)
	}

	for i := 0; i < len(result); i++ {
		sort.Float64s(result[i])
	}

	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			fmt.Printf("%v ", result[i][j])
		}
	}

}
