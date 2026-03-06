/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "bootstrap a new node/cluster",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Init called")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().String("listen", "l", "address to bind on (default: 0.0.0.0:7946)")
	initCmd.Flags().String("join", "j", "address of existing node to join (omit to start new cluster)")
	initCmd.Flags().String("peer", "p", "static peer address, repeatable (--peer a:port --peer b:port)")
	initCmd.Flags().Bool("force", false, "overwrite existing data directory")
}
