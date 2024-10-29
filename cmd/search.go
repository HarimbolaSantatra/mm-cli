package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "io"
    "mm/client"
    "os"
)

func init() {
    rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
  Use:   "search KEYWORD",
  Short: "Search for a word",
  Run: func(cmd *cobra.Command, args []string) {
      if len(args) < 1 {
	  io.WriteString(os.Stderr, "You need to enter a search keyword!")
	  os.Exit(1)
      }
      if len(args) > 1 {
	  fmt.Println("Warning: mm can only handle one keyword so it will only consider the first one!")
      }
      client.Search(args[0])
  },
}
