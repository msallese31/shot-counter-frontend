package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/counting-frontend/data"
	"github.com/counting-frontend/types"
)

// CallShotCounter makes a http POST request to the backend shot counting service
func CallShotCounter(countData data.CountObject) {

	var accData *types.AccelerometerData
	w := countData.Writer

	// bodyBytes, _ := ioutil.ReadAll(countData.Request.Body)
	// bodyString := string(bodyBytes)
	// fmt.Println(bodyString)

	defer countData.Request.Body.Close()
	err := json.NewDecoder(countData.Request.Body).Decode(&accData)
	switch {
	case err == io.EOF:
		// empty body
		w.WriteHeader(http.StatusBadRequest)
		types.SetupAndroidResponse(w, "Empty request body", 0)
		return
	case err != nil:
		// other error
		fmt.Println("Unseen error decoding json: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		types.SetupAndroidResponse(w, "Bad request", 0)
		return
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(accData)

	fmt.Println("id token: " + accData.IDToken)
	fmt.Println(accData.X)

	fmt.Println("Calling backend")
	// TODO: Stick configuration stuff in a context
	backendURL := "http://shot-counter-backend:5000/count"
	resp, err := http.Post(backendURL, "application/json; charset=utf-8", b)
	if err != nil {
		fmt.Println("Error talking to backend: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		types.SetupAndroidResponse(w, "Internal server error", 0)
		return
	}

	// TODO: Return json from backend
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	shotsCounted := buf.String() // Does a complete copy of the bytes in the buffer.
	i, err := strconv.Atoi(shotsCounted)
	if err == nil {
		fmt.Println(i)
	}
	fmt.Println("Shots counted from backend: " + shotsCounted)
	fmt.Println("Shots counted from backend: " + strconv.Itoa(i))
	fmt.Println(resp)
	w.WriteHeader(http.StatusOK)
	types.SetupAndroidResponse(w, "", i)

	return
}
