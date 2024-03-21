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
	rt.router.GET("/users/:uuid", rt.wrap(rt.getUser))
	rt.router.POST("/me/change_username", rt.wrap(rt.changeUsername))

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
