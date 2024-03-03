package services

import (
	"github.com/google/uuid"
	"github.com/wuhoops/silenda/backend/internal/domains"
	"github.com/wuhoops/silenda/backend/internal/repositories"
)

type workspaceService struct {
	workspaceRepository repositories.WorkspaceRepository
}

func NewWorkspaceService(workspaceRepository repositories.WorkspaceRepository) WorkspaceService {
	return &workspaceService{
		workspaceRepository: workspaceRepository,
	}
}

func (s *workspaceService) CreateWorkspace(workspace domains.Workspace) error {
	if err := s.workspaceRepository.CreateWorkspace(workspace); err != nil {
		return ErrorWorkspaceCreate
	}
	return nil
}

func (s *workspaceService) FindAllWorkspaces(id uuid.UUID) ([]domains.Workspace, error) {
	workspaces, err := s.workspaceRepository.FindAllWorkspaces(id)
	if err != nil {
		return nil, ErrorWorkspaceNotFound
	}
	return workspaces, nil
}
