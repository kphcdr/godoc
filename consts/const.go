package consts

import (
	"github.com/astaxie/beego"
	"fmt"
)

const SESSION_UID = "uid"


func IsLogin(ctx beego.Controller) (int,error) {
	var userid int
	var err error
	uid := ctx.GetSession(SESSION_UID)

	if value,ok := uid.(int); ok {
		userid = int(value)
	} else {
		userid = 0
		err = fmt.Errorf("%s", "新增数据失败")

	}
	return userid,err

}