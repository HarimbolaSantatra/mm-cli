package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "mm",
	Short: "mm is a cli client for search malagasy words and proverbs",
	Long:  `A client for searching malagasy words, proverbs and misc`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var Debug bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Use debug mode")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
