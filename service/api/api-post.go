package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"strconv"
)

func (rt *_router) getMyStream(w http.ResponseWriter, _ *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// TODO
}

func (rt *_router) getUserFeed(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

		ctx.Logger.WithError(err).Error("cannot get user")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Get posts
	var posts []database.Post
	posts, err = rt.db.GetPostsByUser(targetUUID)
	if err != nil {

		ctx.Logger.WithError(err).Error("cannot get posts")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Prepare response
	var response []byte
	response, err = json.Marshal(posts)
	if err != nil {

		ctx.Logger.WithError(err).Error("cannot marshal JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Write response
	_, err = w.Write(response)
	if err != nil {

		ctx.Logger.WithError(err).Error("cannot write response")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	w.WriteHeader(http.StatusOK)
}

func (rt *_router) getUserPost(w http.ResponseWriter, _ *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// TODO
}

// TODO: add to yaml
func (rt *_router) getPostImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-Type", "image/webp")

	// Check authorization
	var loggedUser = rt.getAuthorizedUser(r)
	if loggedUser == nil {

		// Token not provided or invalid
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	// Check if loggedUser is banned
	var targetUserUUID = ps.ByName("user_uuid")
	var isBanned, err = rt.db.IsBanned(targetUserUUID, loggedUser.UUID)
	if err != nil {

		ctx.Logger.WithError(err).Error("cannot check if user is banned")
		w.WriteHeader(http.StatusInternalServerError)
		return

	} else if isBanned {

		// User is banned
		w.WriteHeader(http.StatusForbidden)
		return

	}

	// Get image
	var image []byte
	var postUUID = ps.ByName("post_uuid")
	image, err = rt.db.GetImage(postUUID)
	if errors.Is(err, sql.ErrNoRows) {

		// Post not found
		w.WriteHeader(http.StatusNotFound)
		return

	} else if err != nil {

		ctx.Logger.WithError(err).Error("cannot get image")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Write image
	_, err = w.Write(image)
	if err != nil {

		ctx.Logger.WithError(err).Error("cannot write image")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	w.Header().Set("Content-Length", strconv.Itoa(len(image)))
	w.WriteHeader(http.StatusOK)
}

func (rt *_router) addPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check authorization
	var user = rt.getAuthorizedUser(r)
	if user == nil {

		// Token not provided or invalid
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	// Get file
	var file, _, err = r.FormFile("file")
	if file == nil || err != nil {

		// No file provided
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	// Get file content
	var fileContent []byte
	fileContent, err = io.ReadAll(file)
	if err != nil {

		ctx.Logger.WithError(err).Error("cannot read file")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Get caption
	var bodyText = r.FormValue("body")
	var body = struct {
		Caption string `json:"caption"`
	}{}

	if bodyText != "" {

		// Parse JSON
		err = json.Unmarshal([]byte(bodyText), &body)
		if err != nil {

			ctx.Logger.WithError(err).Error("cannot parse JSON")
			w.WriteHeader(http.StatusBadRequest)
			return

		}
	}

	var post = database.Post{
		AuthorUUID: user.UUID,
		Caption:    body.Caption,
	}

	// Upload post
	err = rt.db.SetPost(&post, fileContent)
	if err != nil {

		ctx.Logger.WithError(err).Error("cannot save post")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	w.WriteHeader(http.StatusCreated)
}
