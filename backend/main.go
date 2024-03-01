package main

import (
	"context"
	"github.com/wuhoops/silenda/backend/loaders"
)

type App struct {
	ctx context.Context
}

// @title Silenda API
// @version 1.0
// @contact.name Pasinun.w
// @contact.email pasinun.w@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @schemes http

func main() {
	loaders.SetupDatabases()
	loaders.SetupRoutes()
}
