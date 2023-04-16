package routes

import (
	"blog_server/controller"
	"blog_server/middleware"

	"github.com/gin-gonic/gin"
)

/* routes/rotues.go */
func CollectRoutes(r *gin.Engine) *gin.Engine {
	// 允许跨域访问
	r.Use(middleware.CORSMiddleware())

	// 注册
	r.POST("/register", controller.Register)
	// 登录
	r.POST("/login", controller.Login)

	// user路由组
	userGroup := r.Group("user")
	{
		userGroup.Use(middleware.AuthMiddleware())                       // 使用token验证拦截器
		userGroup.GET("v1/info", controller.GetInfo)                     // 获取用户信息
		userGroup.GET("v1/briefInfo/:id", controller.GetBriefInfo)       // 获取用户简要信息
		userGroup.GET("v1/detailedInfo/:id", controller.GetDetailedInfo) // 获取用户详细信息
		userGroup.PUT("v1/avatar/:id", controller.ModifyAvatar)          // 修改头像
		userGroup.PUT("v1/name/:id", controller.ModifyName)              // 修改用户名

	}

	// image路由组
	imageGroup := r.Group("image")
	{
		imageGroup.POST("/upload", controller.Upload)
	}

	// category分类路由组
	categoryGroup := r.Group("category")
	{
		categoryGroup.GET("v1", controller.SearchCategory)          // 查询分类
		categoryGroup.GET("/v1/:id", controller.SearchCategoryName) // 查询分类名

	}

	// 用户文章路由组
	articleGroup := r.Group("article")
	{
		articleGroup.POST("v1", middleware.AuthMiddleware(), controller.NewArticleController().Create)       // 发布文章
		articleGroup.PUT("v1/:id", middleware.AuthMiddleware(), controller.NewArticleController().Update)    // 修改文章
		articleGroup.DELETE("v1/:id", middleware.AuthMiddleware(), controller.NewArticleController().Delete) // 删除文章
		articleGroup.GET("v1/:id", controller.NewArticleController().Show)                                   // 查看文章
		articleGroup.POST("v1/list", controller.NewArticleController().List)                                 // 显示文章列表
	}

	// 我的收藏
	colRoutesGroup := r.Group("/collects")
	{
		colRoutesGroup.Use(middleware.AuthMiddleware())
		colRoutesGroup.GET("v1/:id", controller.Collects)        // 查询收藏
		colRoutesGroup.PUT("v1/new/:id", controller.NewCollect)  // 新增收藏
		colRoutesGroup.DELETE("v1/:index", controller.UnCollect) // 取消收藏
	}

	// 我的关注
	folRoutesGroup := r.Group("/following")
	{
		folRoutesGroup.Use(middleware.AuthMiddleware())
		folRoutesGroup.GET("v1/:id", controller.Following)      // 查询关注
		folRoutesGroup.PUT("v1/new/:id", controller.NewFollow)  // 新增关注
		folRoutesGroup.DELETE("v1/:index", controller.UnFollow) // 取消关注
	}

	return r
}
