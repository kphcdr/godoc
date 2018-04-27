package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)

type Item struct {
	Id          int64 `json:"item_id"`
	Title 	string `json:"item_name"`
	Description string `json:"item_description"`
	UserId	int64
	Password	string `json:"-"`
	Type	int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func (this *Item) Create() (int64,error) {
	o := orm.NewOrm()

	id ,err := o.Insert(this)
	if err != nil {
		beego.Error(err.Error())
		err = fmt.Errorf("%s", "新增数据失败")
	}

	return id,err
}

func GetMyItem(uid int64) ([]*Item) {
	o := orm.NewOrm()
	var item []*Item
	num,err :=o.QueryTable("item").Filter("user_id", uid).All(&item)
	fmt.Printf("Returned Rows Num: %d, %s, %d", num, err,uid)

	return item
}