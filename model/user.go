package model

import (
	"encoding/json"
	"fmt"
	db "sed/dbHandler"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint      `gorm:"primarykey" json:",omitempty"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	Name       string
	Age        int
	Male       bool
	CreditCard CreditCard `json:"-"`
}

func (u User) String() string {
	return fmt.Sprint(func(id uint) string {
		if id == 0 {
			return ""
		}
		return fmt.Sprint(u.ID, ": ")
	}(u.ID), u.Name, " ", u.Age, ", ", func(s bool) string {
		if s {
			return "male"
		}
		return "female"
	}(u.Male))
}

type Users struct {
	user []User
}

func (u Users) GetValue() []any {
	value := make([]interface{}, len(u.user))
	for i := range u.user {
		value[i] = u.user[i]
	}
	return value
}

func (u *Users) ConsumeJSON(data []byte) error {
	return json.Unmarshal(data, &u.user)
}

func (u *Users) WrapJSON() ([]byte, error) {
	return json.Marshal(&u.user)
}

func (u *Users) SendToDB() *gorm.DB {
	return db.ConnectToDB().Omit("id", "created_at", "updated_at").Create(&u.user)
}

func (u *Users) RetrieveFromDB() *gorm.DB {
	return db.ConnectToDB().Find(&u.user)
}
