package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/util"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getUser(w http.ResponseWriter, _ *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Get request parameters
	var uuid = ps.ByName("uuid")

	// Check if the uuid is empty
	if uuid == "" {
		httpSimpleResponse(http.StatusBadRequest, []byte("uuid is required"), w, ctx)
		return
	}

	// Find requested user
	var user, err = rt.db.GetUser(uuid, database.FilterByUUID)
	if errors.Is(err, sql.ErrNoRows) {

		// User not found
		httpSimpleResponse(http.StatusNotFound, []byte("user not found"), w, ctx)
		return

	} else if err != nil {

		// Generic error
		ctx.Logger.WithError(err).Error("cannot get user")
		httpSimpleResponse(http.StatusInternalServerError, []byte("internal server error"), w, ctx)
		return

	}

	var response []byte
	response, err = json.Marshal(user)
	if err != nil {

		// Error while marshalling the user
		ctx.Logger.WithError(err).Error("cannot marshal user")
		httpSimpleResponse(http.StatusInternalServerError, []byte("internal server error"), w, ctx)
		return
	}

	// Write the response
	httpSimpleResponse(http.StatusOK, response, w, ctx)
}

func (rt *_router) changeUsername(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check authorization
	var userUUID = getAuthorization(r)
	if userUUID == "" {

		// Token not provided or invalid
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	// Find the logged in user
	var user, err = rt.db.GetUser(userUUID, database.FilterByUUID)
	if errors.Is(err, sql.ErrNoRows) {

		// User doesn't exist
		w.WriteHeader(http.StatusUnauthorized)
		return

	} else if err != nil {

		// Generic error
		ctx.Logger.WithError(err).Error("cannot get user")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Parse request body
	var body struct {
		Username string `json:"username"`
	}
	err = json.NewDecoder(r.Body).Decode(&body)

	// Cannot parse body
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return

	}

	// Check validity of new username
	var newUsername = body.Username
	if !util.IsValidUsername(newUsername) {

		w.WriteHeader(http.StatusBadRequest)
		return

	}

	// Check if the username is already taken
	_, err = rt.db.GetUser(newUsername, database.FilterByUsername)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {

		// Generic error
		ctx.Logger.WithError(err).Error("cannot get user")
		w.WriteHeader(http.StatusInternalServerError)
		return

	} else if err == nil {

		// Username is taken
		w.WriteHeader(http.StatusConflict)
		return

	}

	// Update the username
	user.Username = newUsername
	err = rt.db.SetUser(&user)

	if err != nil {

		// Generic error
		ctx.Logger.WithError(err).Error("cannot update user")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Write the response
	w.WriteHeader(http.StatusOK)
}
