package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mm/client"
)

func init() {
	rootCmd.AddCommand(healthCmd)
}

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Check health of malagasyword.org and the HTTP client",
	Run: func(cmd *cobra.Command, args []string) {
		status := client.Health()
		if status != 0 {
			fmt.Printf("Not okay! Exit code is %d", status)
		} else {
			fmt.Println("okay!")
		}
	},
}
