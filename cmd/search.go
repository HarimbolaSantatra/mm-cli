package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "io"
    "log"
    "mm/client"
    "mm/utils"
    "os"
    "encoding/json"
    "strings"
)

func init() {
    rootCmd.AddCommand(searchCmd)
}


var searchCmd = &cobra.Command{
  Use:   "search KEYWORD",
  Short: "Search for a word",
  Run: func(cmd *cobra.Command, args []string) {
      utils.PrintBanner()
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
	      if(strings.Compare(trimedK, "Morphologie") == 0) {
		  utils.PrintLineTitle(trimedK, "")
		  fmt.Println(strings.Replace(vv, ",\n", "\n\t- ", -1))
	      } else {

		  // Remove section 'Mots composés' because it's too long!
		  // Maybe handle this in the future if I'm not lazy!
		  if(strings.Compare(trimedK, "Mots composés") != 0) {

		      // Split into multiple line for section 'Analogues'
		      if(strings.Compare(trimedK, "Analogues") == 0) {
			  utils.PrintLineTitle(trimedK, "")
			  fmt.Println(strings.Replace(vv, ",", "\n\t- ", -1))
		      } else {
			  utils.PrintLineTitle(trimedK, vv)
		      }

		  }
	      }


	  default:
	      fmt.Println(k, " dia type tsy mbola hitako hatr@izay niainana!")
	  }
      }

      utils.PrintRuler()

  },
}
