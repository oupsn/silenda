/*
Copyright © 2024 Pasinun.w
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/wuhoops/silenda/cli/models"
	"github.com/wuhoops/silenda/cli/utils"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Used to initialize the workspace id",
	Long:  `This command is used to initialize the workspace id.`,
	Run: func(cmd *cobra.Command, args []string) {
		wk, err := cmd.Flags().GetString("workspace-id")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		_, err = os.ReadFile(utils.CONFIG_FILE_NAME)
		if err == nil {
			fmt.Println("Workspace id already exists")
			prompt := promptui.Select{
				Label: "Do you want to overwrite it?",
				Items: []string{"Yes", "No"},
			}
			_, result, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed: %v\n", err)
				return
			}
			if result == "No" {
				fmt.Println("Overwrite cancelled")
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
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("workspace-id", "w", "", "Set the workspace id for the project")
}

func writeConfigFile(wk string) error {
	data := models.InitBody{
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
