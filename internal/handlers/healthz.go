package handlers

import (
	"encoding/json"
	"mime"
	"net/http"

	"github.com/ale8k/language_usage_by_so/pkg/errors"
)

func getServerHealth(resp http.ResponseWriter, req *http.Request) {
	data, err := json.Marshal(struct {
		Healthy bool `json:"healthy"`
	}{true})
	// Kill this routine if any error occurs
	errors.HandleErrFatal(err)
	resp.Header().Add("Content-Type", mime.TypeByExtension(".json"))
	resp.Write(data)
}
