package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wuhoops/silenda/backend/internal/domains"
)

var (
	ErrorWorkspaceCreate   = fiber.NewError(fiber.StatusInternalServerError, "Error creating workspace")
	ErrorWorkspaceNotFound = fiber.NewError(fiber.StatusNotFound, "Workspace not found")
)

type WorkspaceService interface {
	CreateWorkspace(workspace domains.Workspace) error
	FindAllWorkspaces(id uuid.UUID) ([]domains.Workspace, error)
	UpdateWorkspace(workspace domains.Workspace) error
	DeleteWorkspace(id uuid.UUID) error
	AddUserToWorkspace(member domains.Member) error
	RemoveUserFromWorkspace(member domains.Member) error
}
