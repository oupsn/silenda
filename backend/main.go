package main

import (
	"context"
	"github.com/wuhoops/silenda/backend/loaders"
)

type App struct {
	ctx context.Context
}

func main() {
	loaders.SetupDatabases()
	loaders.SetupRoutes()
}
