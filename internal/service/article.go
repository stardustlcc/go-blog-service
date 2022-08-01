package service

type ArticleGetRequest struct {
	ID    uint8 `form:"id" binding:"required,gte=1"`
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	Title string `form:"title" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleCreateRequest struct {
	Title     string `form:"title" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleUpdateRequest struct {
	ID         uint8  `form:"id" binding:"required,gte=1"`
	Title      string `form:"title" binding:"min=3, max=100"`
	State      uint8  `form:"state" binding:"required, oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type ArticleDeleteRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
