package main

import (
	"fmt"
	"github.com/oshank-kumar/testrepo/Hello/mathUtil"
)

func arrayOperation() {
	var len int
	fmt.Print("Enter Array size :")
	fmt.Scanln(&len)
	arr := make([]float64, len)
	fmt.Print("Enter Elements : ")
	for i, _ := range arr {
		fmt.Scan(&arr[i])
	}
	avg := mathUtil.Average(arr)
	fmt.Println("Average of", arr, "is", avg)
}
