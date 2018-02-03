package data

import "net/http"

// CountObject is meant to hold information about a "/count" request
type CountObject struct {
	Request      http.Request
	shotsCounted int
	errorMessage error
	// jsonBody     string json
}

// SetCount is a setter method for the shot count of a CountObject
func (c *CountObject) SetCount(count int) {
	c.shotsCounted = count
}

// GetCount is a getter method for the shot count of a CountObject
func (c *CountObject) GetCount() (count int) {
	return c.shotsCounted
}
