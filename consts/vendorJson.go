package consts

type Json struct {
	message string
	code int
	data interface{}
}



func (u *Json) SetMessage(message string) {
	u.message = message
}

func (u *Json) SetCode(code int) {
	u.code = code
}

func (u *Json) SetData(data interface{}) {
	u.data = data
}

func (u *Json) Set(code int,message string) {
	u.SetMessage(message)
	u.SetCode(code)
}

func (u *Json) VendorError() interface{} {
	return map[string]interface{}{"error_code":u.code,"error_message":u.message}
}

func (u *Json) VendorOk() interface{} {
	return map[string]interface{}{"error_code":0,"data":u.data,"message":u.message}
}
