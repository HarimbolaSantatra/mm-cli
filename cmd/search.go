package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"

	"mm/client"
	"mm/utils"
)

func init() {
    rootCmd.AddCommand(searchCmd)
}


var searchCmd = &cobra.Command{
  Use:   "search KEYWORD",
  Short: "Search for a word",
  Run: func(cmd *cobra.Command, args []string) {

      if !Debug {
	  utils.PrintBanner()
      }

      if len(args) < 1 {
	  io.WriteString(os.Stderr, "You need to enter a search keyword!")
	  os.Exit(1)
      }
      if len(args) > 1 {
	  fmt.Println("Warning: 'mm' can only handle one keyword so it will only consider the first one!")
      }

      // Search the user's query to motmalgache.org
      htmlResult := client.Search(args[0])

      // extract the JSON from the html result in string format
      jsonStr := client.ParseString(htmlResult) 

      jsonContent := utils.ConvertToParsedContent([]byte(jsonStr))

      if Debug {
	  log.Printf("JSON result:\n%s\n", jsonStr)
	  jsonContent.DebugPrint()
      }

      // Print the result
      utils.PrintResult(jsonContent, Debug)

      utils.PrintRuler()

  },
}
