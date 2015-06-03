package golassert

import (
	"fmt"
	"reflect"
)

/*
AssertType asserts type of expected and result.
*/
func AssertType(expected interface{}, result interface{}) {
	expectedType := reflect.TypeOf(expected)
	resultType := reflect.TypeOf(result)
	if expectedType != resultType {
		err := "Error: [AssertEqual] Mismatched Types"
		err = fmt.Sprintf("%s\nExpected Value Type: %v\nResult: %v", err, expectedType, resultType)
		panic(err)
	}
}

/*
AssertEqual asserts if expected result is same as returned result.
*/
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

/*
AssertEqualStringArray asserts two string arrays.
*/
func AssertEqualStringArray(expected []string, result []string) {
	AssertType(expected, result)
	if expected == nil && result == nil {
		return
	}
	if len(expected) != len(result) {
		panic(fmt.Sprintf("Error: [] Different count of items\nExpected Value: %v\nResult: %v", expected, result))
	}
	for expectedIdx := range expected {
		elementExists := false
		for resultIdx := range result {
			if result[resultIdx] == expected[expectedIdx] {
				elementExists = true
			}
		}
		if !elementExists {
			panic(fmt.Sprintf("Error: [] Item missing: %v.\nExpected Value: %v\nResult: %v", expected[expectedIdx], expected, result))
		}
	}
}
