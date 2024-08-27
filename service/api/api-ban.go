package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getBannedUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check authorization
	var loggedUser = rt.getAuthorizedUser(r)
	if loggedUser == nil {

		// Token not provided or invalid
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	// Get banned users
	var bannedUsers, err = rt.db.GetBannedUsers(loggedUser.UUID)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot get banned users")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Prepare response
	var response []byte
	response, err = json.Marshal(bannedUsers)
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

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// Ban user
	err = rt.db.AddBan(loggedUser.UUID, targetUUID)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot add ban")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Remove follow from given user
	err = rt.db.DeleteFollow(targetUUID, loggedUser.UUID)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot remove follow")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	w.WriteHeader(http.StatusCreated)
}

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// Remove ban
	err = rt.db.DeleteBan(loggedUser.UUID, targetUUID)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot remove ban")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	w.WriteHeader(http.StatusOK)
}
