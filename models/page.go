package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/astaxie/beego"
)

type Page struct {
	Id int `json:"page_id"`
	AuthorUid int `json:"author_uid"`
	ItemId int `json:"item_id"`
	CatId int `json:"cat_id"`
	SNumber int `json:"s_number"`
	PageTitle string `json:"page_title"`
	PageContent string `json:"page_content"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"addtime"`
}

func GetPagesByItemId(id int) ([]*Page) {

	o := orm.NewOrm()
	var pages []*Page

	num,err :=o.QueryTable("page").Filter("item_id", id).All(&pages)
	fmt.Printf("Returned Rows Num: %d, %s, %d", num, err,id)

	return pages
}

func GetPagesByCatId(id int) ([]*Page) {

	o := orm.NewOrm()
	var pages []*Page

	num,err :=o.QueryTable("page").Filter("cat_id", id).All(&pages)
	fmt.Printf("Returned Rows Num: %d, %s, %d", num, err,id)

	return pages
}


func (this *Page) SavePage() {
	o := orm.NewOrm()

	//判断是新建还是更新
	if this.Id == 0 {
		_ ,err := o.Insert(this)
		if err != nil {
			beego.Error(err.Error())
			err = fmt.Errorf("%s", "新增数据失败")
		}
	} else {
		catalogs := Page{Id: this.Id}
		if o.Read(&catalogs) == nil {
			catalogs = *this
			catalogs.Id = this.Id
			if num, err := o.Update(&catalogs,"item_id","author_uid","cat_id","page_title","page_content","s_number"); err == nil {
				fmt.Println(num)
			}
		}
	}
}

func GetOnePage(id int) (bool,Page) {
	o := orm.NewOrm()
	item := Page{Id:id}

	err := o.Read(&item)

	if err == nil {
		return true,item
	} else {
		return false,item
	}

}