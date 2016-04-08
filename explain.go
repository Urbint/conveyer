package conveyer

import (
	"encoding/json"
	"fmt"
	"github.com/luci/go-render/render"
)

// Explain returns a structured error message for displaying in the goconvey ui
// it is a glorified reimplementation of assertions.serializer
func Explain(message string, expected interface{}, actual interface{}, formatArgs ...interface{}) string {
	view := failureView{
		Message:  fmt.Sprintf(message, formatArgs...),
		Expected: render.Render(expected),
		Actual:   render.Render(actual),
	}
	serialized, err := json.Marshal(view)
	if err != nil {
		return message
	}
	return string(serialized)
}

type failureView struct {
	Message  string `json:"Message"`
	Expected string `json:"Expected"`
	Actual   string `json:"Actual"`
}
