package handlers

import (

	r "pika/server/code"

	"github.com/gin-gonic/gin"
)


type UserLoginReqData struct {
	Username string		`json:"username" form:"username" binding:"required,min=2,max=20"`
	PassWord string		`json:"password" form:"password" binding:"required,min=5,max=30"`
	// 登录等级
	Grade int			`json:"grade" form:"grade" binding:"required"`
}
// 登录
func UserLogin(c *gin.Context){

	var req UserLoginReqData
	if err := c.ShouldBind(&req); err != nil{
		c.JSON(201,r.NewManualFailedData(r.FAILED_CODE,"",r.PARAM_BIND_ERR))
		return
	}

	/*
		逻辑
			数据库查询
			密码校正
			返回token(cookie,header中都存入),暂不做一个用户多个token过期限制
	 */
	
	
	
	
	var token string
	//var data = make(map[string]interface{})
	//data["token"] = token
	var data r.DataMap
	data.InitSet("token",token)

	c.JSON(
		200,
		r.NewManualSuccessData(r.SUCESSS_CODE,"login success",data.GetMap()))
}








// -------------------------------------------------------------------------------------

type UserRegistryReqData struct {
	Username string	`json:"username" form:"username" binding:"required,min=2,max=20"`
	PassWord string	`json:"password" form:"password" binding:"required,min=5,max=30"`
	Nickname string	`json:"nickname" form:"nickname" binding:"required"`
	Grade 	int 	`json:"grade" form:"grade"`
}

func UserRegistry(c *gin.Context){
	var req UserRegistryReqData
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(201,r.NewManualFailedData(r.FAILED_CODE,"",r.PARAM_BIND_ERR))
		return
	}

	/*
		逻辑
			查询账号是否重复
			密码加密
			用户等级校正
			添加数据库
	 */

	// 返回数据
}











// --------------------------------------------------------------------------


type UserCollectionReqData struct {
	Token string  `json:"token"`
}

type UserCollectionResData struct {
	// models
}

// /user/collection 查询收藏本子
func UserCollection(c *gin.Context){

	/*
		用户id在token中或者在headers中
	 */

}





// ----------------------------/user/commit	 查询用户评论------------------------------------
type UserCommitsReqData struct {
	UserCollectionReqData
}
type UserCommitsResData struct {
}
func UserCommits(c *gin.Context){

	/*
		用户id在token中或者在headers中
	*/

}



// ----------------------/user/commit/reply 查询层级评论的回复评论------------------------
type UserCommitReplyReqData struct {
	UserCollectionReqData
	CId 	int `json:"cid" form:"cid" binding:"required"`
}
type UserCommitReplyResData struct {
}
func UserCommitReply(c *gin.Context){

	/*
		用户id在token中或者在headers中
	*/

}


// ----------------------/user/info 	查询用户信息------------------------
type UserInfoReqData struct {
	UserCollectionReqData
}
type UserInfoResData struct {
}
func UserInfo(c *gin.Context){

	/*
		用户id在token中或者在headers中
	*/

}



// ----------------------/user/out	退出登录------------------------
type UserOutReqData struct {
	UserCollectionReqData
}
type UserOutResData struct {
}
func UserOut(c *gin.Context){

	/*
		用户id在token中或者在headers中
	*/

}

