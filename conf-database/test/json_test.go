package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {

	var js = `{ "a":1, "b":1, "c":[{"a":22}] }`

	var mp map[string]interface{}

	json.Unmarshal([]byte(js), &mp)
	fmt.Println(mp["c"].([]interface{}))
}
