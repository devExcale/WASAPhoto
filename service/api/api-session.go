package api

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/util"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params, ctx reqcontext.RequestContext) {

	var body = struct {
		Username string `json:"username"`
	}{}

	// Get username from request body
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {

		// Cannot parse the request body
		httpSimpleResponse(http.StatusBadRequest, []byte("invalid request body"), w, ctx)
		return

	}

	var username = body.Username

	// Check if the username is valid
	if !util.IsValidUsername(username) {

		httpSimpleResponse(http.StatusBadRequest, []byte("invalid username"), w, ctx)
		return

	}

	// Check if user exists
	var user database.User
	user, err = rt.db.GetUserBasicByUsername(username)

	// User not found
	if errors.Is(err, sql.ErrNoRows) {

		// Create user
		user = database.User{
			Username:    username,
			DisplayName: username,
		}

		err = rt.db.SetUser(&user)
	}

	if err != nil {

		// Generic error
		rt.baseLogger.WithError(err).Error("cannot save new user")
		httpSimpleResponse(http.StatusInternalServerError, []byte("internal server error"), w, ctx)
		return

	}

	// Compute authorization token
	var respObj = struct {
		Token    string `json:"token"`
		UserUUID string `json:"user_uuid"`
	}{}
	respObj.Token = base64.StdEncoding.EncodeToString([]byte(user.UUID))
	respObj.UserUUID = user.UUID

	var response []byte
	response, err = json.Marshal(respObj)

	// Error while marshalling the response
	if err != nil {

		rt.baseLogger.WithError(err).Error("cannot marshal response")
		httpSimpleResponse(http.StatusInternalServerError, []byte("internal server error"), w, ctx)
		return

	}

	// Write the response
	httpSimpleResponse(http.StatusOK, response, w, ctx)

}

// getAuthorizedUser returns minimal information about the authenticated user.
func (rt *_router) getAuthorizedUser(r *http.Request) *database.User {

	// Get token
	var token = r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		return nil
	}

	// Decode token
	var uuidByte, err = base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil
	}

	// Get user
	var user database.User
	user, err = rt.db.GetUserBasicByUUID(string(uuidByte))
	if err != nil {
		return nil
	}

	return &user
}

/* Commented unused code in case it comes in handy in the future
// getAuthorizedUserToken returns minimal information about the authenticated user.
func (rt *_router) getAuthorizedUserToken(token string) *database.User {

	// Prepare token
	token = strings.TrimPrefix(token, "Bearer ")
	if token == "" {
		return nil
	}

	// Decode token
	var uuidByte, err = base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil
	}

	// Get user
	var user database.User
	user, err = rt.db.GetUserBasicByUUID(string(uuidByte))
	if err != nil {
		return nil
	}

	return &user
}
*/
