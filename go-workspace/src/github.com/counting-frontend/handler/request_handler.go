package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/counting-frontend/backend"
	"github.com/counting-frontend/data"
	"github.com/counting-frontend/types"
)

// RequestHandler is the general request handler for this server.  It determines where
// requests will go
type RequestHandler struct {
}

func (h *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Handling request")
	if r.URL.Path == "/count" {
		// TODO: Real logging
		fmt.Println("Sending to shot counter")
		handleCountRequest(w, r)
	} else if r.URL.Path == "/sign-in" {
		handleSignInRequest(w, r)
	} else {
		// TODO: Create error message -> error code mapping
		w.WriteHeader(http.StatusBadRequest)
		types.SetupAndroidResponse(w, "Url endpoint path not supported", 0)
	}
}

func handleCountRequest(w http.ResponseWriter, r *http.Request) {

	countData := data.CountObject{}
	countData.Request = *r
	countData.Writer = w
	backend.CallShotCounter(countData)
	return
}

func handleSignInRequest(w http.ResponseWriter, r *http.Request) {
	var signInRequest types.SignInRequest
	fmt.Println("Handling sign in request")
	defer r.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	err := json.NewDecoder(r.Body).Decode(&signInRequest)
	switch {
	case err == io.EOF:
		// empty body
		w.WriteHeader(http.StatusBadRequest)
		types.SetupAndroidResponse(w, "Empty request body", 0)
		return
	case err != nil:
		// other error
		w.WriteHeader(http.StatusBadRequest)
		types.SetupAndroidResponse(w, "Bad request", 0)
		return
	}

	return
}

var _ http.Handler = &RequestHandler{}
