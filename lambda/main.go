package main

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler :
func Handler() (string, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	cfg.Region = endpoints.ApNortheast1RegionID
	service := ssm.New(cfg)

	result := UpdateParamStore(service)
	return result, nil
}

func main() {
	lambda.Start(Handler)
}
