package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wuhoops/silenda/backend/internal/domains"
)

var (
	ErrorSecretNotFound = fiber.NewError(fiber.StatusNotFound, "Secret not found")
	ErrorSecretExists   = fiber.NewError(fiber.StatusBadRequest, "Secret already exists")
)

type SecretService interface {
	FindSecretsByEnvMode(workspaceId string, envMode string) ([]domains.Secret, error)
	CreateSecret(secret domains.Secret) error
}
