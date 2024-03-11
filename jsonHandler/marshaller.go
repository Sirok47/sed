package jsonHandler

import (
	"encoding/json"
	"fmt"
	"os"
)

type testStruct struct {
	Name string
	Age  int
	Male bool
}

func Marshal() {
	result, err := json.Marshal(testStruct{Name: "sed", Age: 333, Male: false})
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.OpenFile("S:/GoProj/json.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error trying to open file: ", err)
		return
	}
	defer file.Close()
	_, err = file.Write(result)
	if err != nil {
		fmt.Println("Error during writing to a file: ", err)
	}
}
