package services

import (
	"github.com/google/uuid"
	"github.com/wuhoops/silenda/backend/internal/domains"
	"github.com/wuhoops/silenda/backend/internal/repositories"
)

type secretService struct {
	secretRepository repositories.SecretRepository
}

func NewSecretService(secretRepository repositories.SecretRepository) SecretService {
	return &secretService{
		secretRepository: secretRepository,
	}
}

func (s secretService) FindSecretsByEnvMode(workspaceId uuid.UUID, envMode string) ([]domains.Secret, error) {
	secrets, err := s.secretRepository.FindSecretsByEnvMode(workspaceId, envMode)
	if err != nil {
		return nil, ErrorSecretNotFound
	}
	return secrets, nil
}

func (s secretService) CreateSecret(secret domains.Secret) error {
	if err := s.secretRepository.CreateSecret(secret); err != nil {
		return ErrorSecretExists
	}
	return nil
}

func (s secretService) DeleteSecretById(id uuid.UUID) error {
	if err := s.secretRepository.DeleteSecretById(id); err !=
		nil {
		return ErrorSecretNotFound
	}
	return nil
}
