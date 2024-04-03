package model

import (
	"encoding/json"
	db "sed/dbHandler"

	"gorm.io/gorm"
)

type CreditCard struct {
	ID     uint `gorm:"primarykey" json:",omitempty"`
	Number int
	UserID uint
}

type CreditCards struct {
	creditCard []CreditCard
}

func (c CreditCards) GetValue() []any {
	value := make([]interface{}, len(c.creditCard))
	for i := range c.creditCard {
		value[i] = c.creditCard[i]
	}
	return value
}

func (c *CreditCards) ConsumeJSON(data []byte) error {
	return json.Unmarshal(data, &c.creditCard)
}

func (c *CreditCards) WrapJSON() ([]byte, error) {
	return json.Marshal(&c.creditCard)
}

func (c *CreditCards) SendToDB() *gorm.DB {
	return db.ConnectToDB().Create(&c.creditCard)
}

func (c *CreditCards) RetrieveFromDB() *gorm.DB {
	return db.ConnectToDB().Find(&c.creditCard)
}
