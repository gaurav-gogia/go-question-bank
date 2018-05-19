package t5

import (
	"fmt"
	"testing"

	"../../questions/q5"

	tests ".."
)

// URL Constants for this test
const (
	URLLow  = "https://g4hziytr7a.execute-api.us-east-2.amazonaws.com/low/low"
	URLHigh = "https://4bl2todicl.execute-api.us-east-2.amazonaws.com/high/high"
)

func TestMakeMap(t *testing.T) {
	fmt.Println("Starting Test")

	mockData := getMockData()
	total := q5.MakeMap(mockData)

	low, err := findLow(mockData)
	if err != nil {
		t.Error(tests.CallError)
	}

	high, err := findHigh(mockData)
	if err != nil {
		t.Error(tests.CallError)
	}

	if total["low"] != low {
		t.Errorf(tests.TestError, total["low"], low)
	}

	if total["high"] != high {
		t.Errorf(tests.TestError, total["high"], high)
	}
}

func getMockData() []int {
	limit := tests.RandomNumber(10, 100)
	var mockData []int

	for i := 0; i <= limit; i++ {
		num := tests.RandomNumber(0, 30000)
		mockData = append(mockData, num)
	}

	return mockData
}

func findLow(ints []int) (int, error) {
	values := map[string][]int{"INTS": ints}
	return tests.CallArray(values, URLLow)
}

func findHigh(ints []int) (int, error) {
	values := map[string][]int{"INTS": ints}
	return tests.CallArray(values, URLHigh)
}
