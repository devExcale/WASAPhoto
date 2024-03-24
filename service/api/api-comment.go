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

func (rt *_router) getUserPostComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// Get comments
	var comments []database.Comment
	comments, err = rt.db.GetCommentsByPost(postUUID)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot get comments")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	// Prepare response
	var response []byte
	response, err = json.Marshal(comments)
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

func (rt *_router) addUserPostComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// Parse request body
	var body struct {
		Comment string `json:"comment"`
	}
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {

		// Cannot parse body
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	var comment = database.Comment{
		PostUUID:   postUUID,
		AuthorUUID: user.UUID,
		Comment:    body.Comment,
	}

	// Set comment
	err = rt.db.SetComment(&comment)
	if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot set comment")
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	w.WriteHeader(http.StatusCreated)
}

func (rt *_router) removeUserPostComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
	var targetUser, err = rt.db.GetUserBasic(targetUUID)
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

	} else if post.AuthorUUID != targetUser.UUID {

		// Given user is not the author of the post
		w.WriteHeader(http.StatusNotFound)
		return

	}

	// Get comment
	var commentUUID = ps.ByName("comment_uuid")
	var comment database.Comment
	comment, err = rt.db.GetCommentByUUID(commentUUID)
	if errors.Is(err, sql.ErrNoRows) {

		// Comment not found
		w.WriteHeader(http.StatusNotFound)
		return

	} else if err != nil {

		// Unknown error
		ctx.Logger.WithError(err).Error("cannot get comment")
		w.WriteHeader(http.StatusInternalServerError)
		return

	} else if comment.PostUUID != postUUID {

		// Comment does not belong to the post
		w.WriteHeader(http.StatusNotFound)
		return

	} else if comment.AuthorUUID != loggedUser.UUID {

		// Comment is not from the logged user
		w.WriteHeader(http.StatusForbidden)
		return

	}

	// Write the response
	w.WriteHeader(http.StatusOK)

}
