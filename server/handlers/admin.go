package handlers

import (
	req "pika/server/struct"

	"github.com/gin-gonic/gin"
)




// ------------------通过用户注册  /admin/verify/unregistry ------------------------
func AdminVerifyUnregistry(c *gin.Context){
	var req req.AdminVerifyUnregistryReqData
}



// ------------------查看等待通过注册用户	/admin/query/unregistry ------------------------


func AdminQueryUnregistry(c *gin.Context){
	var req req.AdminQueryUnregistryReqData
}



// ------------------冻结用户	/admin/frozen/user ------------------------

func AdminFrozenUser(c *gin.Context){
	var req req.AdminFrozenUserReqData
}



// ------------------更改权限	/admin/change/grade ------------------------

func AdminChangeGrade(c *gin.Context){
	var req req.AdminChangeGradeReqData
}



// ------------------获取用户,上传人,管理员列表	/admin/query/user/info ------------------------

func AdminQueryUserInfo(c *gin.Context){
	var req req.AdminQueryUserInfoReqData
}



// ------------------获取用户,上传人,管理员信息	/admin/query/users ------------------------
type AdminQueryUsersReqData struct {
}
type AdminQueryUsersResData struct {
}

func AdminQueryUsers(c *gin.Context){
	var req req.AdminQueryUsersReqData
}





// ------------------ 添加分类 /admin/add/classify------------------------
func AddClassify(c *gin.Context){
	var req req.AddClassifyReqData

}


// ------------------ 删除分类 /admin/del/classify------------------------

func DelClassify(c *gin.Context){
	var req req.DelClassifyReqData

}


// -----------------查看分类,含增加人 /admin/all/classify-------------------------

func AllClassify(c *gin.Context){
	var req req.AllClassifyReqData

}
