package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print version number",
  Run: func(cmd *cobra.Command, args []string) {
      version := "0.1b"
      fmt.Println("Version is", version)
  },
}
