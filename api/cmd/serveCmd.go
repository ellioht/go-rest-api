package cmd

import (
	"github.com/ellioht/go-rest-api/internal/app"
	"github.com/spf13/cobra"
)

func ServeCmd(app *app.App) *cobra.Command {
	serve := &cobra.Command{
		Use:   "serve",
		Short: "Start the API server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return serve
}

func Serve(app *app.App) *cobra.Command {
	serve := ServeCmd(app)
	return serve
}
