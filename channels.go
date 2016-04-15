package conveyer

import "reflect"

// ShouldHaveReceivedSomethingLike reads a channel and requires that a record that looks like
// the specified interface be emitted from the channel. The channel will be fully drained
// before returning
func ShouldHaveReceivedSomethingLike(actual interface{}, args ...interface{}) string {
	channel := reflect.ValueOf(actual)

	results := []interface{}{}

	for {
		rec, isOk := channel.Recv()
		if !isOk {
			break
		}
		results = append(results, rec.Interface())
	}

	if ShouldContainSomethingLike(results, args...) == "" {
		return ""
	}

	return Explain(`Channel never emitted the desired value`, args[0], results)
}
