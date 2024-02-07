package cmd

import (
	"fmt"
	"github.com/ellioht/go-rest-api/config"
	"github.com/ellioht/go-rest-api/internal/app"
	"github.com/spf13/cobra"
	"os"
)

func RootCmd(app *app.App, cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:              "root",
		Short:            "rootCmd represents the base command when called without any subcommands",
		TraverseChildren: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// cfg
			return nil
		},
	}
	return cmd
}

func Root(app *app.App, cfg *config.Config) *cobra.Command {
	root := RootCmd(app, cfg)
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return root
}
