package servers

import (

	// init db
	"github.com/gin-gonic/gin"
	_ "pika/server/db"

)



func Run(address string){

	server := gin.Default()

	userGroup := server.Group("/user")
	{
		// /user/login		登录
		userGroup.POST("/login",nil)
		// /user/registry	注册
		userGroup.POST("/registry",nil)
	}

	// 下面的请求都需要通过验证是否登录用户中间件
	server.Use(nil)
	{
		// /search   查询
		server.GET("/search",nil)
		// 推荐本子
		server.GET("/recommend",nil)


		// /user/collection 查询收藏本子
		userGroup.GET("/collection",nil)
		// /user/commit		查询评论
		userGroup.GET("/commit",nil)
		// /user/info 		查询用户信息
		userGroup.GET("/info",nil)

		// /user/out		退出登录
		userGroup.POST("/out",nil)
	}

	// 本子
	noteBookGroup := server.Group("/benzi")
	{
		// 签名url
		noteBookGroup.POST("/sign",nil)
		// 上传本子
		noteBookGroup.POST("/upload",nil)
		// 删除本子
		noteBookGroup.DELETE("/del",nil)

		// 点赞本子
		noteBookGroup.POST("/add/like",nil)
		// 删除点赞
		noteBookGroup.DELETE("/del/like",nil)

		// 添加收藏
		noteBookGroup.POST("/add/collection",nil)
		// 删除收藏
		noteBookGroup.DELETE("/del/collection",nil)


		// 查询本子id
		noteBookGroup.GET("/id",nil)
		// 本子图片
		noteBookGroup.GET("/",nil)
		// /benzi/insert/commit  	本子中添加评论
		noteBookGroup.POST("/insert/commit",nil)
		// /benzi/query/commit		查询本子评论
		noteBookGroup.GET("/query/commit",nil)

	}

	// tag标签
	tagGroup := server.Group("/tag")
	{
		// 添加tag
		tagGroup.POST("/add",nil)
		// 删除tag
		tagGroup.DELETE("/del",nil)
		// 查看个人添加的tag
		tagGroup.GET("/user/add",nil)
		// 查看所有tag
		tagGroup.GET("/all/tag",nil)
	}


	adminGroup := server.Group("/admin")
	{
		// 注册用户
		// 删除用户
		// 更改权限
		// 获取用户,上传人,管理员列表
		// 获取用户,上传人,管理员信息
		adminGroup.GET("/")
	}






	err := server.Run(address)
	if err != nil {
		panic(err)
	}
}
