package main

import (
	"encoding/json"
	"fmt"
)

// Event :
type Event struct {
	Detail json.RawMessage `json:"detail"`
}

func main() {
	jsonStr := `
{
  "detail": {
		"foo": "bar"
	}
}
`
	jsonBytes := ([]byte)(jsonStr)
	event := new(Event)

	if err := json.Unmarshal(jsonBytes, &event); err != nil {
		panic("error1")
	}

	var detail map[string]interface{}
	err := json.Unmarshal(event.Detail, &detail)
	if err != nil {
		panic("error2")
	}

	fmt.Println(detail)
	fmt.Println(detail["foo"])
}
