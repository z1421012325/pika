package handlers

import (
	"github.com/gin-gonic/gin"
	"pika/config"
	r "pika/server/code"
	model "pika/server/models"
	"pika/tools"
	"strconv"
)

// ------------/benzi/create/storage  创建本子仓库   ----------------------
type CreateBenziStorageReqData struct {
	Title string		`json:"title" form:"title" binding:"required"`	// 标题
	BCover string		`json:"b_cover" form:"b_cover" binding:"required"`	// 封面
	Author string		`json:"author" form:"author" binding:"required"`	// 作者
	Type  []string		`json:"types" form:"types" binding:"required"`		// 分类
	Tags []string 		`json:"tags" form:"tags"`		// 标签
}
type CreateBenziStorageResData struct {

}
func CreateBenziStorage(c *gin.Context){

	uid := c.Request.Header.Get(config.SET_USER_ID_NAME)
	var req CreateBenziStorageReqData
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
type SignReqData struct {
	Files []string	`json:"files" form:"files" binding:"required"`	// 文件名
}
type SignResData struct {
}
func SignUploadUrl(c *gin.Context){
	/*
		aliyun oss or 自建ftp服务
	 */
	// todo oss 环境变量
	var req SignReqData
	if err := c.ShouldBind(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	res := tools.GetImageToken(req.Files)
	c.JSON(200,r.NewSuccessData("",res))
}












// ------------------/benzi/upload 上传本子   ------------------------
type uploadReqData struct {
	ImgUrls []string    `json:"img_url" form:"img_url" binding:"required"`	// 需要上传的url的字符串(经过签名)
	Bid 	int			`json:"bid" form:"bid" binding:"required"`		// 本子id
	Chapter string 		`json:"chapter" form:"chapter" binding:"default=第一章"`	// 章节名
}

type uploadResData struct {
}

func Upload(c *gin.Context){

	uid := c.Request.Header.Get(config.SET_USER_ID_NAME)
	var req uploadReqData
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
type DelReqData struct {
	Bid 	int			`json:"bid" form:"bid" binding:"required"`		// 本子id
}
type DelResData struct {
}

func DelBenzi(c *gin.Context){

	uid := c.Request.Header.Get(config.SET_USER_ID_NAME)
	var req DelReqData
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
type AddLikeReqData struct {
	Bid 	int			`json:"bid" form:"bid" binding:"required"`		// 本子id
}
type AddLikeResData struct {
}

func AddLike(c *gin.Context){

	uid := c.Request.Header.Get(config.SET_TOKEN_NAME)
	var req AddLikeReqData
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
type DelLikeReqData struct {
	Bid 	int			`json:"bid" form:"bid" binding:"required"`		// 本子id
}
type DelLikeResData struct {
}

func DelLike(c *gin.Context){

	uid := c.Request.Header.Get(config.SET_TOKEN_NAME)
	var req DelLikeReqData
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
type AddCollectionReqData struct {
	Bid 	int			`json:"bid" form:"bid" binding:"required"`		// 本子id
}
type AddCollectionResData struct {
}
func AddCollection(c *gin.Context){

	uid := c.Request.Header.Get(config.SET_TOKEN_NAME)
	var req AddCollectionReqData
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
type DelCollectionReqData struct {
	Bid 	int			`json:"bid" form:"bid" binding:"required"`		// 本子id
}
type DelCollectionResData struct {
}

func DelCollection(c *gin.Context){

	uid := c.Request.Header.Get(config.SET_TOKEN_NAME)
	var req DelCollectionReqData
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
type SearchBenziReqData struct {
	Bid 	int			`json:"bid" form:"bid" binding:"required"`		// 本子id
}
type SearchBenziResData struct {
	model.BenZi			// todo 测试gorm映射到该结构体上
	Chapter []model.BenZiImg		`gorm:"column:chapter"`
	Tags []model.Tags			`gorm:"column:tags"`
}

func SearchBenzi(c *gin.Context){
	var req SearchBenziReqData
	if err := c.ShouldBindQuery(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	var res SearchBenziResData
	err := model.QueryBenzi(req.Bid,&res)
	if err != nil {
		r.SearchErrorResult()
	}
	r.NewSuccessData("查询成功",res)
}













// ------------------/benzi/id   本子图片  ------------------------
type BenziImageReqData struct {
	Bid int		`json:"bid" form:"bid" binding:"required"`
	//Record int	`json:"record" form:"record"`
}
type BenziImageResData struct {
	model.BenZiImg
}

func BenziImage(c *gin.Context){

	uid := c.Request.Header.Get(config.SET_TOKEN_NAME)
	var req BenziImageReqData
	if err := c.ShouldBindQuery(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	var res []BenziImageResData
	err := model.RecordQueryAndUpRecord(uid,req.Bid,&res)
	if err != nil {
		r.SearchErrorResult()
	}
	r.NewSuccessData("查询成功",res)

}










// ------------------/benzi/insert/commit  	本子中添加评论   ------------------------
type InsertCommitReqData struct {
	Bid int		`json:"bid" form:"bid"`
	Comment  string    `json:"comment" form:"comment" binding:"required,max=150"`

	Reply string `json:"reply" form:"reply"`		// 默认一级评论,有值则为二级评论
	Cid int       `json:"cid" form:"cid"`
	ReplyUid int   `json:"uid" form:"uid"`
}
type InsertCommitResData struct {
}

func InsertComment(c *gin.Context){
	uid := c.Request.Header.Get(config.SET_TOKEN_NAME)
	var err error
	var req InsertCommitReqData
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
type QueryCommitReqData struct {
	Bid int		`json:"bid" form:"bid"`

	Reply string `json:"reply" form:"reply"`	// 是否为二级评论查询
	CId	int `json:"cid" form:"cid"`

	SearchPageReqData
}
type QueryCommitResData struct {
	model.Comment
	model.ReplyComments
	model.User
}

func QueryComment(c *gin.Context){

	var req QueryCommitReqData
	if err := c.ShouldBindQuery(&req);err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	var res []QueryCommitResData
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

