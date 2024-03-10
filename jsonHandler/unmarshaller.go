package jsonHandler

import (
	"encoding/json"
	"fmt"
)

func Unmarshall() {
	result := testStruct{}
	if err := json.Unmarshal([]byte("{\"Name\":\"sed\",\"Age\":333,\"Male\":false}\n"), &result); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
