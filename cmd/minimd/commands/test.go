package cmd

import (
	"github.com/spf13/cobra"
	"github.com/minichain/blockchainDemo"
)

var testCmd = &cobra.Command{
	Use:   "testCmd",
	Short: "test based code",
	Run: func(cmd *cobra.Command, args []string) {
		c := new (blockchainDemo.CLI)
		c.Run()
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}

