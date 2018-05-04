package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Catalogs struct {
	Id int64 `json:"cat_id"`
	Name string `json:"cat_name"`
	ItemId int `json:"item_id"`
	SNumber int `json:"s_number"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"addtime"`
	ParentCatId int `json:"parent_cat_id"`
	Level int `json:"level"`
	Page []*Page `json:"page" orm:"-"`
}

func GetCatalogsByItemId(id int64) ([]*Catalogs) {

	o := orm.NewOrm()
	var catalogs []*Catalogs

	num,err :=o.QueryTable("catalogs").Filter("item_id", id).All(&catalogs)
	fmt.Printf("Returned Rows Num: %d, %s, %d", num, err,id)


	return catalogs
}