package models

import (
	"time"
)

type Item struct {
	Id          int64
	Title 	string
	Description string
	UserId	int64
	Password	string
	LastUpdateTime int
	Type	int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}