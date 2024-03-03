/*
Copyright © 2024 Pasinun.w
*/
package secretCmd

import (
	"fmt"
	"github.com/wuhoops/silenda/cli/api"
	"github.com/wuhoops/silenda/cli/loaders"
	"github.com/wuhoops/silenda/cli/models"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add secret variable",
	Long:  `Add secret variable to the workspace.`,
	Run: func(cmd *cobra.Command, args []string) {
		env, _ := cmd.Flags().GetString("env")
		key, _ := cmd.Flags().GetString("key")
		value, _ := cmd.Flags().GetString("value")
		if env == "" || key == "" || value == "" {
			fmt.Print("Error: ", "environment mode is required, HINT: try `silenda add --env <env-mode> --key <key> --value <value>`")
			return
		}
		if loaders.Wc.WorkSpaceId == "" {
			fmt.Print("Error: ", "workspace id has not been set, HINT: try `silenda init <workspace-id>` to set workspace id.")
			return
		}
		body := models.CreateSecretBody{
			WorkspaceID: loaders.Wc.WorkSpaceId,
			EnvMode:     env,
			Key:         key,
			Value:       value,
		}

		err := api.CreateSecret(body)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Secret variable has been added successfully.")
	},
}

func init() {
	secretCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("env", "e", "", "environment mode")
	addCmd.Flags().StringP("key", "k", "", "secret key")
	addCmd.Flags().StringP("value", "v", "", "secret value")
}
