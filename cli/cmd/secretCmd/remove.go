/*
Copyright © 2024 Pasinun.w
*/
package secretCmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/wuhoops/silenda/cli/api"
	"github.com/wuhoops/silenda/cli/loaders"
	"github.com/wuhoops/silenda/cli/models"
)

// addCmd represents the add command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove secret variable",
	Long:  `remove secret variable from the workspace.`,
	Run: func(cmd *cobra.Command, args []string) {
		prompt := promptui.Select{
			Label: "Please select the environment",
			Items: []string{"dev", "stg", "prod"},
		}
		_, env, err := prompt.Run()
		body := models.FindSecretsByEnvModeBody{
			EnvMode:     env,
			WorkspaceID: loaders.Wc.WorkSpaceId,
		}
		resp, err := api.GetAllEncryptedSecretVariables(body)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		var secrets []string
		for _, s := range resp {
			secrets = append(secrets, s.Key+": "+s.Value)
		}
		prompt = promptui.Select{
			Label: "Please select the secret variable to remove",
			Items: secrets,
		}
		index, _, err := prompt.Run()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		body2 := models.DeleteSecretByIDBody{
			SecretID: resp[index].ID,
		}
		err = api.DeleteSecret(body2)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Secret", resp[index].Key, "has been removed successfully.")
	},
}

func init() {
	secretCmd.AddCommand(removeCmd)
}
