package code

// 错误,异常返回
func NewFailedData(err string) interface{} {
	return ResData{
		Code: FAILED_CODE,
		Err:  err,
	}
}

func NewManualFailedData(code int, msg, err string) interface{} {

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
func ParamBindErrorResult() interface{} {
	return ResData{
		Code: PARAM_BIND_CODE,
		Msg:  PARAM_BIND_ERR,
	}
}


// 账号密码错误
func PasswordErrorResult() interface{}{
	return ResData{
		Code: FAILED_CODE,
		Msg:  USERNAME_PASSWORD_ERR,
	}
}

// 登录,token错误
func LoginErrorResult() interface{}{
	return ResData{
		Code: FAILED_CODE,
		Err:LOGIN_ERR,
	}
}

// 账号已存在
func UserRepeatErrorResult() interface{}{
	return ResData{
		Code: FAILED_CODE,
		Msg:USERNAME_REPEAT_ERR,
	}
}

// 注册失败
func UserRegistryErrorResult() interface{}{
	return ResData{
		Code: FAILED_CODE,
		Msg: REGISTRY_ERR,
	}
}