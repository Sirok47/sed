package model

import (
	"fmt"
	"time"
)

type User struct {
	ID        uint      `gorm:"primarykey" json:",omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Name      string
	Age       int
	Male      bool
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
