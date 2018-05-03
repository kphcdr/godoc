package models

import "time"

type Page struct {
	Id int `json:"page_id"`
	AuthorUid int `json:"author_uid"`
	ItemId int `json:"item_id"`
	CatId int `json:"cat_id"`
	PageTitle string `json:"page_title"`
	PageContent string `json:"page_content"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)",json:"addtime"`
}

func GetPagesByItemId(id int64) ([]Page) {

	println(id)

	var pages []Page

	return pages
}