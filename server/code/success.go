package code

// 正常返回数据
func NewSuccessData(msg string,data interface{}) interface{}{

	tmp_data := make([]interface{},1)
	tmp_data = append(tmp_data,data)

	return ResData{
		Code: SUCESSS_CODE,
		Msg:  msg,
		Data: tmp_data,
	}
}

func NewManualSuccessData(code int,msg string,data interface{}) interface{}{

	if code < 0 {
		code = SUCESSS_CODE
	}

	tmp_data := make([]interface{},1)
	tmp_data = append(tmp_data,data)

	return ResData{
		Code: code,
		Msg:  msg,
		Data: tmp_data,
	}
}