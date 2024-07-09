package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))
	rt.router.GET("/search", rt.SearchUser)

	rt.router.POST("/photo", rt.AddPhotoHandler)
	rt.router.GET("/photo", rt.GetUserFeed)
	rt.router.DELETE("/deletephoto/:pictureID", rt.DeletePhotoHandler)

	rt.router.POST("/addcomment/:pictureID", rt.AddComment)
	rt.router.DELETE("/deletecomment/:pictureID/comment/:commentID", rt.DeleteComment)

	rt.router.POST("/like/:pictureID", rt.AddLike)
	rt.router.DELETE("/like/:pictureID", rt.DeleteLike)

	rt.router.POST("/follow/:userID", rt.AddFollowUser)
	rt.router.DELETE("/unfollow/:userID", rt.DeleteFollowUser)

	rt.router.POST("/ban/:userID", rt.AddBan)
	rt.router.DELETE("/ban/:userID", rt.DeleteBan)
	rt.router.GET("/ban/:userID", rt.GetIfBanned)

	//rt.router.POST("/getid", rt.GetUserId)
	rt.router.POST("/session", rt.Login)
	rt.router.PUT("/username", rt.SetUsernameHandler)
	rt.router.GET("/profile/:userID", rt.GetProfile)
	rt.router.GET("/stream", rt.GetUserStream)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
