package main

import (
	"context"
	"github.com/ellioht/go-rest-api/cmd"
	"github.com/ellioht/go-rest-api/config"
	"github.com/ellioht/go-rest-api/internal/app"
)

func main() {
	app.Start(appStart)
}

func appStart(ctx context.Context, app *app.App) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	root := cmd.Root(app, cfg)
	root.AddCommand(cmd.Serve(app))

	return nil
}
