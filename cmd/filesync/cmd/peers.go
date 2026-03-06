package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var peersCmd = &cobra.Command{
	Use:   "peers",
	Short: "manage cluster membership",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list members of a cluster",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

var addCmd = &cobra.Command{
	Use:   "add <address>",
	Short: "add a member to the cluster",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}
var removeCmd = &cobra.Command{
	Use:   "remove <node-id>",
	Short: "remove a member to the cluster",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

func init() {
	peersCmd.AddCommand(listCmd, addCmd, removeCmd)
	listCmd.Flags().Bool("json", false, "machine-readable output")
	removeCmd.Flags().Bool("force", false, "remove even if node is currently reachable")
	rootCmd.AddCommand(peersCmd)
}
