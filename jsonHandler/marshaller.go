package jsonHandler

import (
	"encoding/json"
	"os"
	"sed/dbHandler"
	"sed/model"
)

func Marshal(path string, wrMode string) (result []byte, err error) {
	users := make([]model.User, 0)
	query := dbHandler.ConnectToDB().Find(&users)
	if query.Error != nil {
		err = query.Error
		return
	}
	if wrMode == "append" {
		existingJson := make([]model.User, 0)
		existingJson, err = Unmarshall(path)
		if err != nil {
			return
		}
		result, err = json.Marshal(append(existingJson, users...))
	} else {
		result, err = json.Marshal(users)
	}
	if err != nil {
		return
	}
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.Write(result)
	return
}
