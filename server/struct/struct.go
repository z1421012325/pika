package _struct

import "pika/server/models"

type AdminVerifyUnregistryReqData struct {
}
type AdminVerifyUnregistryResData struct {
}


type AdminQueryUnregistryReqData struct {
}
type AdminQueryUnregistryResData struct {
}

type AdminFrozenUserReqData struct {
}
type AdminFrozenUserResData struct {
}

type AdminChangeGradeReqData struct {
}
type AdminChangeGradeResData struct {
}

type AdminQueryUserInfoReqData struct {
}
type AdminQueryUserInfoResData struct {
}


type AdminQueryUsersReqData struct {
}
type AdminQueryUsersResData struct {
}


type AddClassifyReqData struct {
}
type AddClassifyResData struct {
}

type DelClassifyReqData struct {
}
type DelClassifyResData struct {
}


type AllClassifyReqData struct {
}
type AllClassifyResData struct {
}




type SearchReqData struct {
	QueryName string `json:"q" form:"q"`
	SearchPageReqData
}
type SearchData struct {
	models.BenZi
	BenziImgTotal int	`gorm:"column:total"`  // 用于查询页数 查询之后添加到Benzi.title后
	CollectionTotal int `json:"collection_total" gorm:"collection_total"`
}
type SearchResData struct {
	Total int `gorm:"column:total" json:"total"`
	CurrentPage int64 `json:"cur_page"`
	NextPage bool	`json:"next_page"`
	Data []SearchData		`json:"data"`
}



type RecommendReqData struct {
}
type RecommendResData struct {
}





type SearchPageReqData struct {
	Page int64			`json:"page" form:"page" binding:"default=1"`
	Number int64			`json:"number" form:"number" binding:"default=20"`
}

type UserLoginReqData struct {
	//Username string		`json:"username" form:"username" binding:"required,min=2,max=20"`
	//PassWord string		`json:"password" form:"password" binding:"required,min=5,max=30"`
	//// 登录等级
	//Grade int			`json:"grade" form:"grade" binding:"required"`

	// 参数字段 username,paswd,grade
	models.User
}


type UserRegistryReqData struct {
	//Username string	`json:"username" form:"username" binding:"required,min=2,max=20"`
	//PassWord string	`json:"password" form:"password" binding:"required,min=5,max=30"`
	Nickname string	`json:"nickname" form:"nickname" binding:"required,min=3,max=30"`
	//Grade 	int 	`json:"grade" form:"grade"`

	models.User
}


type UserCollectionReqData struct {
	// 拦截器拦截token解析并在header中设置了UID 试试结构体中用gin的 ShouldBind 映射
	//Uid int 	`json:"UID" form:"UID" binding:"required"`
	SearchPageReqData
}



type UserCommentReqData struct {
	SearchPageReqData
}
type UserCommentResData struct {
	Comments []models.Comment
}


type UserCommitReplyReqData struct {
	SearchPageReqData
	Cid int64		`json:"c_id" form:"c_id" binding:"required"`

}
type UserCommentReplyResData struct {
	models.User
	models.ReplyComments
}



type UserInfoReqData struct {
}
type UserInfoResData struct {
	models.User
}


type CreateBenziStorageReqData struct {
	Title string		`json:"title" form:"title" binding:"required"`	// 标题
	BCover string		`json:"b_cover" form:"b_cover" binding:"required"`	// 封面
	Author string		`json:"author" form:"author" binding:"required"`	// 作者
	Type  []string		`json:"types" form:"types" binding:"required"`		// 分类
	Tags []string 		`json:"tags" form:"tags"`		// 标签
}
type CreateBenziStorageResData struct {
}


type SignReqData struct {
	Files []string	`json:"files" form:"files" binding:"required"`	// 文件名
}
type SignResData struct {
}



type UploadReqData struct {
	ImgUrls []string    `json:"img_url" form:"img_url" binding:"required"`	// 需要上传的url的字符串(经过签名)
	Bid 	int			`json:"bid" form:"bid" binding:"required"`		// 本子id
	Chapter string 		`json:"chapter" form:"chapter" binding:"default=第一章"`	// 章节名
}

type uploadResData struct {
}


type DelReqData struct {
	Bid 	int			`json:"bid" form:"bid" binding:"required"`		// 本子id
}
type DelResData struct {
}

type AddLikeReqData struct {
	Bid 	int			`json:"bid" form:"bid" binding:"required"`		// 本子id
}
type AddLikeResData struct {
}



type DelLikeReqData struct {
	Bid 	int			`json:"bid" form:"bid" binding:"required"`		// 本子id
}
type DelLikeResData struct {
}

type AddCollectionReqData struct {
	Bid 	int			`json:"bid" form:"bid" binding:"required"`		// 本子id
}
type AddCollectionResData struct {
}



type DelCollectionReqData struct {
	Bid 	int			`json:"bid" form:"bid" binding:"required"`		// 本子id
}
type DelCollectionResData struct {
}


type SearchBenziReqData struct {
	Bid 	int			`json:"bid" form:"bid" binding:"required"`		// 本子id
}
type SearchBenziResData struct {
	models.BenZi			// todo 测试gorm映射到该结构体上
	Chapter []models.BenZiImg		`gorm:"column:chapter"`
	Tags []models.Tags			`gorm:"column:tags"`
}



type BenziImageReqData struct {
	Bid int		`json:"bid" form:"bid" binding:"required"`
	//Record int	`json:"record" form:"record"`
}
type BenziImageResData struct {
	models.BenZiImg
}


type InsertCommitReqData struct {
	Bid int		`json:"bid" form:"bid"`
	Comment  string    `json:"comment" form:"comment" binding:"required,max=150"`

	Reply string `json:"reply" form:"reply"`		// 默认一级评论,有值则为二级评论
	Cid int       `json:"cid" form:"cid"`
	ReplyUid int   `json:"uid" form:"uid"`
}
type InsertCommitResData struct {
}



type QueryCommitReqData struct {
	Bid int		`json:"bid" form:"bid"`

	Reply string `json:"reply" form:"reply"`	// 是否为二级评论查询
	CId	int `json:"cid" form:"cid"`

	SearchPageReqData
}
type QueryCommitResData struct {
	models.Comment
	models.ReplyComments
	models.User
}