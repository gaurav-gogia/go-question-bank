package t1

import (
	"fmt"
	"testing"

	tests ".."
	"../../questions/q1"
)

// URL constant
const URL = "https://kf78kd16f8.execute-api.us-east-2.amazonaws.com/stage_1/sumFun"

func TestSum(t *testing.T) {
	var r1, r2, total, ans int
	var err error

	for i := 1; i <= 5; i++ {
		fmt.Println("Test #", i)

		r1 = tests.RandomNumber(0, 10000)
		r2 = tests.RandomNumber(0, 10000)

		total = q1.Sum(r1, r2)
		ans, err = sum(r1, r2)

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

func sum(x, y int) (int, error) {
	values := map[string]int{"N1": x, "N2": y}
	return tests.Call(values, URL)
}
