package t3

import (
	"testing"

	"../../questions/q3"
)

func TestFail(t *testing.T) {
	mockData := getMockData()
	mockResult := getMockResult()
	result := q3.Fail(mockData)

	if len(result) != len(mockResult) {
		t.Error("Wrong list of failed students")
	}

	for _, n := range result {
		if ok := equalityTest(n.Marks, mockResult); !ok {
			t.Error("Wrong list of failed students")
		}
	}
}

func getMockData() []q3.Student {
	mockData := []q3.Student{
		{
			Name:  "Ravi",
			Roll:  1,
			Marks: 84,
		},
		{
			Name:  "Jai",
			Roll:  30,
			Marks: 28,
		},
		{
			Name:  "Mehak",
			Roll:  22,
			Marks: 98,
		},
		{
			Name:  "Surbhi",
			Roll:  3,
			Marks: 22,
		},
		{
			Name:  "Mohit",
			Roll:  2,
			Marks: 54,
		},
		{
			Name:  "Akansha",
			Roll:  7,
			Marks: 41,
		},
	}

	return mockData
}

func getMockResult() []q3.Student {
	mockResult := []q3.Student{
		{
			Name:  "Surbhi",
			Roll:  3,
			Marks: 22,
		},
		{
			Name:  "Jai",
			Roll:  30,
			Marks: 28,
		},
	}

	return mockResult
}

func equalityTest(mark int, mockResult []q3.Student) bool {
	for _, m := range mockResult {
		if mark == m.Marks {
			return true
		}
	}

	return false
}
