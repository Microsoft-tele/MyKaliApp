package packages

import (
	"encoding/json"
	"fmt"
)

func ReadConveyFile() {
	user := User{
		Name: "Go",
		File: nil,
	}
	user.CreateFile()
	user.ReadConveyFile()
}
func ShowAccount() {
	hostStr := ReadFileContent("./account.json")
	for i, v := range hostStr {
		fmt.Printf("[%v : %v]\n", i, v)
	}
}

func ScpMenu() bool {
	fmt.Println("1.	Insert account")
	fmt.Println("2.	Read convey file")
	fmt.Println("3.	Show account")
	fmt.Println("4.	Convey")
	fmt.Println("5.	exit")
	fmt.Println("Please input:")
	var choice int
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		host := CreateHost()
		host.InsertAccount()
	case 2:
		ReadConveyFile()
	case 3:
		ShowAccount()
	case 4:
		var HostObjSlice []Host
		HostObjSlice = make([]Host, 0)
		HostJsonSlice := ReadFileContent("./account.json")
		for _, v := range HostJsonSlice {
			// fmt.Printf("[%v : %v]\n", i, v)
			var obj Host
			json.Unmarshal([]byte(v), &obj)
			HostObjSlice = append(HostObjSlice, obj)
		}
		fmt.Println("Already exists account:")
		for i, v := range HostObjSlice {
			fmt.Printf("[%v : %v]\n", i, v)
		}
		fmt.Println("Input:")
		var choice int
		fmt.Scanf("%d", &choice)
		localhost := User{
			Name: "localhost",
			File: nil,
		}
		localhost.ConveyFile(&HostObjSlice[choice])
	case 5:
		return true
	}
	return true
}
