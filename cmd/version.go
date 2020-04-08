package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "0.0.1",
	Long:  `version 0.0.1`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go API Server version 0.0.1")
	},
}
