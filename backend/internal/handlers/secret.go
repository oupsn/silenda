package handlers

type FindSecretsByEnvModeBody struct {
	WorkspaceID string `json:"workspace_id"`
	EnvMode     string `json:"env_mode"`
} // @name FindSecretsByEnvModeBody

type FindSecretsByEnvModeResponse struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
} // @name FindSecretsByEnvModeResponse

type CreateSecretBody struct {
	WorkspaceID string `json:"workspace_id"`
	EnvMode     string `json:"env_mode"`
	Key         string `json:"key"`
	Value       string `json:"value"`
} // @name CreateSecretBody

type DeleteSecretByIdBody struct {
	SecretID string `json:"secret_id"`
} // @name DeleteSecretByIdBody

type FindAllSecretsByWorkspaceIdBody struct {
	WorkspaceID string `json:"workspace_id"`
} // @name FindAllSecretsByWorkspaceIdBody

type FindAllSecretsByWorkspaceIdResponse struct {
	Dev   []FindSecretsByEnvModeResponse `json:"dev"`
	Stage []FindSecretsByEnvModeResponse `json:"stage"`
	Prod  []FindSecretsByEnvModeResponse `json:"prod"`
} // @name FindAllSecretsByWorkspaceIdResponse

type UpdateSecretByIdBody struct {
	SecretID string `json:"secret_id"`
	Value    string `json:"value"`
} // @name UpdateSecretByIdBody
