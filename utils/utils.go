package utils

import (
    "log"
    "encoding/json"

    "mm/datatype"
)

// Get the keys and the values of a JSON and returns two slice representing them
// Takes an empty interface representing a JSON data
// Returns a slice of keys and values
func GetKeysValues( jsonContent interface{}, debug bool) (keys, values []string) {
    x := make([]string, 0)
    y := make([]string, 0)

    for k, v := range jsonContent.(map[string]string) {
	if debug {
	    log.Printf("Append key-value %s-%s", k, v)
	}
	x = append(x, k)
	y = append(y, v)
    }

    if debug {
	log.Printf("Value of keys and value in GetKeysValues:\n%s\n%s", x, y)
	log.Printf("Len of keys: %d", len(x))
	log.Printf("Len of values: %d", len(y))
    }

    return x,y
}

// Convert the json string in ParsedContent
func ConvertToParsedContent(jsonStr []byte) datatype.ParsedContent {
    var jsonContent datatype.ParsedContent
    var m map[string]string

    err := json.Unmarshal(jsonStr, &m)
    if err != nil {
	log.Printf("JSON in string format: \"%s\"", jsonStr)
	log.Fatalf("JSON Unmarshaling error: \"%s\"", err)
    }

    // Using the map `m` as intermediary is necessary because
    // encoding/json package does not allow some characters as a key
    // See: https://stackoverflow.com/questions/79330361/cannot-use-some-unicode-character-as-a-struct-tag/79330696
    jsonContent.Speech = m["Part of speech\u00a0"]
    jsonContent.MlgExplication = m["MlgExplication\u00a0"]
    jsonContent.FrExplication = m["FrExplication\u00a0"]
    jsonContent.Vocabulary = m["Vocabulary\u00a0"]
    jsonContent.Morphology = m["Morphology\u00a0"]
    jsonContent.Analogs = m["Analogs\u00a0"]

    jsonContent.Json = string(jsonStr)

    return jsonContent
}
