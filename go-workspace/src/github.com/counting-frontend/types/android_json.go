package types

import (
	"encoding/json"
	"net/http"
	"time"
)

// JSONObject is the interface for dealing with Json types
// type JSONObject interface {
// }

// AccelerometerData is the json data type we use for incoming requests from Android
type AccelerometerData struct {
	// The `json` struct tag maps between the json name
	// and actual name of the field
	X     *[]float32 `json:"X_Acc"`
	Y     *[]float32 `json:"Y_Acc"`
	Z     *[]float32 `json:"Z_Acc"`
	Email string     `json:"email"`
	Date  time.Time
}

// SignInRequest is the data type for when we recieve a /sign-in request from the android application
type SignInRequest struct {
	IDToken string `json:"idToken"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

// AndroidResponse is the json data type we respond to an Android request with
type AndroidResponse struct {
	// The `json` struct tag maps between the json name
	// and actual name of the field
	ShotsCounted int    `json:"shots_counted"`
	Error        string `json:"error"`
}

// SetupAndroidResponse is a helper function to setup an AndroidResponse on a ResponseWriter
func SetupAndroidResponse(w http.ResponseWriter, requestError string, shotsCounted int) {
	respJSON := AndroidResponse{}
	// TODO: Create error message -> error code mapping
	respJSON.Error = requestError
	respJSON.ShotsCounted = shotsCounted
	json.NewEncoder(w).Encode(respJSON)
}
