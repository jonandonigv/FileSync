package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch <dir>",
	Short: "Start syncing a directory",
	Long:  `A longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Watch called")
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
	watchCmd.Flags().String("name", "", "human label for this sync root (default: dirname)")
	watchCmd.Flags().String("exclude", "", "glob pattern to ignore, repeatable (--exclude '*.temp')")
	watchCmd.Flags().Bool("read-only", false, "sync inbound changes only, never push")
	watchCmd.Flags().Bool("no-delete", false, "don't propagate deleted to peers")
	watchCmd.Flags().Duration("debounce", 50*time.Millisecond, "event debounce window (default: 50ms)")
}
