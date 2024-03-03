/*
Copyright © 2024 Pasinun.w
*/
package secretCmd

import (
	"github.com/spf13/cobra"
	"github.com/wuhoops/silenda/cli/cmd"
)

// secretCmd represents the secret command
var secretCmd = &cobra.Command{
	Use:   "secret",
	Short: "secret command",
	Long:  `This command is used to manage secret variables in the workspace.`,
}

func init() {
	cmd.RootCmd.AddCommand(secretCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// secretCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// secretCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
