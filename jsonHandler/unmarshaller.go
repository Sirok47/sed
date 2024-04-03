package jsonHandler

import (
	"log"
	"os"
)

func Unmarshal(object *JSONable, path string, stamp bool) (err error) {
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

	err = (*object).ConsumeJSON(input)
	if err != nil {
		return
	}
	if !stamp {
		return
	}

	transaction := (*object).SendToDB()
	if transaction.Error != nil {
		log.Fatal(transaction.Error)
	}

	return
}
