package models

import "time"

type UserBasic struct {
	Id        int       `xorm:"id"`
	Identity  string    `xorm:"identity"`
	Name      string    `xorm:"name"`
	Password  string    `xorm:"password"`
	Email     string    `xorm:"email"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (table UserBasic) tableName() string {
	return "user_basic"
}
