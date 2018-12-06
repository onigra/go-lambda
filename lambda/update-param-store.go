package main

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// GetNewManifestPath :
func GetNewManifestPath(service *ssm.SSM, keyname string) string {
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
func UpdateCurrentManifestPath(service *ssm.SSM, newManifestPath string, keyname string) bool {
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
func DeleteNewManifestParam(service *ssm.SSM, keyname string) bool {
	req := service.DeleteParameterRequest(&ssm.DeleteParameterInput{
		Name: &keyname,
	})

	_, err := req.Send()
	if err != nil {
		panic("failed to delete new parameter store, " + err.Error())
	}

	return true
}
