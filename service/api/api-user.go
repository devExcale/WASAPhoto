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

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check authorization
	var loggedUser = rt.getAuthorizedUser(r)
	if loggedUser == nil {
		// Token not provided or invalid
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get request parameters
	var userUUID = ps.ByName("user_uuid")

	// Check if the userUUID is empty
	if userUUID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check bans
	var isBanned, err = rt.db.IsBanned(loggedUser.UUID, userUUID)
	if err != nil {

		ctx.Logger.WithError(err).Error("cannot check ban")
		w.WriteHeader(http.StatusInternalServerError)
		return

	} else if isBanned {

		w.WriteHeader(http.StatusForbidden)
		return

	}

	// Find requested user
	searchedUsed, err := rt.db.GetUserFull(userUUID, database.FilterByUUID)
	if errors.Is(err, sql.ErrNoRows) {

		// User not found
		w.WriteHeader(http.StatusNotFound)
		return

	} else if err != nil {

		// Generic error
		ctx.Logger.WithError(err).Error("cannot get user")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	var response []byte
	response, err = json.Marshal(searchedUsed)
	if err != nil {

		// Error while marshalling the searchedUsed
		ctx.Logger.WithError(err).Error("cannot marshal user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write the response
	httpSimpleResponse(http.StatusOK, response, w, ctx)
}

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check authorization
	var user = rt.getAuthorizedUser(r)
	if user == nil {

		// Token not provided or invalid
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	// Parse request body
	var body struct {
		Username string `json:"username"`
	}
	var err = json.NewDecoder(r.Body).Decode(&body)

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
	_, err = rt.db.GetUserFull(newUsername, database.FilterByUsername)

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
	err = rt.db.SetUser(user)

	if err != nil {

		// Generic error
		ctx.Logger.WithError(err).Error("cannot update user")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Write the response
	w.WriteHeader(http.StatusOK)
}
