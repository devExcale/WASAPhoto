package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// Register routes
	rt.router.GET("/", rt.getAppInfo)

	// Special routes
	rt.router.GET("/liveness", rt.liveness)
	rt.router.POST("/doLogin", rt.wrap(rt.doLogin))

	// User routes
	rt.router.GET("/users/:user_uuid", rt.wrap(rt.getUserProfile))
	rt.router.POST("/me/change_username", rt.wrap(rt.setMyUserName))

	// Post routes
	rt.router.GET("/me/feed", rt.wrap(rt.getMyStream))
	rt.router.GET("/users/:user_uuid/feed", rt.wrap(rt.getUserFeed))
	rt.router.GET("/users/:user_uuid/feed/:post_uuid", rt.wrap(rt.getUserPost))

	// Comment routes
	rt.router.GET("/users/:user_uuid/feed/:post_uuid/comments", rt.wrap(rt.getUserPostComments))
	rt.router.PUT("/users/:user_uuid/feed/:post_uuid/comments", rt.wrap(rt.addUserPostComment))
	rt.router.DELETE("/users/:user_uuid/feed/:post_uuid/comments/:comment_uuid", rt.wrap(rt.removeUserPostComment))

	// Comment routes
	rt.router.GET("/users/:user_uuid/feed/:post_uuid/likes", rt.wrap(rt.getUserPostLikes))
	rt.router.PUT("/users/:user_uuid/feed/:post_uuid/likes", rt.wrap(rt.addUserPostLike))
	rt.router.DELETE("/users/:user_uuid/feed/:post_uuid/likes", rt.wrap(rt.removeUserPostLike))

	// Follow routes
	rt.router.GET("/me/followed_users", rt.wrap(rt.getFollowedUsers))
	rt.router.PUT("/me/followed_users/:user_uuid", rt.wrap(rt.followUser))
	rt.router.DELETE("/me/followed_users/:user_uuid", rt.wrap(rt.unfollowUser))

	// Ban routes
	rt.router.GET("/me/banned_users", rt.wrap(rt.getBannedUsers))
	rt.router.PUT("/me/banned_users/:user_uuid", rt.wrap(rt.banUser))
	rt.router.DELETE("/me/banned_users/:user_uuid", rt.wrap(rt.unbanUser))

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
