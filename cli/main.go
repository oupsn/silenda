/*
Copyright © 2024 Pasinun.w
*/
package main

import (
	"github.com/wuhoops/silenda/cli/cmd"
	"github.com/wuhoops/silenda/cli/loaders"
)

func main() {
	loaders.Configs()
	cmd.Execute()
}
