package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var _ = InitVersionCmd()

func InitVersionCmd() struct{} {
	rootCmd.AddCommand(versionCmd)
	return struct{}{}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "0.0.1",
	Long:  `version 0.0.1`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go API Server version 0.0.1")
	},
}
