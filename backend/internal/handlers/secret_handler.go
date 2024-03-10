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

// FindSecretsByEnvMode godoc
// @tags secret
// @id findSecretsByEnvMode
// @summary Find secrets by env mode
// @accept json
// @produce json
// @Param payload body FindSecretsByEnvModeBody true "FindSecretsByEnvModeBody"
// @Success 200 {object} Response[[]FindSecretsByEnvModeResponse]
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /secret.findSecretsByEnvMode [post]
func (s *SecretHandler) FindSecretsByEnvMode(c *fiber.Ctx) error {
	var body FindSecretsByEnvModeBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	workspaceIdUuid, err := uuid.Parse(body.WorkspaceID)
	if err != nil {
		return err
	}
	secrets, err := s.secretService.FindSecretsByEnvMode(workspaceIdUuid, body.EnvMode)
	if err != nil {
		return err
	}
	var response []FindSecretsByEnvModeResponse
	for _, secret := range secrets {
		response = append(response, FindSecretsByEnvModeResponse{
			ID:    secret.ID.String(),
			Key:   secret.Key,
			Value: secret.Value,
		})
	}

	return Ok(c, response)
}

// CreateSecret godoc
// @tags secret
// @id createSecret
// @summary Create secret
// @accept json
// @produce json
// @Param payload body CreateSecretBody true "CreateSecretBody"
// @Success 200 {object} Response[string]
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /secret.createSecret [post]
func (s *SecretHandler) CreateSecret(c *fiber.Ctx) error {
	var body CreateSecretBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	UuidWorkspaceIdUuid, err := uuid.Parse(body.WorkspaceID)
	if err != nil {
		return err
	}
	secret := domains.Secret{
		WorkspaceID: UuidWorkspaceIdUuid,
		EnvMode:     (domains.EnvMode)(body.EnvMode),
		Key:         body.Key,
		Value:       body.Value,
	}
	if err := s.secretService.CreateSecret(secret); err != nil {
		return err
	}
	return Created(c, "Create secret successfully!")
}

// DeleteSecretById godoc
// @tags secret
// @id deleteSecretById
// @summary Delete secret by id
// @accept json
// @produce json
// @Param payload body DeleteSecretByIdBody true "DeleteSecretByIdBody"
// @Success 200 {object} Response[string]
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /secret.deleteSecretByID [post]
func (s *SecretHandler) DeleteSecretById(c *fiber.Ctx) error {
	var body DeleteSecretByIdBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	secretIdUuid, err := uuid.Parse(body.SecretID)
	if err != nil {
		return err
	}
	err = s.secretService.DeleteSecretById(secretIdUuid)
	if err != nil {
		return err
	}
	return Ok(c, "delete secret successfully")
}

// FindAllSecretsByWorkspaceId godoc
// @tags secret
// @id findAllSecretsByWorkspaceId
// @summary Find all secrets by workspace id
// @accept json
// @produce json
// @Param payload body FindAllSecretsByWorkspaceIdBody true "FindAllSecretsByWorkspaceIdBody"
// @Success 200 {object} FindAllSecretsByWorkspaceIdResponse
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /secret.findAllSecretsByWorkspaceId [post]
func (s *SecretHandler) FindAllSecretsByWorkspaceId(c *fiber.Ctx) error {
	var body FindAllSecretsByWorkspaceIdBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	workspaceIdUuid, err := uuid.Parse(body.WorkspaceID)
	if err != nil {
		return err
	}
	devSecrets, stageSecrets, prodSecrets, err := s.secretService.FindAllSecretsByWorkspaceId(workspaceIdUuid)
	if err != nil {
		return err
	}
	resDevSecrets := make([]FindSecretsByEnvModeResponse, 0)
	for _, secret := range devSecrets {
		resDevSecrets = append(resDevSecrets, FindSecretsByEnvModeResponse{
			ID:    secret.ID.String(),
			Key:   secret.Key,
			Value: secret.Value,
		})
	}
	resStageSecrets := make([]FindSecretsByEnvModeResponse, 0)
	for _, secret := range stageSecrets {
		resStageSecrets = append(resStageSecrets, FindSecretsByEnvModeResponse{
			ID:    secret.ID.String(),
			Key:   secret.Key,
			Value: secret.Value,
		})
	}
	resProdSecrets := make([]FindSecretsByEnvModeResponse, 0)
	for _, secret := range prodSecrets {
		resProdSecrets = append(resProdSecrets, FindSecretsByEnvModeResponse{
			ID:    secret.ID.String(),
			Key:   secret.Key,
			Value: secret.Value,
		})
	}
	response := FindAllSecretsByWorkspaceIdResponse{
		Dev:   resDevSecrets,
		Stage: resStageSecrets,
		Prod:  resProdSecrets,
	}
	return Ok(c, response)
}

// UpdateSecret godoc
// @tags secret
// @id updateSecret
// @summary Update secret
// @accept json
// @produce json
// @Param payload body UpdateSecretByIdBody true "UpdateSecretByIdBody"
// @Success 200 {object} Response[string]
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /secret.updateSecret [post]
func (s *SecretHandler) UpdateSecret(c *fiber.Ctx) error {
	var body UpdateSecretByIdBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	secretIdUuid, err := uuid.Parse(body.SecretID)
	if err != nil {
		return err
	}
	secret := domains.Secret{
		ID:    secretIdUuid,
		Value: body.Value,
	}
	err = s.secretService.UpdateSecret(secret)
	if err != nil {
		return err
	}
	return Ok(c, "update secret successfully")
}
