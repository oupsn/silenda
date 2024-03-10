/*
Copyright © 2024 Pasinun.w
*/
package secretCmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
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
		prompt := promptui.Select{
			Label: "Please select the environment",
			Items: []string{"dev", "stage", "prod", "all"},
		}
		_, env, err := prompt.Run()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		if env == "all" {
			body := models.FindAllSecretsByWorkspaceIDBody{
				WorkspaceID: loaders.Wc.WorkSpaceId,
			}
			resp, err := api.GetAllSecretsByWorkspaceID(body)
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()
			tbl := table.New("Key", "Dev's value", "Stage's value", "Prod's value")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			type res struct {
				Key   string
				Dev   string
				Stage string
				Prod  string
			}
			var t []res
			for i, s := range resp.Dev {
				t = append(t, res{s.Key, s.Value, resp.Dev[i].Value, resp.Prod[i].Value})
			}
			for _, s := range t {
				tbl.AddRow(s.Key, s.Dev, s.Stage, s.Prod)
			}
			tbl.Print()
		} else {
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
		}
	},
}

func init() {
	secretCmd.AddCommand(listCmd)
	listCmd.Flags().StringP("env", "e", "", "environment mode")
}
