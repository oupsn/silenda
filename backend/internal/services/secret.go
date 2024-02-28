package services

import (
	"github.com/wuhoops/silenda/backend/internal/domains"
)

var ()

type SecretService interface {
	FindSecretsByEnvMode(workspaceId string, envMode string) ([]domains.Secret, error)
	CreateSecret(secret domains.Secret) error
}
