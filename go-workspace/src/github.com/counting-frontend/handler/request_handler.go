package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/counting-frontend/auth"
	"github.com/counting-frontend/backend"
	"github.com/counting-frontend/cron"
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
	} else if r.URL.Path == "/readiness" {
		w.WriteHeader(http.StatusOK)
	} else if r.URL.Path == "/AS5Hr6Aoay" {
		t := time.Now()
		fmt.Println("TEST DB DAILY RESET:  " + t.Format("3:04PM"))
	} else if r.URL.Path == "/pdeuPqVVnL" {
		// This endpoint should be it once daily.  It's backed by a k8's "cronJob".
		// This is where our logic for daily DB resets will go.
		handleDailyEndpoint(w, r)
	} else if r.URL.Path == "/6KVmH3cjgX" {
		// This endpoint should be hit once a month by k8's cronjob.
		// This is where we will capture & clear monthly history.
		handleMonthlyEndpoint(w, r)
	} else {
		// TODO: Create error message -> error code mapping
		w.WriteHeader(http.StatusBadRequest)
	}
}

func handleCountRequest(w http.ResponseWriter, r *http.Request) {

	countData := data.CountObject{}
	countData.Request = *r
	countData.Writer = w
	switch r.Method {
	case http.MethodGet:
		backend.GetShotCount(countData)
	case http.MethodPost:
		backend.CallShotCounter(countData)
	}
	return
}

func handleDailyEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In handleDailyEndpoint")
	t := time.Now()
	fmt.Println("DAILY ENDPOINT HIT!! " + t.Format("3:04PM"))
	err := cron.DoDailyWork(w, r)
	if err != nil {
		// TODO: Figure out how to handle errors
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("Daily endpoint work completed successfully")
	w.WriteHeader(http.StatusOK)
}

func handleMonthlyEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In handleMonthlyEndpoint")
	t := time.Now()
	fmt.Println("MONTHLY ENDPOINT HIT!! " + t.Format("3:04PM"))
	err := cron.DoMonthlyWork(w, r)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("Monthly endpoint work completed successfully")
	w.WriteHeader(http.StatusOK)
}

func handleSignInRequest(w http.ResponseWriter, r *http.Request) {
	// TODO: ADD check here to make sure the request is coming from our android app.
	// This will help prevent a DDOS attack
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var signInRequest types.SignInRequest
	fmt.Println("Handling sign in request")

	if r.Body == nil {
		fmt.Println("No body provided in sign-in request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&signInRequest)
	if err != nil {
		fmt.Println("Something went wrong reading bytes during sign in request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = auth.SignIn(&signInRequest)
	switch err {
	case nil:
		w.WriteHeader(http.StatusOK)
		return
	default:
		// other error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

var _ http.Handler = &RequestHandler{}
