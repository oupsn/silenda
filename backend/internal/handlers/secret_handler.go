package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wuhoops/silenda/backend/internal/domains"
	"github.com/wuhoops/silenda/backend/internal/services"
)

type SecretHandler struct {
	secretService services.SecretService
}

func NewSecretHandler(secretService services.SecretService) SecretHandler {
	return SecretHandler{
		secretService: secretService,
	}
}

func (s *SecretHandler) FindSecretsByEnvMode(c *fiber.Ctx) error {
	var body FindSecretsByEnvModeBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	secrets, err := s.secretService.FindSecretsByEnvMode(body.WorkspaceID, body.EnvMode)
	if err != nil {
		return err
	}
	var response []FindSecretsByEnvModeResponse
	for _, secret := range secrets {
		response = append(response, FindSecretsByEnvModeResponse{
			Key:   *secret.Key,
			Value: *secret.Value,
		})
	}

	return Ok(c, response)
}

func (s *SecretHandler) CreateSecret(c *fiber.Ctx) error {
	var body CreateSecretBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	UuidWorkspaceId, err := uuid.Parse(body.WorkspaceID)
	if err != nil {
		return err
	}
	secret := domains.Secret{
		WorkspaceID: &UuidWorkspaceId,
		EnvMode:     (*domains.EnvMode)(&body.EnvMode),
		Key:         &body.Key,
		Value:       &body.Value,
	}
	if err := s.secretService.CreateSecret(secret); err != nil {
		return err
	}
	return Ok(c, "Create secret successfully!")
}
