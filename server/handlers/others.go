package handlers

import (
	"github.com/gin-gonic/gin"
	"pika/server/models"
	r "pika/server/code"
	re "pika/server/struct"
)

// ------------------/search   查询 -------------------------

func Search(c *gin.Context){

	var req re.SearchReqData
	if err := c.ShouldBindQuery(&req); err != nil{
		r.GinParamBindErrorResult(c)
		return
	}

	var res re.SearchResData
	err := models.SearchKeyBenzi(req.QueryName,req.Page,req.Number,res)
	if err != nil {
		r.SearchErrorResult()
		return
	}
	c.JSON(200,r.NewSuccessData("",res))

}


// ------------------/recommend 推荐本子   ------------------------

func Recommend(c *gin.Context){
	var req re.RecommendReqData
}


