package repositories

import (
	"github.com/google/uuid"
	"github.com/wuhoops/silenda/backend/internal/domains"
)

type SecretRepository interface {
	FindSecretsByEnvMode(workspaceId uuid.UUID, envMode string) ([]domains.Secret, error)
	CreateSecret(secret domains.Secret) error
	DeleteSecretById(id uuid.UUID) error
}
