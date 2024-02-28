/*
Copyright © 2024 Pasinun.w
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wuhoops/silenda/cli/api"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all secret variables from the workspace",
	Long:  "List all secret variables from the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		api.GetAllEncryptedSecretVariables()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
