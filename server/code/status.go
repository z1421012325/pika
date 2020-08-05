package code

// 状态码
const (
	// 成功
	SUCESSS_CODE = 10000
	// 通用失败
	FAILED_CODE = 10001
	// 参数错误
	PARAM_BIND_CODE = 20001
	// 添加错误
	ADD_ERR_CODE  = 10007
	// 更新异常
	UPLOAD_ERR_CODE = 10008
	// 删除异常
	DEL_ERR_code = 10009
)

// 异常信息,错误信息,成功信息
const (
	USERNAME_REPEAT_ERR   = "账号已存在"
	USERNAME_PASSWORD_ERR = "账号或者密码错误"
	PARAM_BIND_ERR        = "参数异常"
	LOGIN_ERR			  = "登录异常"
	REGISTRY_ERR 		  = "注册异常"
	SEARCH_ERR			  = "查询异常"
)



// 返回数据模型
type ResData struct {
	Code int    `json:"code,omitempy"`
	Msg  string `json:"msg,omitempy"`
	Err  string `json:"err,omitempy"`
	// omitempy 该列数据没有则json时不序列化,string表示序列化后显示为字符串
	Data []interface{} `json:"data,omitempy,string"`
}









type DataMap struct {
	maps map[string]interface{}
}

func (s *DataMap) InitSet(k string, v interface{}) {
	s.maps = make(map[string]interface{})
	s.maps[k] = v
}

func (s *DataMap) Set(k string, v interface{}) {
	s.maps[k] = v
}

func (s *DataMap) Get(k string) interface{} {
	v, ok := s.maps[k]
	if !ok {
		return nil
	}
	return v
}

func (s *DataMap) GetMap() interface{} {
	return s.maps
}
