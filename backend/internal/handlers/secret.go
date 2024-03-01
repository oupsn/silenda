package handlers

type FindSecretsByEnvModeBody struct {
	WorkspaceID string `json:"workspace_id"`
	EnvMode     string `json:"env_mode"`
} // @name FindSecretsByEnvModeBody

type FindSecretsByEnvModeResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
} // @name FindSecretsByEnvModeResponse

type CreateSecretBody struct {
	WorkspaceID string `json:"workspace_id"`
	EnvMode     string `json:"env_mode"`
	Key         string `json:"key"`
	Value       string `json:"value"`
} // @name CreateSecretBody
