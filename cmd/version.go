/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version string
	commit  string
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "prints current nft-generator version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version: ", version)
		fmt.Println("commit: ", commit)
	},
}
