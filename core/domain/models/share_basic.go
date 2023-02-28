package models

import "time"

type ShareBasic struct {
	Id                     int
	Identity               string
	UserIdentity           string
	UserRepositoryIdentity string
	RepositoryIdentity     string
	ExpiredTime            int
	ClickNum               int
	CreatedAt              time.Time `xorm:"created"`
	UpdatedAt              time.Time `xorm:"updated"`
	DeletedAt              time.Time `xorm:"deleted"`
}

type ShareBasicDetail struct {
	RepositoryIdentity string `json:"repository_identity"`
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Size               int64  `json:"size"`
	Path               string `json:"path"`
}

func (table ShareBasic) TableName() string {
	return "share_basic"
}
