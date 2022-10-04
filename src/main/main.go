package main

import (
	"App/src/packages"
	"fmt"
	"os"
)

/*
#include<stdio.h>
#include<stdlib.h>
#include<string.h>
int gdb(){
// 	FILE* fp;
// 	char buffer[100];
// 	fp = popen("pwd", "r");
// 	fgets(buffer, sizeof(buffer), fp);
// 	char cmd[100] = "python3 ";
// 	strcat(cmd, buffer);

// 	cmd[strlen(cmd) - 1] = 0;

// 	strcat(cmd, "/gdbinit.py");
// 	printf("%s", cmd);
	system("python3 /home/kali/Desktop/git/gdbinit/gdbinit.py");
	return 0;
}
*/
import "C"

func main() {
	for {
		Menu()
	}
}

func Menu() {
	var choice int
	fmt.Println("1.	Install")
	fmt.Println("2.	Initial gdb")
	fmt.Println("3.	Scp")
	fmt.Println("0.	Exit")
	fmt.Printf("Input:")
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		install()
	case 2:
		gdbinit()
	case 3:
		scp()
	case 0:
		os.Exit(1)
	}
}

func install() {
	packages.InstallAllPkgs(packages.GetAllPkgs())
	// Add : choose install function
}
func gdbinit() {
	//packages.ExecCommandBinBash("python gdbinit.py")
	C.gdb()
}
func scp() {
	packages.ScpMenu()
}
