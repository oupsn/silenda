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

func (s secretService) FindAllSecretsByWorkspaceId(id uuid.UUID) ([]domains.Secret, []domains.Secret, []domains.Secret, error) {
	devSecrets, err := s.secretRepository.FindSecretsByEnvMode(id, "dev")
	if err != nil {
		return nil, nil, nil, ErrorSecretNotFound
	}
	stageSecrets, err := s.secretRepository.FindSecretsByEnvMode(id, "stage")
	if err != nil {
		return nil, nil, nil, ErrorSecretNotFound
	}
	prodSecrets, err := s.secretRepository.FindSecretsByEnvMode(id, "prod")
	if err != nil {
		return nil, nil, nil, ErrorSecretNotFound
	}
	return devSecrets, stageSecrets, prodSecrets, nil
}

func (s secretService) UpdateSecret(secret domains.Secret) error {
	if err := s.secretRepository.UpdateSecret(secret); err != nil {
		return ErrorSecretNotFound
	}
	return nil
}
