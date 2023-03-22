package model

type ParamRegister struct {
	Username        string `json:"username" binding:"required" form:"username"`
	Password        string `json:"password" binding:"required" form:"password"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password" form:"confirm_password"`
}
type ParamLogin struct {
	Username string `json:"username" binding:"required" form:"username"`
	Password string `json:"password" binding:"required" form:"password"`
}
type ParamNewArticle struct {
	Title   string `json:"title" binding:"required" form:"title"`
	Content string `json:"content" binding:"required" form:"content"`
	Short   string `json:"short" binding:"required" form:"short"`
	Tags    string `json:"tags" binding:"required" form:"tags"`
}
