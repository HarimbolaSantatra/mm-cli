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
