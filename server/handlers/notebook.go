package handlers

import (
	"github.com/gin-gonic/gin"
	"pika/config"
	r "pika/server/code"
	model "pika/server/models"
	"pika/tools"
	"strconv"
	re "pika/server/struct"

)

// ------------/benzi/create/storage  创建本子仓库   ----------------------

func CreateBenziStorage(c *gin.Context){

	uid := c.Request.Header.Get(config.SET_USER_ID_NAME)
	var req re.CreateBenziStorageReqData
	if err := c.ShouldBind(&req);err!=nil{
		r.GinParamBindErrorResult(c)
		return
	}

	err := model.CreateBenziStorge(uid,req.Title,req.BCover,req.Author,req.Type,req.Tags)
	if err != nil {
		r.AddErrorResult("本子仓库创建失败",err.Error())
	}

	r.NewSuccessData("本子仓库创建成功",nil)
}









// ------------------/benzi/sign 签名url   ------------------------

func SignUploadUrl(c *gin.Context){
	/*
		aliyun oss or 自建ftp服务
	 */
	// todo oss 环境变量
	var req re.SignReqData
	if err := c.ShouldBind(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	res := tools.GetImageToken(req.Files)
	c.JSON(200,r.NewSuccessData("",res))
}












// ------------------/benzi/upload 上传本子   ------------------------


func Upload(c *gin.Context){

	uid := c.Request.Header.Get(config.SET_USER_ID_NAME)
	var req re.UploadReqData
	if err := c.ShouldBind(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	uidInt,_ := strconv.Atoi(uid)
	if err := model.UploadBnezi(req.Bid,uidInt,req.ImgUrls,req.Chapter); err != nil{
		c.JSON(201,r.NewManualFailedData(r.FAILED_CODE,"创建失败",err.Error()))
		return
	}

	c.JSON(200,r.NewSuccessData("上传成功",nil))
}











// ------------------/benzi/del 删除本子   ------------------------


func DelBenzi(c *gin.Context){

	uid := c.Request.Header.Get(config.SET_USER_ID_NAME)
	var req re.DelReqData
	if err := c.ShouldBind(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	err := model.DelBenzi(uid,req.Bid)
	if err != nil {
		r.DelErrorResult("删除失败",err.Error())
	}
	r.NewSuccessData("删除成功",nil)
}







// ------------------/benzi/add/like 点赞本子   ------------------------

func AddLike(c *gin.Context){

	uid := c.Request.Header.Get(config.SET_TOKEN_NAME)
	var req re.AddLikeReqData
	if err := c.ShouldBind(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	err := model.UserLikeBenzi(req.Bid,uid)
	if err != nil {
		r.OperationErrorResult("点赞失败",err.Error())
	}
	r.NewSuccessData("点赞成功",nil)

}



// ------------------/benzi/del/like  删除点赞  ------------------------


func DelLike(c *gin.Context){

	uid := c.Request.Header.Get(config.SET_TOKEN_NAME)
	var req re.DelLikeReqData
	if err := c.ShouldBind(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	err := model.CancelUserLikeBenzi(req.Bid,uid)
	if err != nil {
		r.OperationErrorResult("取消点赞失败",err.Error())
	}
	r.NewSuccessData("取消点赞成功",nil)

}


// ------------------/benzi/add/collection  添加收藏 ------------------------

func AddCollection(c *gin.Context){

	uid := c.Request.Header.Get(config.SET_TOKEN_NAME)
	var req re.AddCollectionReqData
	if err := c.ShouldBind(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	err := model.UserAddBenziCollection(req.Bid,uid)
	if err != nil {
		r.OperationErrorResult("收藏失败",err.Error())
	}
	r.NewSuccessData("收藏成功",nil)

}


// ------------------/benzi/del/collection  删除收藏  ------------------------

func DelCollection(c *gin.Context){

	uid := c.Request.Header.Get(config.SET_TOKEN_NAME)
	var req re.DelCollectionReqData
	if err := c.ShouldBind(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	err := model.UnUserAddBenziCollection(req.Bid,uid)
	if err != nil {
		r.OperationErrorResult("删除收藏失败",err.Error())
	}
	r.NewSuccessData("删除收藏成功",nil)

}


// ------------------/benzi/  查询本子id query参数 b_id   ------------------------


func SearchBenzi(c *gin.Context){
	var req re.SearchBenziReqData
	if err := c.ShouldBindQuery(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	var res re.SearchBenziResData
	err := model.QueryBenzi(req.Bid,&res)
	if err != nil {
		r.SearchErrorResult()
	}
	r.NewSuccessData("查询成功",res)
}













// ------------------/benzi/id   本子图片  ------------------------

func BenziImage(c *gin.Context){

	uid := c.Request.Header.Get(config.SET_TOKEN_NAME)
	var req re.BenziImageReqData
	if err := c.ShouldBindQuery(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	var res []re.BenziImageResData
	err := model.RecordQueryAndUpRecord(uid,req.Bid,&res)
	if err != nil {
		r.SearchErrorResult()
	}
	r.NewSuccessData("查询成功",res)

}










// ------------------/benzi/insert/commit  	本子中添加评论   ------------------------


func InsertComment(c *gin.Context){
	uid := c.Request.Header.Get(config.SET_TOKEN_NAME)
	var err error
	var req re.InsertCommitReqData
	if err := c.ShouldBindQuery(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	if req.Reply == "" {		// 一级评论
		err = model.AddBenziComment(uid,req.Bid,req.Comment)
	}else {						// 二级评论
		err = model.AddReplyComment(uid,req.Cid,req.ReplyUid,req.Comment)
	}
	if err != nil {
		r.AddErrorResult("添加评论失败",err.Error())
	}
	r.NewSuccessData("添加评论成功",nil)

}


// ------------------/benzi/query/commit		查询本子评论   ------------------------


func QueryComment(c *gin.Context){

	var req re.QueryCommitReqData
	if err := c.ShouldBindQuery(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	var res []re.QueryCommitResData
	var err error
	if req.Reply == "" {
		err = model.QueryBenziComment(req.Bid,req.Page,req.Number,&res)
	}else {
		err = model.QueryBenziReplyComment(req.CId,req.Page,req.Number,&res)
	}
	if err != nil {
		r.SearchErrorResult()
	}
	r.NewSuccessData("查询成功",res)
}

