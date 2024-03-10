package repositories

import (
	"github.com/google/uuid"
	"github.com/wuhoops/silenda/backend/internal/domains"
)

type WorkspaceRepository interface {
	CreateWorkspace(workspace domains.Workspace) error
	FindWorkspaceById(id uuid.UUID) (*domains.Workspace, error)
	FindAllWorkspaces(ownerId uuid.UUID) ([]domains.Workspace, error)
	UpdateWorkspace(workspace domains.Workspace) error
	DeleteWorkspace(id uuid.UUID) error
	AddUserToWorkspace(member domains.Member) error
	RemoveUserFromWorkspace(member domains.Member) error
}
