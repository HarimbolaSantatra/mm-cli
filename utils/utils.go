package utils

import (
    "log"
)

// Get the keys and the values of a JSON and returns two slice representing them
// Arguments: and empty interfacem representing a JSON data
// Returns: an slice of keys and values
func GetKeysValues( jsonContent interface{}, debug bool) (keys, values []string) {
    x := make([]string, 0)
    y := make([]string, 0)
    x = append(x, "Name","Last name", "age")
    y = append(y, "Richard", "Rapport", "28")

    if debug {
	log.Printf("Value of keys and value in GetKeysValues:\n%s\n%s", x, y)
	log.Printf("Len of keys: %d", len(x))
	log.Printf("Len of values: %d", len(y))
    }

    return x,y
}
