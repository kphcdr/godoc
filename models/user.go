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
	Id          int	`json:"id"`
	Email        string `json:"email"`
	Password	string	`json:"-"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"addtime"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func GetOneUser(id int) (bool,User) {
	o := orm.NewOrm()
	user := User{Id:id}

	err := o.Read(&user)

	if err == nil {
		return true,user
	} else {
		return false,user
	}

}

func (base *User) Create() (int,error) {
	o := orm.NewOrm()

	id ,err := o.Insert(base)
	if err != nil {
		beego.Error(err.Error())
		err = fmt.Errorf("%s", "新增数据失败")

	}

	return int(id),err
}

func Login(email string,password string) (bool,User) {
	o := orm.NewOrm()
	var user User

	err := o.QueryTable(user).Filter("Email", email).Filter("Password", password).One(&user)
	return err != orm.ErrNoRows, user
}

func CryptPassword(password string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func (this *User) Save() {
	o := orm.NewOrm()

	//判断是新建还是更新
	if this.Id == 0 {
		_ ,err := o.Insert(this)
		if err != nil {
			beego.Error(err.Error())
			err = fmt.Errorf("%s", "新增数据失败")
		}
	} else {
		user := User{Id: this.Id}
		if o.Read(&user) == nil {
			user = *this
			user.Id = this.Id
			if num, err := o.Update(&user,"password","email"); err == nil {
				fmt.Println(num)
			}
		}
	}
}