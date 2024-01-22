package main

import (
	"encoding/json"
	"os"
	"testing"
)

// Test to check function serialize go to php for number
func TestWriteSerializeNumber(t *testing.T) {
	// test assert
	num := 12
	t.Log("Serialize number : ", WriteSerializeNumber(num))

	if WriteSerializeNumber(num) != "i:12;" {
		t.Errorf("Wrong! Serialize number must be: %.2vv", "i:12;")
	}

	// insert to json to check using php
	// declare array of object
	var data []map[string]interface{}

	// declare json object serialize
	serializeData := map[string]interface{}{
		"code":   "GoSerializeNumber",
		"base":   num,
		"result": WriteSerializeNumber(num),
	}

	data = append(data, serializeData)

	// write to file json
	file, _ := json.Marshal(data)
	_ = os.WriteFile("test.json", file, 0644)
}

// Test to check function serialize go to php for number
func TestSerializeString(t *testing.T) {
	// test assert
	str := "test string complex 12345678"
	t.Log("Serialize string : ", WriteSerializeString(str))

	if WriteSerializeString(str) != "s:26:\"test string complex 12345678\";" {
		t.Errorf("Wrong! Serialize string must be: %.2v", "s:26:\"testing bigbox bigenvelope\";")
	}

	// insert to json to check using php
	// declare array of object
	var data []map[string]interface{}

	// declare json object serialize
	serializeData := map[string]interface{}{
		"code":   "GoSerializeNumber",
		"base":   str,
		"result": WriteSerializeString(str),
	}

	data = append(data, serializeData)

	// write to file json
	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile("test.json", file, 0644)
}
