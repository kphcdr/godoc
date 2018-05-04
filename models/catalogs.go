package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
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
	Page []*Page `json:"pages" orm:"-"`
}

func GetOneCataLogs(id int64) (bool,Catalogs) {
	o := orm.NewOrm()
	catalogs := Catalogs{Id:id}

	err := o.Read(&catalogs)

	if err == nil {
		return true,catalogs
	} else {
		return false,catalogs
	}

}

func (this *Catalogs) Save() {
	o := orm.NewOrm()

	//判断是新建还是更新
	if this.Id == 0 {
		_ ,err := o.Insert(this)
		if err != nil {
			beego.Error(err.Error())
			err = fmt.Errorf("%s", "新增数据失败")
		}
	} else {
		catalogs := Catalogs{Id: this.Id}
		if o.Read(&catalogs) == nil {
			catalogs = *this
			catalogs.Id = this.Id
			if num, err := o.Update(&catalogs,"item_id","name","s_number","parent_cat_id","level"); err == nil {
				fmt.Println(num)
			}
		}
	}
}

func (this *Catalogs) Delete() (error) {
	o := orm.NewOrm()


	 _,err := o.Delete(this)

	 return err

}

func GetCatalogsByItemId(id int64) ([]*Catalogs) {

	o := orm.NewOrm()
	var catalogs []*Catalogs

	num,err :=o.QueryTable("catalogs").Filter("item_id", id).All(&catalogs)
	fmt.Printf("catalogs Returned Rows Num: %d, %s, %d", num, err,id)

	for _,value := range catalogs {
		value.Page = GetPagesByCatId(value.Id)
	}


	return catalogs
}