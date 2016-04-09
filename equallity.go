package conveyer

import (
	"fmt"
	"reflect"
)

// ShouldLookLike recursively examines non-zero fields and asserts
// their equality.
func ShouldLookLike(actual interface{}, args ...interface{}) string {
	valActual := reflect.ValueOf(actual)
	valExpected := reflect.ValueOf(args[0])

	typeActual := valActual.Type()
	typeExpected := valExpected.Type()

	numFieldsExpected := typeExpected.NumField()

	for i := 0; i < numFieldsExpected; i++ {
		structField := typeExpected.Field(i)
		expectedField := valExpected.Field(i)
		// If the field we are examining is the zero value, carry on
		if reflect.DeepEqual(reflect.Zero(expectedField.Type()), expectedField.Interface()) {
			continue
		}

		_, hasField := typeActual.FieldByName(structField.Name)
		actualField := valExpected.FieldByName(structField.Name)

		if !hasField {
			return fmt.Sprintf(`Type "%s" has no field named "%s" but was expected to`, typeActual.Name(), structField.Name)
		}

		if !reflect.DeepEqual(expectedField.Interface(), actualField.Interface()) {
			return Explain(`Field "%s" has the wrong value`, expectedField.Interface(), actualField.Interface(), structField.Name)
		}
	}

	return ""
}

// ShouldContainSomethingLike iterates over an array and requires that it has
// at least one element that matches via ShouldLookLike
func ShouldContainSomethingLike(actual interface{}, args ...interface{}) string {
	valArray := reflect.ValueOf(actual)

	for i := 0; i < valArray.Len(); i++ {
		// If we find a match return
		if ShouldLookLike(valArray.Index(i).Interface(), args[0]) == "" {
			return ""
		}
	}
	return Explain("Could not find anything that looked like the supplied value", args[0], actual)
}
