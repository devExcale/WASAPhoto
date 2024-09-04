package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AppInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

func (rt *_router) getStatus(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	info := AppInfo{
		Name:    "WASAPhoto",
		Version: "1.0.0",
		Status:  "Running",
	}

	var err = json.NewEncoder(w).Encode(info)
	if err != nil {
		rt.baseLogger.WithError(err).Error("cannot marshal app info")
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Default Status OK
}
