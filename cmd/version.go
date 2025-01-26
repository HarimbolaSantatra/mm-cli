package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mm/utils"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version is", utils.GetVersion())
	},
}
