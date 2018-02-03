package handler

import (
	"fmt"
	"net/http"

	"github.com/counting-frontend/backend"
	"github.com/counting-frontend/data"
	"github.com/counting-frontend/types"
)

type RequestHandler struct {
}

func (h *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Handling request")
	if r.URL.Path == "/count" {
		fmt.Println("Sending to shot counter")
		handleCountRequest(w, r)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func handleCountRequest(w http.ResponseWriter, r *http.Request) {
	var accData types.AccelerometerData
	if r.Body == nil {
		http.Error(w, "No accelerometer data was received", 400)
		return
	}

	// err := json.NewDecoder(r.Body).Decode(&accData)
	// if err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, err.Error(), 400)
	// 	return
	// }

	fmt.Println(accData.X)
	//
	// bodyBytes, _ := ioutil.ReadAll(r.Body)
	// bodyString := string(bodyBytes)
	// fmt.Println(string(bodyString))
	// err := json.NewDecoder(r.Body).Decode(&u)
	// if err != nil {
	// http.Error(w, err.Error(), 400)
	// }
	// fmt.Println(u.Id)

	countData := data.CountObject{}
	countData.Request = *r
	// TODO: Should this be a handler?
	backend.CallShotCounter(countData)
	return
}

var _ http.Handler = &RequestHandler{}
