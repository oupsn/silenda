/*
Copyright © 2024 Pasinun.w
*/
package secretCmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
	"github.com/wuhoops/silenda/cli/api"
	"github.com/wuhoops/silenda/cli/loaders"
	"github.com/wuhoops/silenda/cli/models"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all secret variables from the workspace",
	Long:  "List all secret variables from the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Print("Error: ", "environment mode is required, HINT: try `silenda list <env-mode>`")
			return
		}
		env := args[0]
		if loaders.Wc.WorkSpaceId == "" {
			fmt.Print("Error: ", "workspace id has not been set, HINT: try `silenda init <workspace-id>`")
			return
		}
		body := models.FindSecretsByEnvModeBody{
			EnvMode:     env,
			WorkspaceID: loaders.Wc.WorkSpaceId,
		}

		resp, err := api.GetAllEncryptedSecretVariables(body)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()
		tbl := table.New("Key", "Value")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		for _, s := range resp {
			tbl.AddRow(s.Key, s.Value)
		}

		tbl.Print()
	},
}

func init() {
	secretCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("env", "e", "", "environment mode")
}
