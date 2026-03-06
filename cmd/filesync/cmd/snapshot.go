package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "manage Raft snapshots",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("snapshot called")
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new snapshot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

var snapshotListCmd = &cobra.Command{
	Use:   "list",
	Short: "list snapshots",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

var restoreCmd = &cobra.Command{
	Use:   "restore <snapshot-id>",
	Short: "restore snapshot",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("restore called")
	},
}

func init() {
	rootCmd.AddCommand(snapshotCmd)
	restoreCmd.Flags().Bool("force", false, "restore even if node has newer state")
	createCmd.Flags().Bool("compress", false, "gzip the snapshot before writing")
	snapshotCmd.AddCommand(createCmd, snapshotListCmd, restoreCmd)
}
