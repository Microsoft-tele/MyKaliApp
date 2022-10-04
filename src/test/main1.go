package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// command
	co := "ls -al"

	shell := exec.Command("/bin/zsh", "-c", co) // create a shell
	output, err := shell.Output()
	if err != nil {
		return
	}
	fmt.Printf("%s", output)
}
