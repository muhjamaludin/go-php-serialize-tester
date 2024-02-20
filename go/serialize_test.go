package main

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
	"testing"
)

// Test to check function serialize go to php for number
func TestSerializeNumber(t *testing.T) {
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
		"code":   "GoSerializeString",
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

// Test to check function serialize go to php for number
func TestSerializeSliceString(t *testing.T) {
	// test assert
	sliceStr := []string{"abc", "1Dor", "", "9_000_000"}

	// declare json object serialize
	data := map[string]interface{}{
		"code":   "GoSerializeSliceString",
		"base":   sliceStr,
		"result": WriteSerializeSliceString(sliceStr),
	}

	// write to file json
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile("./result/goSerializeSliceString.json", file, 0644)

	// read phpSerializeString json
	jsonFile, err := os.ReadFile("./result/phpSerializeArrayString.json")
	if err != nil {
		log.Fatal(err)
	}

	// type struct serialize
	type SliceStr struct {
		Base   []string `json:"base"`
		Code   string   `json:"code"`
		Result string   `json:"result"`
	}

	// get base data
	var strs SliceStr
	json.Unmarshal(jsonFile, &strs)

	// compare unserialize go to serialize php
	if reflect.DeepEqual(ReadSerializeSliceString(strs.Result), strs.Base) {
		t.Errorf("Wrong! Serialize string must be: %.2v", strs.Base)
	}
}

func TestSerializeSliceInteger(t *testing.T) {
	// test assert
	sliceInt := []int{200, 0, 1, 9_000_000, 1000}

	// declare json object serialize
	data := map[string]interface{}{
		"code":   "GoSerializeSliceNumber",
		"base":   sliceInt,
		"result": WriteSerializeSliceInteger(sliceInt),
	}

	// write to file json
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile("./result/goSerializeSliceInt.json", file, 0644)

	// read phpSerializeString json
	jsonFile, err := os.ReadFile("./result/phpSerializeArrayInteger.json")
	if err != nil {
		log.Fatal(err)
	}

	// type struct serialize
	type SliceStr struct {
		Base   []string `json:"base"`
		Code   string   `json:"code"`
		Result string   `json:"result"`
	}

	// get base data
	var strs SliceStr
	json.Unmarshal(jsonFile, &strs)

	// compare unserialize go to serialize php
	if reflect.DeepEqual(ReadSerializeSliceInt(strs.Result), strs.Base) {
		t.Errorf("Wrong! Serialize string must be: %.2v", strs.Base)
	}
}

func TestSerializeSliceStringInteger(t *testing.T) {
	// test assert
	sliceStrInt := []interface{}{0, 1, 2_000_000, "luca", "gunung Tambora", 7}

	// declare json object serialize
	data := map[string]interface{}{
		"code":   "GoSerializeSliceStringNumber",
		"base":   sliceStrInt,
		"result": WriteSerializeSliceStringIntegerBoolean(sliceStrInt),
	}

	// write to file json
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile("./result/goSerializeSliceStringInteger.json", file, 0644)

	// read phpSerializeString json
	jsonFile, err := os.ReadFile("./result/phpSerializeArrayStringInteger.json")
	if err != nil {
		log.Fatal(err)
	}

	// type struct serialize
	type SliceStr struct {
		Base   []string `json:"base"`
		Code   string   `json:"code"`
		Result string   `json:"result"`
	}

	// get base data
	var strs SliceStr
	json.Unmarshal(jsonFile, &strs)

	// compare unserialize go to serialize php
	if reflect.DeepEqual(ReadSerializeSliceStringIntBoolean(strs.Result), strs.Base) {
		t.Errorf("Wrong! Serialize string must be: %.2v", strs.Base)
	}
}

func TestSerializeSliceStringIntegerBoolean(t *testing.T) {
	// test assert
	sliceStrIntBool := []interface{}{true, 0, 1, 9_000_000, "dora", "eminem", false, 100}

	// declare json object serialize
	data := map[string]interface{}{
		"code":   "GoSerializeSliceStringNumberBool",
		"base":   sliceStrIntBool,
		"result": WriteSerializeSliceStringIntegerBoolean(sliceStrIntBool),
	}

	// write to file json
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile("./result/goSerializeSliceStringIntegerBoolean.json", file, 0644)

	// read phpSerializeString json
	jsonFile, err := os.ReadFile("./result/phpSerializeArrayStringIntegerBoolean.json")
	if err != nil {
		log.Fatal(err)
	}

	// type struct serialize
	type SliceStr struct {
		Base   []string `json:"base"`
		Code   string   `json:"code"`
		Result string   `json:"result"`
	}

	// get base data
	var strs SliceStr
	json.Unmarshal(jsonFile, &strs)

	// compare unserialize go to serialize php
	if reflect.DeepEqual(ReadSerializeSliceStringIntBoolean(strs.Result), strs.Base) {
		t.Errorf("Wrong! Serialize string must be: %.2v", strs.Base)
	}
}
