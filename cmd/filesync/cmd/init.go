/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var files []string

// addCmd represents the add command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Files added to sync: %v \n", files)
		if len(files) < 1 {
			fmt.Println("No files added")
		} else {
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().String("listen", "l", "address to bind on (default: 0.0.0.0:7946)")
	initCmd.Flags().String("join", "j", "address of existing node to join (omit to start new cluster)")
	initCmd.Flags().String("peer", "p", "static peer address, repeatable (--peer a:port --peer b:port)")
	initCmd.Flags().Bool("force", false, "overwrite existing data directory")
	initCmd.Flags().StringSliceVarP(&files, "files", "f", files, "Add files to be sync")
	/* For the Dir flag it should take the specified directory or the CWD if not set  */
	/* 	addCmd.Flags().String() */
}
