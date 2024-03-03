/*
Copyright © 2024 Pasinun.w
*/
package main

import (
	"github.com/wuhoops/silenda/cli/cmd"
	_ "github.com/wuhoops/silenda/cli/cmd/secretCmd"
	"github.com/wuhoops/silenda/cli/loaders"
)

func main() {
	loaders.WorkspaceConfigInit()
	cmd.Execute()
}
