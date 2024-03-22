package jsonHandler

import (
	"encoding/json"
	"log"
	"os"
	"sed/dbHandler"
	"sed/model"
)

func Unmarshall(path string, stamp bool) (result []model.User, err error) {
	result = make([]model.User, 0)
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	fileStat, _ := file.Stat()
	input := make([]byte, fileStat.Size())
	_, err = file.Read(input)
	if err != nil {
		return
	}
	err = json.Unmarshal(input, &result)
	if err != nil {
		return
	}
	if stamp {
		transaction := dbHandler.ConnectToDB().Create(&result)
		if transaction.Error != nil {
			log.Fatal(transaction.Error)
		}
	}
	return
}
