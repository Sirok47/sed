package jsonHandler

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(string(result))
}
