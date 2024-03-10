/*
Copyright © 2024 Pasinun.w
*/
package secretCmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
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
		if loaders.Wc.WorkSpaceId == "" {
			fmt.Print("Error: ", "workspace id has not been set, HINT: try `silenda init <workspace-id>` to set workspace id.")
			return
		}
		prompt := promptui.Select{
			Label: "Please select the environment",
			Items: []string{"dev", "stage", "prod"},
		}
		_, env, err := prompt.Run()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		prompt2 := promptui.Prompt{
			Label: "Enter the key",
		}
		key, err := prompt2.Run()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		prompt3 := promptui.Prompt{
			Label: "Enter the value",
		}
		value, err := prompt3.Run()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		body := models.CreateSecretBody{
			WorkspaceID: loaders.Wc.WorkSpaceId,
			EnvMode:     env,
			Key:         key,
			Value:       value,
		}
		err = api.CreateSecret(body)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Secret variable has been added successfully.")
	},
}

func init() {
	secretCmd.AddCommand(addCmd)
}
