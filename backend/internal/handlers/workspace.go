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
