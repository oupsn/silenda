package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wuhoops/silenda/backend/internal/domains"
	"github.com/wuhoops/silenda/backend/internal/services"
)

type WorkspaceHandler struct {
	workspaceService services.WorkspaceService
}

func NewWorkspaceHandler(workspaceService services.WorkspaceService) WorkspaceHandler {
	return WorkspaceHandler{
		workspaceService: workspaceService,
	}
}

// CreateWorkspace godoc
// @tags workspace
// @id createWorkspace
// @summary Create a new workspace
// @accept json
// @produce json
// @Param payload body CreateWorkspaceBody true "CreateWorkspaceBody"
// @Success 200 {object} Response[string]
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /workspace.createWorkspace [post]
func (h *WorkspaceHandler) CreateWorkspace(c *fiber.Ctx) error {
	var body CreateWorkspaceBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	OwnerUuid, err := uuid.Parse(body.Owner)
	if err != nil {
		return err
	}
	payload := domains.Workspace{
		Name:  body.Name,
		Owner: OwnerUuid,
	}
	if err := h.workspaceService.CreateWorkspace(payload); err != nil {
		return err
	}

	return Ok(c, "Workspace created")
}

// FindAllWorkspaces godoc
// @tags workspace
// @id findAllWorkspaces
// @summary Find all workspaces
// @accept json
// @produce json
// @Param payload body FindAllWorkspacesBody true "FindAllWorkspacesBody"
// @Success 200 {object} Response[[]FindAllWorkspacesResponse]
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /workspace.findAllWorkspaces [post]
func (h *WorkspaceHandler) FindAllWorkspaces(c *fiber.Ctx) error {
	var body FindAllWorkspacesBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	ownerUuid, err := uuid.Parse(body.Owner)
	if err != nil {
		return err
	}
	workspaces, err := h.workspaceService.FindAllWorkspaces(ownerUuid)
	if err != nil {
		return err
	}
	var response []FindAllWorkspacesResponse
	for _, workspace := range workspaces {
		response = append(response, FindAllWorkspacesResponse{
			ID:   workspace.ID.String(),
			Name: workspace.Name,
		})
	}

	return Ok(c, workspaces)
}
