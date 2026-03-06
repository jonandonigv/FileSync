/*
Copyright © 2026 Jon Andoni Galdos <jonandonigv@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "FileSync",
	Short: "FileSync is a distributed file synchronization system written in Go that keeps directories in sync across multiple nodes with strong consistency guarantees.",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.FileSync.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.PersistentFlags().String("node-id", "", "override node Id (defualt: hostname)")
	rootCmd.PersistentFlags().String("log-format", "", "text | json (default: text)")
	rootCmd.PersistentFlags().String("log-level", "", "debug | info | warn | error (default: info)")
	rootCmd.PersistentFlags().String("data-dir", "", "directory for node data, logs, snapshots (~/.filesync/data)")
	rootCmd.PersistentFlags().String("config", "", "path to config file (defaul: ~/.filesync/config.yaml)")
}
