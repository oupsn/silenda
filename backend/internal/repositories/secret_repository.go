package repositories

import (
	"fmt"
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

func (s secretRepository) FindSecretsByEnvMode(workspaceId uuid.UUID, envMode string) ([]domains.Secret, error) {
	var secrets []domains.Secret
	if err := s.DB.Where("workspace_id = ?", workspaceId).Where("env_mode = ?", envMode).Find(&secrets).Error; err != nil {
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

func (s secretRepository) DeleteSecretById(id uuid.UUID) error {
	if err := s.DB.Where("id = ?", id).Delete(&domains.Secret{}).Error; err != nil {
		return err
	}
	return nil
}

func (s secretRepository) UpdateSecret(secret domains.Secret) error {
	if err := s.DB.Where("id = ?", secret.ID).Updates(secret).Error; err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}
