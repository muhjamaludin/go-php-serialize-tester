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

// // function write string to php serialize format
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
