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

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check authorization
	var user = rt.getAuthorizedUser(r)
	if user == nil {

		// Token not provided or invalid
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	// Get posts
	var posts, err = rt.db.GetPostsByFollowed(user.UUID)
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

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// Get post
	var postUUID = ps.ByName("post_uuid")
	var post database.Post

	post, err = rt.db.GetPost(postUUID)
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

	// Prepare response
	var response []byte
	response, err = json.Marshal(post)
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

func (rt *_router) getPhotoImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
	var isBanned, err = rt.db.IsBanned(loggedUser.UUID, targetUserUUID)
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
	var post database.Post
	post, err = rt.db.GetPost(postUUID)
	if errors.Is(err, sql.ErrNoRows) {

		// Post not found
		w.WriteHeader(http.StatusNotFound)
		return

	} else if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot get post")
		w.WriteHeader(http.StatusInternalServerError)
		return

	} else if post.AuthorUUID != targetUserUUID {

		// Post not from target (post not found)
		w.WriteHeader(http.StatusNotFound)
		return

	}

	// Get image
	var image []byte
	image, err = rt.db.GetImage(post.UUID)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot get image")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Write image
	_, err = w.Write(image)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot write image")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	w.Header().Set("Content-Length", strconv.Itoa(len(image)))
	w.WriteHeader(http.StatusOK)
}

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {
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

		// Unknown error
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

			// Unknown error
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

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot save post")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	w.WriteHeader(http.StatusCreated)
}

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	// Check authorization
	var loggedUser = rt.getAuthorizedUser(r)
	if loggedUser == nil {

		// Token not provided or invalid
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

	// Get post
	var postUUID = ps.ByName("post_uuid")
	var post, err = rt.db.GetPost(postUUID)
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

	// Delete post's comments
	err = rt.db.DeletePostComments(postUUID)

	if err == nil {
		// Delete post's likes
		err = rt.db.DeletePostLikes(postUUID)
	}

	if err == nil {
		// Delete post
		err = rt.db.DeletePost(postUUID)
	}

	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot delete post")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

}
