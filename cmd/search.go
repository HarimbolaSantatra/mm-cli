package cmd

import (
    "github.com/spf13/cobra"
    "mm/client"
)

func init() {
    rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
  Use:   "search",
  Short: "Search for a word",
  Run: func(cmd *cobra.Command, args []string) {
      client.TestIt()
  },
}
