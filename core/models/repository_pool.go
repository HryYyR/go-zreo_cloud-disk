package models

import "time"

type RepositoryPool struct {
	Id        int
	Identity  string
	Hash      string
	Name      string
	Ext       string
	Size      int64
	Path      string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (table RepositoryPool) tableName() string {
	return "repository_pool"
}
