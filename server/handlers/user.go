package handlers

import (
	"fmt"
	r "pika/server/code"
	model "pika/server/models"
	"pika/tools"
	"pika/config"
	re "pika/server/struct"

	"github.com/gin-gonic/gin"
)




// -------------------------------/user/login --------------------------------------------


// 登录
func UserLogin(c *gin.Context){

	var req re.UserLoginReqData
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


func UserRegistry(c *gin.Context){
	var req re.UserRegistryReqData
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

	var req re.UserCollectionReqData
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

func UserComments(c *gin.Context){
	/*
		用户id在token中或者在headers中
	*/
	uid := c.Request.Header.Get(config.SET_USER_ID_NAME)
	var req re.UserCommentReqData
	if err := c.ShouldBindJSON(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	var res re.UserCommentResData
	if err := model.QueryUserComments(res.Comments,uid,req.Page,req.Number);err!=nil{
		c.JSON(201,r.SearchErrorResult())
		return
	}
	c.JSON(200,r.NewSuccessData("",res))
}



// ----------------------/user/commit/reply 查询层级评论的回复评论------------------------

func UserCommentReply(c *gin.Context){

	/*
		用户id在token中或者在headers中
	*/
	var req re.UserCommitReplyReqData
	if err := c.ShouldBindJSON(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	var res []re.UserCommentReplyResData
	if err:= model.QueryUserReplyComments(req.Cid,req.Page,req.Number,res);err!=nil{
		c.JSON(201,r.SearchErrorResult())
		return
	}
	c.JSON(200,r.NewSuccessData("",res))
}






// ----------------------/user/info 	查询用户信息------------------------

func UserInfo(c *gin.Context){

	/*
		用户id在token中或者在headers中
	*/
	uid := c.Request.Header.Get(config.SET_USER_ID_NAME)

	var res re.UserInfoReqData
	if err:=model.QueryUserInfo(uid,res);err != nil{
		c.JSON(201,r.SearchErrorResult())
		return
	}
	c.JSON(200,r.NewSuccessData("",res))
}



// ----------------------/user/out	退出登录------------------------

func UserOut(c *gin.Context){

	/*
		由于没有设置token唯一值(redis缓存...),所以直接返回由前端删除token或者后端更改token
	*/
	c.Request.Header.Set(config.SET_TOKEN_NAME,"")

	c.JSON(200,r.NewManualSuccessData(0,"user logout!",nil))
}

