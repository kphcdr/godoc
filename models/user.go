package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)

type User struct {
	Id          int
	Name        string
	Password	string
}


func (base *User) Create() error {
	o := orm.NewOrm()

	_,err := o.Insert(base)
	if err != nil {
		beego.Error(err.Error())
		err = fmt.Errorf("%s", "新增数据失败")

	}

	return err
}

func CryptPassword(password string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}