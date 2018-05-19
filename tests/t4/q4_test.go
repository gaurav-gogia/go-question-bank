package t4

import (
	"testing"

	"../../questions/q4"
)

func TestGetCandidates(t *testing.T) {
	mockData := getMockData()
	mockSubjects := getMockSubjects()
	mockResult := getMockResult()
	result := q4.GetCandidates(mockData, mockSubjects)

	if len(result) != len(mockResult) {
		t.Error("Wrong list of candidates")
	}

	for _, n := range result {
		if ok := equalityTest(n.Roll, n.Name, mockResult); !ok {
			t.Error("Wrong list of candidates")
		}
	}
}

func getMockData() q4.DB {
	Compiler := []q4.Candidate{
		{
			Name: "User4",
			Roll: 4,
		},
		{
			Name: "User2",
			Roll: 2,
		},
		{
			Name: "User3",
			Roll: 3,
		},
	}

	NPMC := []q4.Candidate{
		{
			Name: "User2",
			Roll: 2,
		},
		{
			Name: "User3",
			Roll: 3,
		},
		{
			Name: "User4",
			Roll: 4,
		},
	}

	AI := []q4.Candidate{
		{
			Name: "User2",
			Roll: 2,
		},
		{
			Name: "User4",
			Roll: 4,
		},
	}

	subjects := []q4.Subject{
		{SubjectName: "Compiler", Student: Compiler},
		{SubjectName: "NPMC", Student: NPMC},
		{SubjectName: "AI", Student: AI},
	}

	mockData := q4.DB{
		SubjectInfo: subjects,
	}

	return mockData
}

func getMockSubjects() []string {
	subjects := []string{
		"AI",
	}

	return subjects
}

func getMockResult() []q4.Candidate {
	mockResult := []q4.Candidate{
		{
			Name: "User2",
			Roll: 2,
		},
		{
			Name: "User4",
			Roll: 4,
		},
	}

	return mockResult
}

func equalityTest(roll int, name string, mockResult []q4.Candidate) bool {
	for _, m := range mockResult {
		if name == m.Name && roll == m.Roll {
			return true
		}
	}

	return false
}
