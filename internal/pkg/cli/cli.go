package cli

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "api",
		Short: "Simple Go API server",
		Long:  `High performance, extensible, minimalist REST API`,
	}
)

// execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
