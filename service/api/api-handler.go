package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	var r = rt.router

	// Generic routes
	r.GET("/", rt.getAppInfo)
	r.GET("/liveness", rt.liveness)

	// Session routes
	r.POST("/doLogin", rt.wrap(rt.doLogin))

	// User routes
	r.GET("/users", rt.wrap(rt.findUser))
	r.GET("/users/:user_uuid", rt.wrap(rt.getUserProfile))
	r.PATCH("/me/username", rt.wrap(rt.setMyUserName))

	// Post routes
	r.GET("/me/feed", rt.wrap(rt.getMyStream))
	r.PUT("/me/feed", rt.wrap(rt.uploadPhoto))
	r.DELETE("/me/feed/:post_uuid", rt.wrap(rt.deletePhoto))
	r.GET("/users/:user_uuid/feed", rt.wrap(rt.getUserFeed))
	r.GET("/users/:user_uuid/feed/:post_uuid", rt.wrap(rt.getPhoto))
	r.GET("/users/:user_uuid/feed/:post_uuid/webp", rt.wrap(rt.getPostImage)) // TODO: api.yaml

	// Comment routes
	r.GET("/users/:user_uuid/feed/:post_uuid/comments", rt.wrap(rt.getComments))
	r.PUT("/users/:user_uuid/feed/:post_uuid/comments", rt.wrap(rt.commentPhoto))
	r.DELETE("/users/:user_uuid/feed/:post_uuid/comments/:comment_uuid", rt.wrap(rt.uncommentPhoto))

	// Likes routes
	r.GET("/users/:user_uuid/feed/:post_uuid/likes", rt.wrap(rt.getLikes))
	r.PUT("/users/:user_uuid/feed/:post_uuid/likes", rt.wrap(rt.likePhoto))
	r.DELETE("/users/:user_uuid/feed/:post_uuid/likes", rt.wrap(rt.unlikePhoto))

	// Follow routes
	r.GET("/me/followed_users", rt.wrap(rt.getFollowedUsers))
	r.PUT("/me/followed_users/:user_uuid", rt.wrap(rt.followUser))
	r.DELETE("/me/followed_users/:user_uuid", rt.wrap(rt.unfollowUser))

	// Ban routes
	r.GET("/me/banned_users", rt.wrap(rt.getBannedUsers))
	r.PUT("/me/banned_users/:user_uuid", rt.wrap(rt.banUser))
	r.DELETE("/me/banned_users/:user_uuid", rt.wrap(rt.unbanUser))

	return rt.router
}

func httpSimpleResponse(code int, message []byte, w http.ResponseWriter, ctx reqcontext.RequestContext) {

	w.WriteHeader(code)

	if message == nil || len(message) == 0 {
		return
	}

	_, err := w.Write(message)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't write http request response")
	}

}
