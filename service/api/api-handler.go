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
	r.GET("/users/:user_uuid", rt.wrap(rt.getUserProfile))
	r.POST("/me/change_username", rt.wrap(rt.setMyUserName))

	// Post routes
	r.GET("/me/feed", rt.wrap(rt.getMyStream))
	r.PUT("/me/feed", rt.wrap(rt.addPost))
	r.GET("/users/:user_uuid/feed", rt.wrap(rt.getUserFeed))
	r.GET("/users/:user_uuid/feed/:post_uuid", rt.wrap(rt.getUserPost))
	r.GET("/users/:user_uuid/feed/:post_uuid/webp", rt.wrap(rt.getPostImage))

	// Comment routes
	r.GET("/users/:user_uuid/feed/:post_uuid/comments", rt.wrap(rt.getUserPostComments))
	r.PUT("/users/:user_uuid/feed/:post_uuid/comments", rt.wrap(rt.addUserPostComment))
	r.DELETE("/users/:user_uuid/feed/:post_uuid/comments/:comment_uuid", rt.wrap(rt.removeUserPostComment))

	// Likes routes
	r.GET("/users/:user_uuid/feed/:post_uuid/likes", rt.wrap(rt.getUserPostLikes))
	r.PUT("/users/:user_uuid/feed/:post_uuid/likes", rt.wrap(rt.addUserPostLike))
	r.DELETE("/users/:user_uuid/feed/:post_uuid/likes", rt.wrap(rt.removeUserPostLike))

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
