package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// ParseDetail :
func ParseDetail(event events.CloudWatchEvent) map[string]interface{} {
	var detail map[string]interface{}

	err := json.Unmarshal(event.Detail, &detail)
	if err != nil {
		panic("detail Unmarshal error.")
	}

	return detail
}

// Handler :
func Handler(event events.CloudWatchEvent) (string, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	cfg.Region = endpoints.ApNortheast1RegionID
	service := ssm.New(cfg)

	detail := ParseDetail(event)
	deployParams := detail["parameters"].(map[string]interface{})

	if detail["status"] == "Success" {
		newManifestPath := GetNewManifestPath(service, deployParams["newParam"].(string))
		UpdateCurrentManifestPath(service, newManifestPath, deployParams["currentParam"].(string))
	}

	DeleteNewManifestParam(service, deployParams["newParam"].(string))

	return "ok", nil
}

func main() {
	lambda.Start(Handler)
}
