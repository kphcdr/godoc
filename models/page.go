package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Page struct {
	Id int `json:"page_id"`
	AuthorUid int `json:"author_uid"`
	ItemId int `json:"item_id"`
	CatId int `json:"cat_id"`
	PageTitle string `json:"page_title"`
	PageContent string `json:"page_content"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"addtime"`
}

func GetPagesByItemId(id int64) ([]*Page) {

	o := orm.NewOrm()
	var pages []*Page

	num,err :=o.QueryTable("page").Filter("item_id", id).All(&pages)
	fmt.Printf("Returned Rows Num: %d, %s, %d", num, err,id)

	return pages
}

func GetPagesByCatId(id int64) ([]*Page) {

	o := orm.NewOrm()
	var pages []*Page

	num,err :=o.QueryTable("page").Filter("cat_id", id).All(&pages)
	fmt.Printf("Returned Rows Num: %d, %s, %d", num, err,id)

	return pages
}