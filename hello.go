package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func main() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	cfg.Region = endpoints.ApNortheast1RegionID

	svc := ssm.New(cfg)
	keyname := "/foo/bar"
	decryption := false

	req := svc.GetParameterRequest(&ssm.GetParameterInput{
		Name: &keyname,
		WithDecryption: &decryption,
	})

	resp, err := req.Send()
	if err != nil {
		panic("failed to describe table, "+err.Error())
	}

	fmt.Println("Response", resp)
}
