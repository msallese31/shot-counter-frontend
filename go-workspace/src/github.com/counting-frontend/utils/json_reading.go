package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/counting-frontend/types"
)

// ReadJSONToInterface unmarshalls the data from an io.Reader into some interface
func ReadJSONToInterface(reader io.Reader, in interface{}) (err error) {
	err = json.NewDecoder(reader).Decode(&in)
	if err != nil {
		return err
	}
	return nil
}

// CheckJSONDecodeError checks for an error while reading a body into a json object.
// If it finds an error it will return the appropriate response back to the user.
func CheckJSONDecodeError(err error, w http.ResponseWriter) {
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
}
