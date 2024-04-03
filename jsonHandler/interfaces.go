package jsonHandler

import "gorm.io/gorm"

type JSONable interface {
	GetValue() []any

	ConsumeJSON(data []byte) error
	WrapJSON() ([]byte, error)

	SendToDB() *gorm.DB
	RetrieveFromDB() *gorm.DB
}
