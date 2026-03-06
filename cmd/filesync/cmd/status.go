package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "show cluster and sync health",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("status called")
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().Bool("watch", false, "live-reresh output")
	statusCmd.Flags().String("peer", "", "show status for a specific peer only")
	statusCmd.Flags().Bool("json", false, "machine-readable output")
}
