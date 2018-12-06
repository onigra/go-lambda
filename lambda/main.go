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

// ParseDeployParams :
func ParseDeployParams(params string) map[string]interface{} {
	deployParamsBytes := ([]byte)(params)
	var deployParams map[string]interface{}

	err := json.Unmarshal(deployParamsBytes, &deployParams)
	if err != nil {
		panic("deployParams Unmarshal error.")
	}

	return deployParams
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
	deployParams := ParseDeployParams(detail["parameters"].(string))

	if detail["status"] == "Success" {
		newManifestPath := GetNewManifestPath(service, deployParams["newParam"].([]interface{})[0].(string))
		UpdateCurrentManifestPath(service, newManifestPath, deployParams["currentParam"].([]interface{})[0].(string))
	}
	DeleteNewManifestParam(service, deployParams["newParam"].([]interface{})[0].(string))

	return "ok", nil
}

func main() {
	lambda.Start(Handler)
}
