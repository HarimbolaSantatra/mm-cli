package parser

import (
    "fmt"
    "os/exec"
)

func Run() {
    cmd := exec.Command("ls", "-lh")
    if err := cmd.Run(); err != nil {
	fmt.Println("Error: ", err)
    }
}
