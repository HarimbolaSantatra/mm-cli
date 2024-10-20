package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
  Use:   "search",
  Short: "Search for a word",
  Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Search test")
  },
}
