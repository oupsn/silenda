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

func GetAllEncryptedSecretVariables(body models.FindSecretsByEnvModeBody) ([]models.FindSecretsByEnvModeResponse, error) {
	if body.EnvMode == "" {
		return nil, errors.New("environment mode is required")
	}
	resp, err := client.Secret.FindSecretsByEnvMode(&secret.FindSecretsByEnvModeParams{
		Payload: &body,
	})
	if err != nil {
		return nil, err
	}
	var response []models.FindSecretsByEnvModeResponse
	for _, s := range resp.Payload.Data {
		response = append(response, models.FindSecretsByEnvModeResponse{
			ID:    s.ID,
			Key:   s.Key,
			Value: s.Value,
		})
	}
	return response, nil
}

func CreateSecret(body models.CreateSecretBody) error {
	if body.EnvMode == "" {
		return errors.New("environment mode is required")
	}
	if body.Key == "" {
		return errors.New("key is required")
	}
	if body.Value == "" {
		return errors.New("value is required")
	}
	_, err := client.Secret.CreateSecret(&secret.CreateSecretParams{
		Payload: &body,
	})
	if err != nil {
		return err
	}
	return nil
}

func DeleteSecret(body models.DeleteSecretByIDBody) error {
	if body.SecretID == "" {
		return errors.New("id is required")
	}
	_, err := client.Secret.DeleteSecretByID(&secret.DeleteSecretByIDParams{
		Payload: &body,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetAllSecretsByWorkspaceID(body models.FindAllSecretsByWorkspaceIDBody) (*models.FindAllSecretsByWorkspaceIDResponse, error) {
	if body.WorkspaceID == "" {
		return nil, errors.New("workspace id is required")
	}
	resp, err := client.Secret.FindAllSecretsByWorkspaceID(&secret.FindAllSecretsByWorkspaceIDParams{
		Payload: &body,
	})
	if err != nil {
		return nil, err
	}
	var response *models.FindAllSecretsByWorkspaceIDResponse
	for _, s := range resp.Payload.Dev {
		response.Dev = append(response.Dev, &models.FindSecretsByEnvModeResponse{
			ID:    s.ID,
			Key:   s.Key,
			Value: s.Value,
		})
	}
	for _, s := range resp.Payload.Stage {
		response.Stage = append(response.Stage, &models.FindSecretsByEnvModeResponse{
			ID:    s.ID,
			Key:   s.Key,
			Value: s.Value,
		})
	}
	for _, s := range resp.Payload.Prod {
		response.Prod = append(response.Prod, &models.FindSecretsByEnvModeResponse{
			ID:    s.ID,
			Key:   s.Key,
			Value: s.Value,
		})
	}
	return response, nil
}
