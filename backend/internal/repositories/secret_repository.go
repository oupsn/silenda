package repositories

import (
	"github.com/google/uuid"
	"github.com/wuhoops/silenda/backend/internal/domains"
	"gorm.io/gorm"
)

type secretRepository struct {
	DB *gorm.DB
}

func NewSecretRepository(db *gorm.DB) SecretRepository {
	return &secretRepository{
		DB: db,
	}
}

func (s secretRepository) FindSecretsByEnvMode(workspaceId string, envMode string) ([]domains.Secret, error) {
	var secrets []domains.Secret
	UuidWorkspaceId, err := uuid.Parse(workspaceId)
	if err != nil {
		return nil, err
	}
	if err := s.DB.Where("workspace_id = ?", UuidWorkspaceId).Where("env_mode = ?", envMode).Find(&secrets).Error; err != nil {
		return nil, err
	}
	return secrets, nil
}

func (s secretRepository) CreateSecret(secret domains.Secret) error {
	if err := s.DB.Create(&secret).Error; err != nil {
		return err
	}
	return nil
}
