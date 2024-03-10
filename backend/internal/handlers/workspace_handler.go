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
		Name:   body.Name,
		UserID: OwnerUuid,
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

// UpdateWorkspace godoc
// @tags workspace
// @id updateWorkspace
// @summary Update workspace
// @accept json
// @produce json
// @Param payload body UpdateWorkspaceBody true "UpdateWorkspaceBody"
// @Success 200 {object} Response[string]
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /workspace.updateWorkspace [post]
func (h *WorkspaceHandler) UpdateWorkspace(c *fiber.Ctx) error {
	var body UpdateWorkspaceBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	workspaceUuid, err := uuid.Parse(body.ID)
	if err != nil {
		return err
	}
	payload := domains.Workspace{
		ID:   workspaceUuid,
		Name: body.Name,
	}
	if err := h.workspaceService.UpdateWorkspace(payload); err != nil {
		return err
	}

	return Ok(c, "Workspace updated")
}

// DeleteWorkspace godoc
// @tags workspace
// @id deleteWorkspace
// @summary Delete workspace
// @accept json
// @produce json
// @Param payload body DeleteWorkspaceBody true "DeleteWorkspaceBody"
// @Success 200 {object} Response[string]
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /workspace.deleteWorkspace/ [delete]
func (h *WorkspaceHandler) DeleteWorkspace(c *fiber.Ctx) error {
	var body DeleteWorkspaceBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	workspaceUuid, err := uuid.Parse(body.ID)
	if err != nil {
		return err
	}
	if err := h.workspaceService.DeleteWorkspace(workspaceUuid); err != nil {
		return err
	}

	return Ok(c, "Workspace deleted")
}

// AddUserToWorkspace godoc
// @tags workspace
// @id addUserToWorkspace
// @summary Add user to workspace
// @accept json
// @produce json
// @Param payload body AddUserToWorkspaceBody true "AddUserToWorkspaceBody"
// @Success 200 {object} Response[string]
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /workspace.addUserToWorkspace [post]
func (h *WorkspaceHandler) AddUserToWorkspace(c *fiber.Ctx) error {
	var body AddUserToWorkspaceBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	workspaceUuid, err := uuid.Parse(body.WorkspaceID)
	if err != nil {
		return err
	}
	userUuid, err := uuid.Parse(body.UserID)
	if err != nil {
		return err
	}
	payload := domains.Member{
		WorkspaceID: workspaceUuid,
		UserID:      userUuid,
	}
	if err := h.workspaceService.AddUserToWorkspace(payload); err != nil {
		return err
	}

	return Ok(c, "User added to workspace")
}

// RemoveUserFromWorkspace godoc
// @tags workspace
// @id removeUserFromWorkspace
// @summary Remove user from workspace
// @accept json
// @produce json
// @Param payload body RemoveUserFromWorkspaceBody true "RemoveUserFromWorkspaceBody"
// @Success 200 {object} Response[string]
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} ErrResponse
// @Router /workspace.removeUserFromWorkspace [delete]
func (h *WorkspaceHandler) RemoveUserFromWorkspace(c *fiber.Ctx) error {
	var body RemoveUserFromWorkspaceBody
	if err := c.BodyParser(&body); err != nil {
		return err
	}
	workspaceUuid, err := uuid.Parse(body.WorkspaceID)
	if err != nil {
		return err
	}
	userUuid, err := uuid.Parse(body.UserID)
	if err != nil {
		return err
	}
	payload := domains.Member{
		WorkspaceID: workspaceUuid,
		UserID:      userUuid,
	}
	if err := h.workspaceService.RemoveUserFromWorkspace(payload); err != nil {
		return err
	}

	return Ok(c, "User removed from workspace")
}
