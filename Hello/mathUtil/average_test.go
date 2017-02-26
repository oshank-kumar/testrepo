package mathUtil_test

import (
	. "github.com/oshank-kumar/testrepo/Hello/mathUtil"
	"testing"
)

type TestResult struct {
	slice  []float64
	result float64
}

func TestAverage(t *testing.T) {
	myResult := []TestResult{
		{[]float64{12, 10, 15, 4.5, 8.5}, 10},
		{[]float64{5, 6, 9, 10}, 7.5},
		{[]float64{8, 10, 9, 5}, 8},
	}
	for _, val := range myResult {
		if tempAvg := Average(val.slice); tempAvg != val.result {
			t.Log("Expected", val.result, "got", tempAvg)
			t.Fail()
		}
	}
}
