package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
	"time"
)

type User struct {
	Id          int64
	Email        string
	Password	string
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}


func (base *User) Create() (int64,error) {
	o := orm.NewOrm()

	id ,err := o.Insert(base)
	if err != nil {
		beego.Error(err.Error())
		err = fmt.Errorf("%s", "新增数据失败")

	}

	return id,err
}

func CryptPassword(password string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}