package api

import (
	"errors"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	client2 "github.com/wuhoops/silenda/cli/client"
	"github.com/wuhoops/silenda/cli/client/secret"
	"github.com/wuhoops/silenda/cli/models"
)

var client *client2.API

func init() {
	transport := httptransport.New("localhost:8080", "/api/v1", []string{"http"})
	client = client2.New(transport, strfmt.Default)
}

func GetAllEncryptedSecretVariables(body secret.FindSecretsByEnvModeParams) ([]models.FindSecretsByEnvModeResponse, error) {
	if body.Payload.EnvMode == "" {
		return nil, errors.New("environment mode is required")
	}
	resp, err := client.Secret.FindSecretsByEnvMode(&secret.FindSecretsByEnvModeParams{
		Payload: &models.FindSecretsByEnvModeBody{
			EnvMode:     body.Payload.EnvMode,
			WorkspaceID: body.Payload.WorkspaceID,
		},
	})
	if err != nil {
		return nil, err
	}
	var response []models.FindSecretsByEnvModeResponse
	for _, s := range resp.Payload.Data {
		response = append(response, models.FindSecretsByEnvModeResponse{
			Key:   s.Key,
			Value: s.Value,
		})
	}
	return response, nil
}

func CreateSecret(body secret.CreateSecretParams) error {
	if body.Payload.EnvMode == "" {
		return errors.New("environment mode is required")
	}
	if body.Payload.Key == "" {
		return errors.New("key is required")
	}
	if body.Payload.Value == "" {
		return errors.New("value is required")
	}
	_, err := client.Secret.CreateSecret(&secret.CreateSecretParams{
		Payload: &models.CreateSecretBody{
			EnvMode:     body.Payload.EnvMode,
			WorkspaceID: body.Payload.WorkspaceID,
			Key:         body.Payload.Key,
			Value:       body.Payload.Value,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
