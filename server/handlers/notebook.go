package handlers

import "github.com/gin-gonic/gin"

// ------------------/benzi/sign 签名url   ------------------------
type SignReqData struct {
}
type SignResData struct {
}
func SignUploadUrl(c *gin.Context){
}


// ------------------/benzi/upload 上传本子   ------------------------
type uploadReqData struct {
}
type uploadResData struct {
}

func Upload(c *gin.Context){
}

// ------------------/benzi/del 删除本子   ------------------------
type DelReqData struct {
}
type DelResData struct {
}

func DelBenzi(c *gin.Context){
}

// ------------------/benzi/add/like 点赞本子   ------------------------
type AddLikeReqData struct {
}
type AddLikeResData struct {
}

func AddLike(c *gin.Context){
}



// ------------------/benzi/del/like  删除点赞  ------------------------
type DelLikeReqData struct {
}
type DelLikeResData struct {
}

func DelLike(c *gin.Context){
}


// ------------------/benzi/add/collection  添加收藏 ------------------------
type AddCollectionReqData struct {
}
type AddCollectionResData struct {
}

func AddCollection(c *gin.Context){
}


// ------------------/benzi/del/collection  删除收藏  ------------------------
type DelCollectionReqData struct {
}
type DelCollectionResData struct {
}

func DelCollection(c *gin.Context){
}


// ------------------/benzi/  查询本子id query参数 b_id   ------------------------
type SearchBenziReqData struct {
}
type SearchBenziResData struct {
}

func SearchBenzi(c *gin.Context){
}
// ------------------/benzi/id   本子图片  ------------------------
type BenziImageReqData struct {
}
type BenziImageResData struct {
}

func BenziImage(c *gin.Context){
}

// ------------------/benzi/insert/commit  	本子中添加评论   ------------------------
type InsertCommitReqData struct {
}
type InsertCommitResData struct {
}

func InsertCommit(c *gin.Context){
}


// ------------------/benzi/query/commit		查询本子评论   ------------------------
type QueryCommitReqData struct {
}
type QueryCommitResData struct {
}

func QueryCommit(c *gin.Context){
}

