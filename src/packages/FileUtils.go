package packages

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func GetFiles(folder string) (fileName []string) {
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		if file.IsDir() {
			// fmt.Println(folder + "/" + file.Name())
			fileName = append(fileName, file.Name())
			// GetFiles(folder + "/" + file.Name())   //递归
		} else {
			// fmt.Println(folder + "/" + file.Name())
			fileName = append(fileName, file.Name())
		}
	}
	return
}

func ReadFileContent(filepath string) (list []string) {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
	if err != nil {
		return
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		line = line[:len(line)-1]
		list = append(list, line)
	}
}

func WriteStringToFile(data string, filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666) //必须控制文件类型
	if err != nil {
		fmt.Println("失败")
	}
	writer := bufio.NewWriter(file)
	writer.WriteString(data + "\n")
	writer.Flush()
	//fmt.Println(data + "insert successful!")
	defer file.Close()
}
