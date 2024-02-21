package main

import (
	"github.com/elliotchance/phpserialize"
)

// function write number to php serialize format
func WriteSerializeNumber(n int) string {
	out, err := phpserialize.Marshal(n, nil)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func ReadSerializeNumber(serialize string) int {
	var result int
	err := phpserialize.Unmarshal([]byte(serialize), &result)
	if err != nil {
		panic(err)
	}

	return result
}

// function write string to php serialize format
func WriteSerializeString(str string) string {
	out, err := phpserialize.Marshal(str, nil)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func ReadSerializeString(serialize string) string {
	var result string
	err := phpserialize.Unmarshal([]byte(serialize), &result)
	if err != nil {
		panic(err)
	}

	return result
}

// function write slice string to php serialize format
func WriteSerializeSliceString(sliceStr []string) string {
	out, err := phpserialize.Marshal(sliceStr, nil)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func ReadSerializeSliceString(serialize string) []interface{} {
	var result []interface{}
	err := phpserialize.Unmarshal([]byte(serialize), &result)
	if err != nil {
		panic(err)
	}

	return result
}

// function writer slice integer to php serialize format
func WriteSerializeSliceInteger(sliceInt []int) string {
	out, err := phpserialize.Marshal(sliceInt, nil)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func ReadSerializeSliceInt(serialize string) []interface{} {
	var result []interface{}
	err := phpserialize.Unmarshal([]byte(serialize), &result)
	if err != nil {
		panic(err)
	}

	return result
}

// function writer slice integer to php serialize format
func WriteSerializeSliceStringIntegerBoolean(sliceInt []interface{}) string {
	out, err := phpserialize.Marshal(sliceInt, nil)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func ReadSerializeSliceStringIntBoolean(serialize string) []interface{} {
	var result []interface{}
	err := phpserialize.Unmarshal([]byte(serialize), &result)
	if err != nil {
		panic(err)
	}

	return result
}

// function writer struct to php serialize format
func WriteSerializeMapInterface(goStruct map[string]interface{}) string {
	out, err := phpserialize.Marshal(goStruct, nil)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func ReadSerializeMapInterface(serialize string) map[interface{}]interface{} {
	var result map[interface{}]interface{}
	err := phpserialize.Unmarshal([]byte(serialize), &result)
	if err != nil {
		panic(err)
	}

	return result
}
