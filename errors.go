package conveyer

// ShouldHaveMessage asserts on the string returned by err.Error()
func ShouldHaveMessage(actual interface{}, args ...interface{}) string {
	msg := actual.(error).Error()
	expected := args[0].(string)

	if msg != expected {
		return Explain("Error did not have expected message", expected, msg)
	}

	return ""
}
