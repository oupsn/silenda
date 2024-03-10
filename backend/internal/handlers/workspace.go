package handlers

type CreateWorkspaceBody struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
} // @name CreateWorkspaceBody

type FindAllWorkspacesBody struct {
	Owner string `json:"owner"`
} // @name FindAllWorkspacesBody

type FindAllWorkspacesResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
} // @name FindAllWorkspacesResponse

type UpdateWorkspaceBody struct {
	ID   string `json:"id"`
	Name string `json:"name"`
} // @name UpdateWorkspaceBody

type DeleteWorkspaceBody struct {
	ID string `json:"workspace_id"`
} // @name DeleteWorkspaceBody

type AddUserToWorkspaceBody struct {
	WorkspaceID string `json:"workspace_id"`
	UserID      string `json:"user_id"`
	Role        string `json:"role"`
} // @name AddUserToWorkspaceBody

type RemoveUserFromWorkspaceBody struct {
	WorkspaceID string `json:"workspace_id"`
	UserID      string `json:"user_id"`
} // @name RemoveUserFromWorkspaceBody
