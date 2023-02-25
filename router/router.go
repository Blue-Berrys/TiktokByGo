package router

import (
	"github.com/gin-gonic/gin"
	"github.com/neverTanking/TiktokByGo/controller"
	"github.com/neverTanking/TiktokByGo/controller/comment"
	"github.com/neverTanking/TiktokByGo/controller/video"
	"github.com/neverTanking/TiktokByGo/db"
	"github.com/neverTanking/TiktokByGo/middleware/JWT"
	"github.com/neverTanking/TiktokByGo/middleware/NoAuth"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	// public directory is used to serve static resources
	db.Init()
	r.Static("/static", "./static")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	//apiRouter.GET("/user/", controller.UserInfo)
	//apiRouter.POST("/user/register/", controller.Register)
	//apiRouter.POST("/user/login/", controller.Login)
	//apiRouter.POST("/publish/action/", controller.Publish)
	//apiRouter.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", JWT.JWTMiddleware(), video.LikeActionController)
	apiRouter.GET("/favorite/list/", NoAuth.NoAuthMiddleWare(), video.FavoriteListController)
	apiRouter.POST("/comment/action/", JWT.JWTMiddleware(), comment.CommentActionController)
	apiRouter.GET("/comment/list/", JWT.JWTMiddleware(), comment.QueryCommentListController)
	//
	//// extra apis - II
	//apiRouter.POST("/relation/action/", controller.RelationAction)
	//apiRouter.GET("/relation/follow/list/", controller.FollowList)
	//apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	//apiRouter.GET("/relation/friend/list/", controller.FriendList)
	//apiRouter.GET("/message/chat/", controller.MessageChat)
	//apiRouter.POST("/message/action/", controller.MessageAction)
	return r
}
