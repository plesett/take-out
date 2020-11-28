package tool

type SuccessMessageStruct struct {
	Code	 int
	Msg		 string
	Date	 interface{}
}

type ErrorMessageStruct struct {
	Code	 int 			`json:"code"`
	Msg		 string			`json:"msg"`
	Date	 interface{}	`json:"date"`
}

const (
	SuccessSendCode = "验证码已成功发送"
	SuccessSubmit   = "提交成功"
	ErrorMobile		= "手机号码错误"
	ErrorMessage	= "信息填写错误"
	ErrorMessageDefect	= "信息请填写完整"
)

func ErrorReturnForMat(msg string, data interface{}) ErrorMessageStruct {
	return ErrorMessageStruct{
		Code: -1,
		Msg: msg,
		Date: data,
	}
}


