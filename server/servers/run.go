package servers

import (
	// init
	_ "pika/config"
	_ "pika/server/db"

	handler "pika/server/handlers"
	"pika/server/middleware"

	"github.com/gin-gonic/gin"
)

func Run(address string) {

	server := gin.Default()
	userGroup := server.Group("/user")

	// 不需要验证登录
	{
		// /user/login		登录
		userGroup.POST("/login", handler.UserLogin)
		// /user/registry	注册
		userGroup.POST("/registry", handler.UserRegistry)
	}

	// 下面的请求都需要通过验证是否登录用户中间件
	server.Use(middleware.VerifyUserLogin())
	{
		// /search   查询
		server.GET("/search", handler.Search)
		// recommend 推荐本子
		server.GET("/recommend", handler.Recommend)

		// /user/collection 查询收藏本子
		userGroup.GET("/collection", handler.UserCollection)
		// /user/commits		查询用户评论
		userGroup.GET("/commits", handler.UserCommits)
		// /user/commit/reply		查询层级评论的回复评论
		userGroup.GET("/commit/reply", handler.UserCommitReply)
		// /user/info 		查询用户信息
		userGroup.GET("/info", handler.UserInfo)

		// /user/out		退出登录
		userGroup.POST("/out", handler.UserOut)
	}

	// 本子
	noteBookGroup := server.Group("/benzi")
	{
		// /benzi/sign 签名url
		noteBookGroup.POST("/sign", handler.SignUploadUrl)
		//  /benzi/upload 上传本子
		noteBookGroup.POST("/upload", handler.Upload)
		// /benzi/del 删除本子
		noteBookGroup.DELETE("/del", handler.DelBenzi)

		// /benzi/add/like 点赞本子
		noteBookGroup.POST("/add/like", handler.AddLike)
		// /benzi/del/like删除点赞
		noteBookGroup.DELETE("/del/like", handler.DelLike)

		// /benzi/add/collection  添加收藏
		noteBookGroup.POST("/add/collection", handler.AddCollection)
		// /benzi/del/collection  删除收藏
		noteBookGroup.DELETE("/del/collection", handler.DelCollection)

		// 查询本子id  	返回数据含有观看记录,是否点赞,收藏,点赞数,评论数
		// /benzi/b_id=4513213
		noteBookGroup.GET("/", handler.SearchBenzi) // query参数 b_id
		// /benzi/id   本子图片
		noteBookGroup.GET("/id", handler.BenziImage) // query参数 b_id
		// /benzi/insert/commit  	本子中添加评论
		noteBookGroup.POST("/insert/commit", handler.InsertCommit)
		// /benzi/query/commit		查询本子评论
		noteBookGroup.GET("/query/commit", handler.QueryCommit)

	}

	// tag标签
	tagGroup := server.Group("/tag")
	{
		// /tag/add  添加tag
		tagGroup.POST("/add", handler.TagAdd)
		// /tag/del  删除tag
		tagGroup.DELETE("/del", handler.TagDel)
		// /tag/user/up 查看上传人添加的tag
		tagGroup.GET("/user/up", handler.TagUserUp)
		// /tag/all 查看所有tag
		tagGroup.GET("/all", handler.TagAll)
	}

	adminGroup := server.Group("/admin")
	{

		// 通过用户注册  /admin/verify/unregistry
		adminGroup.POST("/verify/unregistry", handler.AdminVerifyUnregistry)
		// 查看等待通过注册用户	/admin/query/unregistry
		adminGroup.GET("/query/unregistry", handler.AdminQueryUnregistry)
		// 冻结用户	/admin/frozen/user
		adminGroup.DELETE("/frozen/user", handler.AdminFrozenUser)
		// 更改权限	/admin/change/grade
		adminGroup.POST("/change/grade", handler.AdminChangeGrade)
		// 获取用户,上传人,管理员列表	/admin/query/user/info
		adminGroup.GET("/query/user/info", handler.AdminQueryUserInfo)
		// 获取用户,上传人,管理员信息	/admin/query/users
		adminGroup.GET("/query/users", handler.AdminQueryUsers)
	}

	err := server.Run(address)
	if err != nil {
		panic(err)
	}
}
