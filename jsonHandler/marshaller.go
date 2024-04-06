package jsonHandler

import (
	"encoding/json"
	"log"
	"os"
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

func Marshal(object *JSONable, path string, wrMode string) (result []byte, err error) {
	query := (*object).RetrieveFromDB()
	if query.Error != nil {
		err = query.Error
		return
	}
	if wrMode == "append" {
		var existingJson JSONable
		switch (*object).(type) {
		case *model.Users:
			existingJson = &model.Users{}
		case *model.CreditCards:
			existingJson = &model.CreditCards{}
		}
		err = Unmarshal(&existingJson, path, false)
		if err != nil {
			return
		}
		result, err = json.Marshal(append(existingJson.GetValue(), (*object).GetValue()...)) //переделать без передачи интерфейсов
		fileWrite(path, result)
	} else if wrMode == "rewrite" {
		result, err = json.Marshal((*object).GetValue())
		fileWrite(path, result)
	} else {
		result, err = json.Marshal((*object).GetValue())
	}
	return
}
