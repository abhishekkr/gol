package golassert

import (
	"fmt"
	"reflect"
)

func AssertType(expected interface{}, result interface{}) {
	expected_type := reflect.TypeOf(expected)
	result_type := reflect.TypeOf(result)
	if expected_type != result_type {
		err := "Error: [AssertEqual] Mismatched Types"
		err = fmt.Sprintf("%s\nExpected Value Type: %v\nResult: %v", err, expected_type, result_type)
		panic(err)
	}
}

// assert if expected result is same as returned result
func AssertEqual(expected interface{}, result interface{}) {
	AssertType(expected, result)
	if expected == nil && result == nil {
		return
	}
	switch result.(type) {
	case string, int, error:
		if expected != result {
			panic(fmt.Sprintf("Error: [] Mismatched Values\nExpected value: %v\nResult: %v", expected, result))
		}
	default:
		panic("Error: AssertEqual doesn't handles this type yet.")
	}

}
