package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/counting-frontend/data"
	"github.com/counting-frontend/types"
)

// CallShotCounter makes a http POST request to the backend shot counting service
func CallShotCounter(countData data.CountObject) {
	fmt.Println("Calling backend")

	var accData types.AccelerometerData

	err := json.NewDecoder(countData.Request.Body).Decode(&accData)
	if err != nil {
		fmt.Println(err)
		return
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(accData)

	// TODO: Stick configuration stuff in a context
	backendURL := "http://localhost:5000/count"
	req, err := http.Post(backendURL, "application/json; charset=utf-8", b)
	if err != nil {
		fmt.Printf("http.NewRequest() error: %v\n", err)
		return
	}
	fmt.Println(req)

	// c := &http.Client{}
	// resp, err := c.Do(req)
	// if err != nil {
	// 	fmt.Printf("http.Do() error: %v\n", err)
	// 	return
	// }
	// defer resp.Body.Close()
	// data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("ioutil.ReadAll() error: %v\n", err)
	// 	return
	// }

	// fmt.Printf("read resp.Body successfully:\n%v\n", string(data))
	return
}
