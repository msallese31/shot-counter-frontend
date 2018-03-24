package data

import "net/http"
import "github.com/counting-frontend/types"

// CountObject is meant to hold information about a "/count" request
type CountObject struct {
	Request           http.Request
	Writer            http.ResponseWriter
	ShotsCounted      int
	errorMessage      error
	AccelerometerData *types.AccelerometerData
	// jsonBody     string json
}
