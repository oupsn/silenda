/*
Copyright © 2024 Pasinun.w
*/
package main

import (
	"github.com/wuhoops/silenda/cmd"
	"github.com/wuhoops/silenda/loaders"
)

func main() {
	loaders.Configs()
	cmd.Execute()
}
