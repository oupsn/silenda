package repositories

import (
	"github.com/wuhoops/silenda/backend/internal/domains"
)

type SecretRepository interface {
	FindSecretsByEnvMode(workspaceId string, envMode string) ([]domains.Secret, error)
	CreateSecret(secret domains.Secret) error
}
