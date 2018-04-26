package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)

type Item struct {
	Id          int64
	Title 	string
	Description string
	UserId	int64
	Password	string
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

func (item *Item) Format() (struct{
	item_id int64
	item_name string
	item_description string
}) {
	//ret :=make(map[string]interface{})
	//println(item.Id)
	var test struct{
		item_id int64
		item_name string
		item_description string
	}

	test.item_id = 2
	test.item_name = "name"
	test.item_description = "desc"



	return test
}

func GetMyItem(uid int64) ([]*Item) {
	o := orm.NewOrm()
	var item []*Item
	num,err :=o.QueryTable("item").Filter("user_id", uid).All(&item)
	fmt.Printf("Returned Rows Num: %d, %s, %d", num, err,uid)

	return item
}