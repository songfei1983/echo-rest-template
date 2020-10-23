package cli

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "restapi",
		Short: "Go REST API template",
		Long:  `High performance, extensible, minimalist REST API`,
	}
)

// execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
