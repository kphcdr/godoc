package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)

type Template struct {
	Id          int `json:"id"`
	UserId	int `json:"uid"`
	Title 	string `json:"template_title"`
	Content	string `json:"template_content"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"addtime"`
	Username string `json:"username" orm:"-"`
}


func (this *Template) Save() {
	o := orm.NewOrm()

	//判断是新建还是更新
	if this.Id == 0 {
		_ ,err := o.Insert(this)
		if err != nil {
			beego.Error(err.Error())
			err = fmt.Errorf("%s", "新增数据失败")
		}
	} else {
		template := Template{Id: this.Id}
		if o.Read(&template) == nil {
			template = *this
			template.Id = this.Id
			if num, err := o.Update(&template,"user_id","title","content"); err == nil {
				fmt.Println(num)
			}
		}
	}
}

func GetTemplateByUid(id int) ([]*Template) {

	o := orm.NewOrm()
	var template []*Template

	num,err :=o.QueryTable("template").Filter("user_id", id).All(&template)
	fmt.Printf("Returned Rows Num: %d, %s, %d", num, err,id)

	return template
}

func GetOneTemplate(id int) (bool,Template) {
	o := orm.NewOrm()
	template := Template{Id:id}

	err := o.Read(&template)

	if err == nil {
		return true,template
	} else {
		return false,template
	}

}

func (this *Template) Delete() (error) {
	o := orm.NewOrm()

	_,err := o.Delete(this)

	return err

}