package tests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Constants
const (
	ApplicationType = "application/json"
	CallError       = "Could NOT complete test"
	TestError       = "INCORRECT, got: %v, wanted: %v."
)

// RandomNumber function for test cases
func RandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

// Call function to make the call
func Call(values map[string]int, url string) (int, error) {
	json, _ := json.Marshal(values)
	return call(json, url)
}

// CallArray function to make the call using an array as input
func CallArray(values map[string][]int, url string) (int, error) {
	json, _ := json.Marshal(values)
	return call(json, url)
}

func call(json []byte, url string) (int, error) {
	res, err := http.Post(url, ApplicationType, bytes.NewBuffer(json))
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return strconv.Atoi(string(body))
}
