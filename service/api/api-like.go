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

func (rt *_router) getLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
	var _, err = rt.db.GetUserBasicByUUID(targetUUID)
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

	// Get post
	var postUUID = ps.ByName("post_uuid")
	_, err = rt.db.GetPost(postUUID)
	if errors.Is(err, sql.ErrNoRows) {

		// Post not found
		w.WriteHeader(http.StatusNotFound)
		return

	} else if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot get post")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Get users who liked the post
	var users []database.User
	users, err = rt.db.GetUsersLikedPost(postUUID)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot get users who liked post")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Prepare response
	var response []byte
	response, err = json.Marshal(users)
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

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check authorization
	var loggedUser = rt.getAuthorizedUser(r)
	if loggedUser == nil {

		// Token not provided or invalid
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	// Get given user
	var targetUUID = ps.ByName("user_uuid")
	var _, err = rt.db.GetUserBasicByUUID(targetUUID)
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
	isBanned, err = rt.db.IsBanned(loggedUser.UUID, targetUUID)
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

	// Get post
	var postUUID = ps.ByName("post_uuid")
	_, err = rt.db.GetPost(postUUID)
	if errors.Is(err, sql.ErrNoRows) {

		// Post not found
		w.WriteHeader(http.StatusNotFound)
		return

	} else if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot get post")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Add like
	err = rt.db.AddLike(postUUID, loggedUser.UUID)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot add like")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	w.WriteHeader(http.StatusCreated)
}

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check authorization
	var loggedUser = rt.getAuthorizedUser(r)
	if loggedUser == nil {

		// Token not provided or invalid
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	// Get given user
	var targetUUID = ps.ByName("user_uuid")
	var _, err = rt.db.GetUserBasicByUUID(targetUUID)
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
	isBanned, err = rt.db.IsBanned(loggedUser.UUID, targetUUID)
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

	// Get post
	var postUUID = ps.ByName("post_uuid")
	_, err = rt.db.GetPost(postUUID)
	if errors.Is(err, sql.ErrNoRows) {

		// Post not found
		w.WriteHeader(http.StatusNotFound)
		return

	} else if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot get post")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Remove like
	err = rt.db.DeleteLike(postUUID, loggedUser.UUID)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot add like")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	w.WriteHeader(http.StatusCreated)
}
