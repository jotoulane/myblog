package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web01/controllers"
	"web01/middleware"
)

func SetUp() *gin.Engine {

	r := gin.New()

	//加载静态文件
	r.Static("xxx", "static")
	//模板解析
	r.LoadHTMLGlob("static/pages/*")
	//注册页面
	r.GET("/register", func(context *gin.Context) {
		context.HTML(http.StatusOK, "register.tmpl", gin.H{}) //模板渲染
	})
	//登录页面
	r.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.tmpl", gin.H{}) //模板渲染
	})
	//主页面
	//r.GET("/", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "home.tmpl", gin.H{}) //模板渲染
	//})
	r.GET("/", controllers.HomePage)
	//新建文章
	r.GET("/new", func(context *gin.Context) {
		context.HTML(http.StatusOK, "add.tmpl", gin.H{}) //模板渲染
	})
	//管理页面
	r.GET("/manage", func(context *gin.Context) {
		context.HTML(http.StatusOK, "manage.tmpl", gin.H{}) //模板渲染
	})

	//错误提示页面
	r.GET("/alarm", func(context *gin.Context) {
		context.HTML(http.StatusOK, "alarm.tmpl", gin.H{}) //模板渲染
	})
	//照片墙
	r.GET("/photos", func(context *gin.Context) {
		context.HTML(http.StatusOK, "photos.tmpl", gin.H{})
	})

	//注册
	r.POST("/register", controllers.RegisterHandler)
	//登录
	r.POST("/login", controllers.LoginHandler)

	group := r.Group("api/v1")
	group.Use(middleware.JWTAuthMiddleware())

	//新建文章
	r.POST("/new", controllers.NewArticle)
	//文章列表
	r.GET("/lists", controllers.ListArticle)
	//根据标签分类查询文章
	r.GET("/lists/:tags", controllers.ListTagsArticle)
	//文章详情
	r.GET("/detail/:id", controllers.DetailArticle)
	//文章管理列表
	r.GET("/manage/lists", controllers.ManageListArticle)
	////文章管理列表,包括已经删除的文章
	//r.GET("/recycle/lists", controllers.RecycleListArticle)
	//文章管理详情
	r.GET("/manage/detail/:id", controllers.ManageDetailArticle)
	//编辑文章页面
	r.GET("/update/:id", controllers.ManageUpdate)
	//编辑文章提交
	r.POST("/update/:id", controllers.ManageUpdateSubmit)
	//删除文章
	r.GET("/delete/:id", controllers.ManageDeleteSubmit)
	//点赞文章
	r.GET("/like/:id", controllers.LikeArticle)

	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
