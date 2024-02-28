package handlers

type FindSecretsByEnvModeBody struct {
	WorkspaceID string `json:"workspace_id"`
	EnvMode     string `json:"env_mode"`
}

type FindSecretsByEnvModeResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CreateSecretBody struct {
	WorkspaceID string `json:"workspace_id"`
	EnvMode     string `json:"env_mode"`
	Key         string `json:"key"`
	Value       string `json:"value"`
}
