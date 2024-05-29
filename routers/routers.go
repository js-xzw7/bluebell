package routers

import (
	"bluebell/controller"
	"bluebell/middleware"
	"bluebell/middleware/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// r := gin.Default()
	r := gin.New()
	r.Use(middleware.Zap())
	v1 := r.Group("/api/v1")
	v1.POST("/login", controller.LoginHandler)
	v1.POST("/signup", controller.SignUpHandler)
	// v1.GET("/refreshToken", controller.RefreshTokenHandler)

	v1.POST("/addCommunity", controller.AddCommunityHandler)
	v1.GET("/getCommunityList", controller.GetCommunityListHandler)
	v1.GET("/getCommunityById", controller.GetCommunityByIdHandler)

	v1.POST("/addPost", controller.AddPostHandler)
	v1.GET("/getPostById", controller.GetPostByIdHandler)
	v1.POST("/getPostListByIds", controller.GetPostListByIdsHandler)
	v1.GET("/getPostList", controller.GetPostListHandler)

	v1.POST("/addComment", controller.AddCommentHandler)
	v1.POST("/getCommentListByIds", controller.GetCommentListByIdsHandler)
	v1.POST("/getCommentList", controller.GetCommentListHandler)

	v1.Use(auth.JWTAuthMiddleware())
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
