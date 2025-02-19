
package cmd

import (
	"github.com/spf13/cobra"

	"mm/client"
)

func init() {
	rootCmd.AddCommand(browseCmd)
}

var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "Browse",
	Run: func(cmd *cobra.Command, args []string) {
                client.OpenBrowser()
	},
}
