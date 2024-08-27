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
	searchedUser, err := rt.db.GetUserFull(userUUID, loggedUser.UUID)
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
	response, err = json.Marshal(searchedUser)
	if err != nil {

		// Error while marshalling the searchedUser
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
		Username    string `json:"username"`
		DisplayName string `json:"displayName"`
	}
	var err = json.NewDecoder(r.Body).Decode(&body)

	// Cannot parse body
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return

	}

	var newUsername = body.Username
	var newDisplayName = body.DisplayName
	var lenNewUsername = len(newUsername)
	var lenNewDisplayName = len(newDisplayName)

	// Both fields are empty
	if lenNewUsername == 0 && lenNewDisplayName == 0 {

		w.WriteHeader(http.StatusBadRequest)
		return

		// Check validity of new username if present
	} else if lenNewUsername > 0 && !util.IsValidUsername(newUsername) {

		w.WriteHeader(http.StatusBadRequest)
		return

		// Check validity of new display name if present
	} else if lenNewDisplayName > 0 && !util.IsValidDisplayName(newDisplayName) {

		w.WriteHeader(http.StatusBadRequest)
		return

	}

	// Handle username
	if lenNewUsername > 0 {

		// Check if the username is already taken
		var userNotExists bool
		userNotExists, err = rt.db.IsUsernameAvailable(newUsername)

		if err != nil {

			// Generic error
			ctx.Logger.WithError(err).Error("cannot get user")
			w.WriteHeader(http.StatusInternalServerError)
			return

		} else if !userNotExists {

			// Username is taken
			w.WriteHeader(http.StatusConflict)
			return

		}

		// Update the username
		user.Username = newUsername

	}

	// Handle display name
	if lenNewDisplayName > 0 {

		// Update the username
		user.DisplayName = newDisplayName

	}

	// Update user object
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

func (rt *_router) setMyProfilePicture(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check authorization
	var loggedUser = rt.getAuthorizedUser(r)
	if loggedUser == nil {

		// Token not provided or invalid
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	// Get post uuid
	var body struct {
		PostUUID string `json:"post_uuid"`
	}
	var err = json.NewDecoder(r.Body).Decode(&body)

	// Cannot parse body
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return

	}

	// Get post
	var post database.Post
	post, err = rt.db.GetPost(body.PostUUID)

	if errors.Is(err, sql.ErrNoRows) {

		// Post not found
		w.WriteHeader(http.StatusNotFound)
		return

	} else if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot get post")
		w.WriteHeader(http.StatusInternalServerError)
		return

	} else if post.AuthorUUID != loggedUser.UUID {

		// Given user is not the author of the post
		w.WriteHeader(http.StatusNotFound)
		return

	}

	// Update the profile picture
	loggedUser.PictureURL = post.ImageURL
	err = rt.db.SetUser(loggedUser)

	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot delete post")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

}

func (rt *_router) findUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check authorization
	var loggedUser = rt.getAuthorizedUser(r)
	if loggedUser == nil {

		// Token not provided or invalid
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	// Get request parameters
	var targetUsername = r.URL.Query().Get("username")

	// Check if the username is empty
	if targetUsername == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Search users
	var users []database.User
	users, err := rt.db.GetUsersWithUsernameSubstr(targetUsername, loggedUser.UUID)
	if err != nil {

		// Generic error
		ctx.Logger.WithError(err).Error("cannot get users")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	var response []byte
	response, err = json.Marshal(users)
	if err != nil {

		// Error while marshalling the searchedUsed
		ctx.Logger.WithError(err).Error("cannot marshal user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write the response
	httpSimpleResponse(http.StatusOK, response, w, ctx)
}

func (rt *_router) deleteSelfUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check authorization
	var loggedUser = rt.getAuthorizedUser(r)
	if loggedUser == nil {

		// Token not provided or invalid
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	// Delete post
	var err = rt.db.DeleteUser(loggedUser.UUID, nil)

	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot delete user")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	w.WriteHeader(http.StatusNoContent)
}
