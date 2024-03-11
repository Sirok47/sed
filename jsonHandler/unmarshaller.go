package jsonHandler

import (
	"encoding/json"
	"fmt"
	"os"
)

func Unmarshall() {
	//result := make([]testStruct, 0)
	result := testStruct{}
	file, err := os.Open("S:/GoProj/json.txt")
	if err != nil {
		fmt.Println("Error trying to open file: ", err)
		return
	}
	defer file.Close()
	fileStat, _ := file.Stat()
	input := make([]byte, fileStat.Size())
	_, err = file.Read(input)
	if err != nil {
		fmt.Println("Error during reading file: ", err)
		return
	}
	fmt.Println(string(input))
	if err := json.Unmarshal(input, &result); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
