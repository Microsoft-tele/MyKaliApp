package packages

import (
	"fmt"
)

func AptInstall(pkg string) {
	fmt.Println("Installing " + pkg + " ----------------------------------------------------------------------")
	ExecCommandBinBash("sudo apt -y install " + pkg)
}

func GetAllPkgs() (pkgs []string) {
	pkgs = ReadFileContent("./pkgs.list")
	return
}

func InstallAllPkgs(pkgs []string) {
	for _, v := range pkgs {
		AptInstall(v)
	}
}
