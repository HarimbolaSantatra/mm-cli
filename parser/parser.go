package parser

import (
    "log"
    "os/exec"
)

func ParseString(s string) string {
    cmd := exec.Command("./scraping", s)
    bt, err := cmd.Output()
    if err != nil {
	log.Fatalf("Command execution error: ", err)
    }
    return string(bt)
}
