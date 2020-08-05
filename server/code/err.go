package code

import "github.com/gin-gonic/gin"

// 错误,异常返回
func NewFailedData(err string) ResData {
	return ResData{
		Code: FAILED_CODE,
		Err:  err,
	}
}

func NewManualFailedData(code int, msg, err string) ResData {

	if code < 0 {
		code = FAILED_CODE
	}

	return ResData{
		Code: code,
		Msg:  msg,
		Err:  err,
	}
}

// 参数错误
func ParamBindErrorResult() ResData {
	return ResData{
		Code: PARAM_BIND_CODE,
		Msg:  PARAM_BIND_ERR,
	}
}

func GinParamBindErrorResult(c *gin.Context) {
	c.JSON(201,ParamBindErrorResult())
}


// 账号密码错误
func PasswordErrorResult() ResData{
	return ResData{
		Code: FAILED_CODE,
		Msg:  USERNAME_PASSWORD_ERR,
	}
}

// 登录,token错误
func LoginErrorResult() ResData{
	return ResData{
		Code: FAILED_CODE,
		Err:LOGIN_ERR,
	}
}

// 账号已存在
func UserRepeatErrorResult() ResData{
	return ResData{
		Code: FAILED_CODE,
		Msg:USERNAME_REPEAT_ERR,
	}
}

// 注册失败
func UserRegistryErrorResult() ResData{
	return ResData{
		Code: FAILED_CODE,
		Msg: REGISTRY_ERR,
	}
}

// 查询异常
func SearchErrorResult() ResData {
	return ResData{
		Code: FAILED_CODE,
		Err: SEARCH_ERR,
	}
}

// 添加异常
func AddErrorResult(msg ,err string) ResData {
	return ResData{
		Code: ADD_ERR_CODE,
		Msg:  msg,
		Err:  err,
		Data: nil,
	}
}

// 删除异常
func DelErrorResult(msg,err string) ResData {
	return ResData{
		Code: DEL_ERR_code,
		Msg:  msg,
		Err:  err,
	}
}

// 操作异常
func OperationErrorResult(msg,err string) ResData {
	return ResData{
		Code: FAILED_CODE,
		Msg:  msg,
		Err:  err,
	}
}