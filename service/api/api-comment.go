package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getUserPostComments(w http.ResponseWriter, _ *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check authorization
	var user = rt.getAuthorizedUser(r)
	if user == nil {

		// Token not provided or invalid
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	// Get given user
	var targetUUID = ps.ByName("user_uuid")
	var _, err = rt.db.GetUserBasic(targetUUID)
	if errors.Is(err, sql.ErrNoRows) {

		// User not found
		w.WriteHeader(http.StatusNotFound)
		return

	} else if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot get user")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Check ban
	var isBanned bool
	isBanned, err = rt.db.IsBanned(user.UUID, targetUUID)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot check if user is banned")
		w.WriteHeader(http.StatusInternalServerError)
		return

	} else if isBanned {

		// User is banned
		w.WriteHeader(http.StatusForbidden)
		return

	}

	// Get posts
	var posts []database.Post
	posts, err = rt.db.GetPostsByUser(targetUUID)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot get posts")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Prepare response
	var response []byte
	response, err = json.Marshal(posts)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot marshal response")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Write response
	_, err = w.Write(response)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot write response")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	w.WriteHeader(http.StatusOK)
}

func (rt *_router) addUserPostComment(w http.ResponseWriter, _ *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// TODO
}

func (rt *_router) removeUserPostComment(w http.ResponseWriter, _ *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// TODO w
}
