package handlers

import (
	"fmt"
	r "pika/server/code"
	model "pika/server/models"
	"pika/tools"
	"pika/config"

	"github.com/gin-gonic/gin"
)


type SearchPageReqData struct {
	Page int64			`json:"page" form:"page" binding:"default=1"`
	Number int64			`json:"number" form:"number" binding:"default=20"`
}


// -------------------------------/user/login --------------------------------------------

type UserLoginReqData struct {
	//Username string		`json:"username" form:"username" binding:"required,min=2,max=20"`
	//PassWord string		`json:"password" form:"password" binding:"required,min=5,max=30"`
	//// 登录等级
	//Grade int			`json:"grade" form:"grade" binding:"required"`

	// 参数字段 username,paswd,grade
	model.User
}
// 登录
func UserLogin(c *gin.Context){

	var req UserLoginReqData
	if err := c.ShouldBind(&req); err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	/*
		逻辑
			数据库查询
			密码校正
			返回token(cookie,header中都存入),暂不做一个用户多个token过期限制
	 */

	if !req.VerifyUser(req.PassWord) {
		c.JSON(201,r.PasswordErrorResult())
		return
	}

	tokenStr,ok := tools.NewToken(req.Uid)
	if !ok {
		c.JSON(201,r.LoginErrorResult())
		return
	}
	//var data = make(map[string]interface{})
	//data["token"] = token
	var data r.DataMap
	data.InitSet(config.SET_TOKEN_NAME,tokenStr)

	c.JSON(
		200,
		r.NewManualSuccessData(r.SUCESSS_CODE,"login success",data.GetMap()))
}








// --------------------------------/user/registry 用户注册---------------------------------

type UserRegistryReqData struct {
	//Username string	`json:"username" form:"username" binding:"required,min=2,max=20"`
	//PassWord string	`json:"password" form:"password" binding:"required,min=5,max=30"`
	Nickname string	`json:"nickname" form:"nickname" binding:"required,min=3,max=30"`
	//Grade 	int 	`json:"grade" form:"grade"`

	model.User
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
			用户等级校正  -> pass 看能不能通过tag过滤 oneof
			添加数据库
	 */
	req.QueryUser()
	if req.Uid != 0 {
		c.JSON(201,r.UserRepeatErrorResult())
		return
	}

	req.Encryption()
	req.CheckUserGrade()
	if req.RegistryUser(req.Nickname) != nil {
		c.JSON(201,r.UserRegistryErrorResult())
		return
	}

	// 返回数据
	c.JSON(200,r.NewSuccessData("registry success",nil))
}











// --------------------/user/collection 查询收藏本子------------------------------


type UserCollectionReqData struct {
	// 拦截器拦截token解析并在header中设置了UID 试试结构体中用gin的 ShouldBind 映射
	//Uid int 	`json:"UID" form:"UID" binding:"required"`
	SearchPageReqData
}

type UserCollectionResData struct {
	// models
	model.BenZi
	model.Collection
}

// /user/collection 查询收藏本子
func UserCollection(c *gin.Context){
	/*
		用户id在token中或者在headers中
	 */
	//c.ShouldBind()

	uid := c.Request.Header.Get(config.SET_USER_ID_NAME)
	fmt.Println(uid)

	var req UserCollectionReqData
	if err := c.ShouldBindQuery(&req);err!=nil{
		r.GinParamBindErrorResult(c)
		return
	}

	var res UserCollectionResData
	if err := model.QueryUserCollectionBenzis(&res,uid,req.Page,req.Number);err != nil{
		c.JSON(201,r.SearchErrorResult())
		return
	}
	c.JSON(200,r.NewSuccessData("",res))
}





// ----------------------------/user/commit	 查询用户所有评论------------------------------------
type UserCommentReqData struct {
	SearchPageReqData
}
type UserCommentResData struct {
	Comments []model.Comment
}
func UserComments(c *gin.Context){
	/*
		用户id在token中或者在headers中
	*/
	uid := c.Request.Header.Get(config.SET_USER_ID_NAME)
	var req UserCommentReqData
	if err := c.ShouldBindJSON(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	var res UserCommentResData
	if err := model.QueryUserComments(res.Comments,uid,req.Page,req.Number);err!=nil{
		c.JSON(201,r.SearchErrorResult())
		return
	}
	c.JSON(200,r.NewSuccessData("",res))
}



// ----------------------/user/commit/reply 查询层级评论的回复评论------------------------
type UserCommitReplyReqData struct {
	SearchPageReqData
	Cid int64		`json:"c_id" form:"c_id" binding:"required"`

}
type UserCommentReplyResData struct {
	model.User
	model.ReplyComments
}
func UserCommentReply(c *gin.Context){

	/*
		用户id在token中或者在headers中
	*/
	var req UserCommitReplyReqData
	if err := c.ShouldBindJSON(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	var res []UserCommentReplyResData
	if err:= model.QueryUserReplyComments(req.Cid,req.Page,req.Number,res);err!=nil{
		c.JSON(201,r.SearchErrorResult())
		return
	}
	c.JSON(200,r.NewSuccessData("",res))
}






// ----------------------/user/info 	查询用户信息------------------------
type UserInfoReqData struct {
}
type UserInfoResData struct {
	model.User
}
func UserInfo(c *gin.Context){

	/*
		用户id在token中或者在headers中
	*/
	uid := c.Request.Header.Get(config.SET_USER_ID_NAME)

	var res UserInfoReqData
	if err:=model.QueryUserInfo(uid,res);err != nil{
		c.JSON(201,r.SearchErrorResult())
		return
	}
	c.JSON(200,r.NewSuccessData("",res))
}



// ----------------------/user/out	退出登录------------------------
type UserOutReqData struct {
}
type UserOutResData struct {
}
func UserOut(c *gin.Context){

	/*
		由于没有设置token唯一值(redis缓存...),所以直接返回由前端删除token或者后端更改token
	*/
	c.Request.Header.Set(config.SET_TOKEN_NAME,"")

	c.JSON(200,r.NewManualSuccessData(0,"user logout!",nil))
}

