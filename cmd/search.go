package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "io"
    "log"
    "mm/client"
    "os"
    "encoding/json"
    "strings"
)

func init() {
    rootCmd.AddCommand(searchCmd)
}

var Reset = "\033[0m" 
var Red = "\033[31m" 
var Green = "\033[32m" 
var Yellow = "\033[33m" 
var Blue = "\033[34m" 
var Magenta = "\033[35m" 
var Cyan = "\033[36m" 
var Gray = "\033[37m" 
var White = "\033[97m"

var searchCmd = &cobra.Command{
  Use:   "search KEYWORD",
  Short: "Search for a word",
  Run: func(cmd *cobra.Command, args []string) {
      if len(args) < 1 {
	  io.WriteString(os.Stderr, "You need to enter a search keyword!")
	  os.Exit(1)
      }
      if len(args) > 1 {
	  fmt.Println("Warning: 'mm' can only handle one keyword so it will only consider the first one!")
      }
      htmlResult := client.Search(args[0])
      jsonStr := client.ParseString(htmlResult) // `jsonStr` contains a JSON

      // Decoding content of the JSON
      // we are decoding arbitrary data because the JSON contains unicode chars
      // See https://go.dev/blog/json#decoding-arbitrary-data
      var fmtResp interface{}
      jsonErr := json.Unmarshal([]byte(jsonStr), &fmtResp)
      if jsonErr != nil {
	  log.Printf("Error in JSON content:\n%s\n", string([]byte(jsonStr)))
	  log.Fatalf("Misy erreur ny JSON anao!: %s\n", jsonErr)
      }

      // Print the result
      m := fmtResp.(map[string]interface{})
      for k, v := range m {
	  trimedK := strings.TrimSpace(k) // Remove enventual whitespaces
	  switch vv := v.(type) {
	  case string:

	      // 'Morphologie' contains multiple strings
	      if(strings.Compare(trimedK, "Morphologie") != 0) {
		  fmt.Println(Green + trimedK + ": " + Reset + vv)
	      } else {
		  fmt.Println(Green + trimedK + ":", Reset)
		  fmt.Println(strings.Replace(vv, ",\n", "\n\t- ", -1))
	      }

	  default:
	      fmt.Println(k, " dia type tsy mbola hitako hatr@izay niainana!")
	  }
      }

  },
}
