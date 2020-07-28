package handlers

import "github.com/gin-gonic/gin"




// ------------------通过用户注册  /admin/verify/unregistry ------------------------
type AdminVerifyUnregistryReqData struct {
}
type AdminVerifyUnregistryResData struct {
}

func AdminVerifyUnregistry(c *gin.Context){
}



// ------------------查看等待通过注册用户	/admin/query/unregistry ------------------------
type AdminQueryUnregistryReqData struct {
}
type AdminQueryUnregistryResData struct {
}

func AdminQueryUnregistry(c *gin.Context){
}



// ------------------冻结用户	/admin/frozen/user ------------------------
type AdminFrozenUserReqData struct {
}
type AdminFrozenUserResData struct {
}

func AdminFrozenUser(c *gin.Context){
}



// ------------------更改权限	/admin/change/grade ------------------------
type AdminChangeGradeReqData struct {
}
type AdminChangeGradeResData struct {
}

func AdminChangeGrade(c *gin.Context){
}



// ------------------获取用户,上传人,管理员列表	/admin/query/user/info ------------------------
type AdminQueryUserInfoReqData struct {
}
type AdminQueryUserInfoResData struct {
}

func AdminQueryUserInfo(c *gin.Context){
}



// ------------------获取用户,上传人,管理员信息	/admin/query/users ------------------------
type AdminQueryUsersReqData struct {
}
type AdminQueryUsersResData struct {
}

func AdminQueryUsers(c *gin.Context){
}


