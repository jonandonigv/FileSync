package cmd

import "github.com/spf13/cobra"

var syncCmd = &cobra.Command{
	Use:   "sync <dir>",
	Short: "manually trigger a one-shot sync (no persistent watch)",
}

func init() {
	rootCmd.AddCommand(syncCmd)
	syncCmd.Flags().Bool("dry-run", false, "show what would be transferred without doin it")
	syncCmd.Flags().String("direction", "both", "push | pull | both (default: both)")
	syncCmd.Flags().String("chunk-size", "4KB", "override target chunk size (default: 4KB)")
}
