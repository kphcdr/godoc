package models

import "time"

type Catalogs struct {
	Id int64 `json:"cat_id"`
	Name string `json:"cat_name"`
	ItemId int `json:"item_id"`
	SNumber int `json:"s_number"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)",json:"addtime"`
	ParentCatId int `json:"parent_cat_id"`
	Level int `json:"level"`
	Page []Page `json:"page"`
	Catalogs []Catalogs `json:"catalogs"`
}