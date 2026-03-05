package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var unwatchCmd = &cobra.Command{
	Use:   "unwatch <dir|name>",
	Short: "Stop syncing a directory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unwatch called")
	},
}

func init() {
	rootCmd.AddCommand(unwatchCmd)
	unwatchCmd.Flags().Bool("purge", false, "also delete synced data from peers")
}
