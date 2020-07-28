package code


// 错误,异常返回
func NewFailedData(err string) interface{}{
	return ResData{
		Code: FAILED_CODE,
		Err:  err,
	}
}

func NewManualFailedData(code int,msg,err string) interface{}{

	if code < 0 {
		code = FAILED_CODE
	}

	return ResData{
		Code: code,
		Msg:  msg,
		Err:  err,
	}
}