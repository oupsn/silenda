/*
Copyright © 2024 Pasinun.w
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/wuhoops/silenda/cli/utils"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type InitBody struct {
	WorkSpaceId string `json:"workspace_id"`
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Used to initialize the workspace id",
	Long:  `This command is used to initialize the workspace id.`,
	Run: func(cmd *cobra.Command, args []string) {
		wk := args[0]
		if wk == "" {
			fmt.Println("Error: workspace id is required, HINT: try `silenda init <workspace-id>`")
			return
		}

		_, err := os.ReadFile(utils.CONFIG_FILE_NAME)
		if err == nil {
			prompt := promptui.Select{
				Label: "Workspace id already set. Do you want to overwrite it?",
				Items: []string{"Yes", "No"},
			}
			_, result, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed: %v\n", err)
				return
			}
			if result == "No" {
				fmt.Println("Initialization canceled")
				return
			}
			if result == "Yes" {
				err := writeConfigFile(wk)
				if err != nil {
					fmt.Println("Error: ", err)
				}
			}
		} else {
			err := writeConfigFile(wk)
			if err != nil {
				fmt.Println("Error: ", err)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}

func writeConfigFile(wk string) error {
	data := InitBody{
		WorkSpaceId: wk,
	}
	d, err := json.Marshal(&data)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	err = os.WriteFile(utils.CONFIG_FILE_NAME, d, 0644)
	if err != nil {
		return err
	}
	fmt.Println("Workspace key initialized")
	return nil
}
