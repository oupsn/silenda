package services

import (
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

func (s secretService) FindSecretsByEnvMode(workspaceId string, envMode string) ([]domains.Secret, error) {
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
