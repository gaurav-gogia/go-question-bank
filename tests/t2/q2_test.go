package t2

import (
	"fmt"
	"testing"

	tests ".."
	"../../questions/q2"
)

// URL constant
const URL = "https://0rkw2hfvbh.execute-api.us-east-2.amazonaws.com/fact_stage/fact"

func TestFact(t *testing.T) {
	var r, total, ans int
	var err error

	for i := 1; i <= 5; i++ {
		fmt.Println("Test #", i)

		r = tests.RandomNumber(1, 20)

		total = q2.Fact(r)
		ans, err = fact(r)

		if err != nil {
			t.Error(tests.CallError)
			break
		}

		if total != ans {
			t.Errorf(tests.TestError, total, ans)
			break
		}
	}
}

func fact(num int) (int, error) {
	values := map[string]int{"N": num}
	return tests.Call(values, URL)
}
