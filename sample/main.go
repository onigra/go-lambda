package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler :
func Handler(event events.CloudWatchEvent) (string, error) {
	var detail map[string]interface{}
	err := json.Unmarshal(event.Detail, &detail)
	if err != nil {
		panic("Json Unmarshal error.")
	}

	command := detail["parameters"].(map[string]interface{})["commands"]
	// fmt.Printf("%#v", detail)
	fmt.Println(command.([]interface{})[0])

	return "ok", nil
}

func main() {
	lambda.Start(Handler)
}
