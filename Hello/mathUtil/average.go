package mathUtil

import (
	_ "fmt"
)

func Average(arr []float64) (avg float64) {
	sum := 0.0
	for _, val := range arr {
		sum += val
	}
	avg = sum / float64(len(arr))
	return
}
