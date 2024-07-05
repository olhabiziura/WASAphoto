package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))
	

	rt.router.POST("/photo", rt.AddPhotoHandler)
	rt.router.DELETE("/deletephoto/:pictureID", rt.DeletePhotoHandler)



	rt.router.POST("/addcomment/:pictureID", rt.AddComment)
	rt.router.DELETE("/deletecomment/:pictureID/:commentID", rt.DeleteComment)

	rt.router.POST("/addlike/:pictureID", rt.AddLike)
	rt.router.DELETE("/deletelike/:pictureID/", rt.DeleteLike)

	rt.router.POST("/follow/:userID", rt.AddFollowUser)
	rt.router.DELETE("/unfollow/:userID", rt.DeleteFollowUser)

	rt.router.POST("/ban/:userID", rt.AddBan)
	rt.router.DELETE("/ban/:userID", rt.DeleteBan)


	//rt.router.POST("/adduser", rt.AddUser)
    rt.router.POST("/session", rt.Login)
	rt.router.PUT("/username",rt.SetUsernameHandler)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
