package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "mm/client"
    "mm/parser"
)

func init() {
    rootCmd.AddCommand(healthCmd)
}

var healthCmd = &cobra.Command{
  Use:   "health",
  Short: "Check health of motmalgache.org and the HTTP client",
  Run: func(cmd *cobra.Command, args []string) {
      status := client.Health()
      if status != 0 {
	  fmt.Println("Not okay!")
      } else {
	  // fmt.Println("okay!")
	  parser.Run()
      }
  },
}
