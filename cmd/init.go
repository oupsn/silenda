/*
Copyright © 2024 Pasinun.w
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/wuhoops/silenda/models"
	"github.com/wuhoops/silenda/utils"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Used to initialize the workspace key",
	Long:  `This command is used to initialize the workspace key.`,
	Run: func(cmd *cobra.Command, args []string) {
		wk, err := cmd.Flags().GetString("workspace-key")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		_, err = os.ReadFile(utils.CONFIG_FILE_NAME)
		if err == nil {
			fmt.Println("Workspace key already exists")
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
	initCmd.Flags().StringP("workspace-key", "w", "", "Set the workspace key for the project")
}

func writeConfigFile(wk string) error {
	data := models.InitBody{
		WorkSpaceKey: wk,
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
