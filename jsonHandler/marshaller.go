package jsonHandler

import (
	"encoding/json"
	"log"
	"os"
	"sed/dbHandler"
	"sed/model"
)

func fileWrite(path string, value []byte) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.Write(value)
	if err != nil {
		log.Fatal(err)
	}
}

func Marshal(path string, wrMode string) (result []byte, err error) {
	users := make([]model.User, 0)
	query := dbHandler.ConnectToDB().Select("name", "age", "male").Find(&users)
	if query.Error != nil {
		err = query.Error
		return
	}
	if wrMode == "append" {
		existingJson := make([]model.User, 0)
		existingJson, err = Unmarshall(path, false)
		if err != nil {
			return
		}
		result, err = json.Marshal(append(existingJson, users...))
		fileWrite(path, result)
	} else if wrMode == "rewrite" {
		result, err = json.Marshal(users)
		fileWrite(path, result)
	}
	if err != nil {
		return
	}
	return
}
