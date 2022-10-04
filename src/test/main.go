package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	co := "ls"
	shell := exec.Command("/bin/bash", "-c", co)
	fmt.Println(shell.Args)
	//stdout, err := shell.StdoutPipe()
	shell.Stdout = os.Stdout
	shell.Stdin = os.Stdin
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	shell.Run()
	//reader := bufio.NewReader(shell.Stdout)
	//for {
	//	line, err2 := reader.ReadString('\n')
	//	if err2 != nil || io.EOF == err2 {
	//		break
	//	}
	//	fmt.Println(line)
	//}
	shell.Wait()
}

func ExecCommand(command string) {
	shell := exec.Command("/bin/bash", "-c", command)
	fmt.Println(shell.Args)
	shell.Stdout = os.Stdout
	shell.Stdin = os.Stdin
	shell.Run()
	shell.Wait()
}
