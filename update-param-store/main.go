package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// GetNewManifestPath :
func GetNewManifestPath(service *ssm.SSM) string {
	keyname := "/app/deploy/new"

	req := service.GetParameterRequest(&ssm.GetParameterInput{
		Name: &keyname,
	})

	resp, err := req.Send()
	if err != nil {
		panic("failed to describe new parameter store, " + err.Error())
	}

	return *resp.Parameter.Value
}

// UpdateCurrentManifestPath :
func UpdateCurrentManifestPath(service *ssm.SSM, newManifestPath string) bool {
	keyname := "/app/deploy/current"
	overwrite := true

	req := service.PutParameterRequest(&ssm.PutParameterInput{
		Name:      &keyname,
		Value:     &newManifestPath,
		Type:      ssm.ParameterTypeString,
		Overwrite: &overwrite,
	})

	_, err := req.Send()
	if err != nil {
		panic("failed to update current parameter store, " + err.Error())
	}

	return true
}

// DeleteNewManifestParam :
func DeleteNewManifestParam(service *ssm.SSM) bool {
	keyname := "/app/deploy/new"

	req := service.DeleteParameterRequest(&ssm.DeleteParameterInput{
		Name: &keyname,
	})

	_, err := req.Send()
	if err != nil {
		panic("failed to delete new parameter store, " + err.Error())
	}

	return true
}

func main() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	cfg.Region = endpoints.ApNortheast1RegionID
	svc := ssm.New(cfg)

	newManifestPath := GetNewManifestPath(svc)
	UpdateCurrentManifestPath(svc, newManifestPath)
	DeleteNewManifestParam(svc)

	fmt.Println("New manifest path: ", newManifestPath)
}
