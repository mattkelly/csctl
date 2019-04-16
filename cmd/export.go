package cmd

import (
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a resource",
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
