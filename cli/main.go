/*
Copyright © 2024 Pasinun.w
*/
package main

import (
	"github.com/wuhoops/silenda/cli/cmd"
	_ "github.com/wuhoops/silenda/cli/cmd/secretCmd"
	_ "github.com/wuhoops/silenda/cli/cmd/workspaceCmd"
	"github.com/wuhoops/silenda/cli/loaders"
)

func main() {
	loaders.WorkspaceConfigInit()
	cmd.Execute()
}
