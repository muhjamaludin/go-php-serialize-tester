package main

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

// Test to check function serialize go to php for number
func TestWriteSerializeNumber(t *testing.T) {
	// write serialize go to json file
	num := 12

	// declare json object serialize
	data := map[string]interface{}{
		"code":   "GoSerializeNumber",
		"base":   num,
		"result": WriteSerializeNumber(num),
	}

	// write in the file
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile("./result/goSerializeNum.json", file, 0644)

	// read phpSerializeNumber json
	jsonFile, err := os.ReadFile("./result/phpSerializeNum.json")
	if err != nil {
		log.Fatal(err)
	}

	// type struct serialize
	type Num struct {
		Base   int    `json:"base"`
		Code   string `json:"code"`
		Result string `json:"result"`
	}

	// get base data
	var nums Num
	json.Unmarshal(jsonFile, &nums)

	// compare unserialize go to serialize php
	if ReadSerializeNumber(nums.Result) != nums.Base {
		t.Errorf("Wrong! Serialize number must be: %.2v", nums.Base)
	}
}

// Test to check function serialize go to php for number
func TestSerializeString(t *testing.T) {
	// test assert
	str := "test string go complex 12345678"

	// declare json object serialize
	data := map[string]interface{}{
		"code":   "GoSerializeNumber",
		"base":   str,
		"result": WriteSerializeString(str),
	}

	// write to file json
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile("./result/goSerializeString.json", file, 0644)

	// read phpSerializeString json
	jsonFile, err := os.ReadFile("./result/phpSerializeString.json")
	if err != nil {
		log.Fatal(err)
	}

	// type struct serialize
	type Str struct {
		Base   string `json:"base"`
		Code   string `json:"code"`
		Result string `json:"result"`
	}

	// get base data
	var strs Str
	json.Unmarshal(jsonFile, &strs)

	// compare unserialize go to serialize php
	if ReadSerializeString(strs.Result) != strs.Base {
		t.Errorf("Wrong! Serialize string must be: %.2v", strs.Base)
	}
}
