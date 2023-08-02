package models

import "time"

type UserRepository struct {
	Id                 int
	Identity           string
	UserIdentity       string
	ParentId           int
	RepositoryIdentity string
	Ext                string
	Name               string
	CreatedAt          time.Time `xorm:"created"`
	UpdatedAt          time.Time `xorm:"updated"`
	DeletedAt          time.Time `xorm:"deleted"`
}

func (table UserRepository) tableName() string {
	return "user_repository"
}
