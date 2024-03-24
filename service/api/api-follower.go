package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getFollowedUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check authorization
	var user = rt.getAuthorizedUser(r)
	if user == nil {

		// Token not provided or invalid
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	// Get followed users
	var followedUsers, err = rt.db.GetFollowedUsers(user.UUID)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot get followed users")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Prepare response
	var response []byte
	response, err = json.Marshal(followedUsers)
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

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// Follow user
	err = rt.db.AddFollow(user.UUID, targetUUID)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot follow user")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Write the response
	w.WriteHeader(http.StatusOK)
}

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// Follow user
	err = rt.db.DeleteFollow(user.UUID, targetUUID)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot unfollow user")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Write the response
	w.WriteHeader(http.StatusOK)

}
