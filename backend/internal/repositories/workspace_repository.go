package repositories

import (
	"github.com/google/uuid"
	"github.com/wuhoops/silenda/backend/internal/domains"
	"gorm.io/gorm"
)

type workspaceRepository struct {
	DB *gorm.DB
}

func NewWorkspaceRepository(db *gorm.DB) WorkspaceRepository {
	return &workspaceRepository{
		DB: db,
	}
}

func (r *workspaceRepository) CreateWorkspace(workspace domains.Workspace) error {
	if err := r.DB.Create(&workspace).Error; err != nil {
		return err
	}
	return nil
}

func (r *workspaceRepository) FindWorkspaceById(id uuid.UUID) (*domains.Workspace, error) {
	var workspace *domains.Workspace
	if err := r.DB.Where("id = ?", id).First(&workspace).Error; err != nil {
		return nil, err
	}
	return workspace, nil
}

func (r *workspaceRepository) FindAllWorkspaces(ownerId uuid.UUID) ([]domains.Workspace, error) {
	var workspaces []domains.Workspace
	if err := r.DB.Where("owner = ?", ownerId).Find(&workspaces).Error; err != nil {
		return nil, err
	}
	return workspaces, nil
}
