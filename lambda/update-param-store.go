package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// GetNewManifestPath :
func GetNewManifestPath(service *ssm.SSM) (string, error){
	keyname := "/app/deploy/new"

	req := service.GetParameterRequest(&ssm.GetParameterInput{
		Name: &keyname,
	})

	resp, err := req.Send()
	return *resp.Parameter.Value, err
}

// UpdateCurrentManifestPath :
func UpdateCurrentManifestPath(service *ssm.SSM, newManifestPath string) (int64, error){
	keyname := "/app/deploy/current"
	overwrite := true

	req := service.PutParameterRequest(&ssm.PutParameterInput{
		Name: &keyname,
		Value: &newManifestPath,
		Type: ssm.ParameterTypeString,
		Overwrite: &overwrite,
	})

	resp, err := req.Send()
	return *resp.Version, err
}

// DeleteNewManifestParam :
func DeleteNewManifestParam(service *ssm.SSM) (bool, error) {
	keyname := "/app/deploy/new"

	req := service.DeleteParameterRequest(&ssm.DeleteParameterInput{
		Name: &keyname,
	})

	_, err := req.Send()
	if err != nil {
		return false, err
	}

	return true, err
}

// UpdateParamStore :
func UpdateParamStore(svc *ssm.SSM) bool {
	newManifestPath, err := GetNewManifestPath(svc)
	if err != nil {
		panic("failed to describe new parameter store, " + err.Error())
	}

	_, err = UpdateCurrentManifestPath(svc, newManifestPath)
	if err != nil {
		panic("failed to update current parameter store, " + err.Error())
	}

	_, err = DeleteNewManifestParam(svc)
	if err != nil {
		panic("failed to delete new parameter store, " + err.Error())
	}

	fmt.Println("Success")
	return true
}
