package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "inspect and edit config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show config",
}

var setCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set config",
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validates the config",
}

func init() {
	configCmd.AddCommand(showCmd, setCmd, validateCmd)
	rootCmd.AddCommand(configCmd)
}
