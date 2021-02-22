package main

import (
	"encoding/json"
	"fmt"
	"github.com/tmccombs/hcl2json/convert"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("hclToJSON", wrapper())
	<-c
}

func wrapper() js.Func {
	f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		input := args[0].String()
		jsonBytes, err := convert.Bytes([]byte(input), "", convert.Options{})
		if err != nil {
			return fmt.Errorf("convert to bytes: %w", err)
		}

		var v interface{}
		if err := json.Unmarshal(jsonBytes, &v); err != nil {
			return fmt.Errorf("unmarshal hcl2: %w", err)
		}

		pretty, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return fmt.Errorf("stringify hcl2: %w", err)
		}

		return string(pretty)
	})
	return f
}
