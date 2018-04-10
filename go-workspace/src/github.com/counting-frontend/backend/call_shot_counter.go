package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/counting-frontend/data"
	"github.com/counting-frontend/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CallShotCounter makes a http POST request to the backend shot counting service
func CallShotCounter(countData data.CountObject) {
	var accData *types.AccelerometerData
	w := countData.Writer

	if countData.Request.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
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

	fmt.Println("email: " + accData.Email)
	countData.AccelerometerData = accData
	fmt.Println("Calling backend")
	// TODO: Stick configuration stuff in a context
	backendURL := "http://shot-counter-backend:5000/count"
	resp, err := http.Post(backendURL, "application/json; charset=utf-8", b)
	defer resp.Body.Close()

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
	shotsCountedInt, err := strconv.Atoi(shotsCounted)
	if err != nil {
		fmt.Println("Error converting shotsCounted to an int: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		types.SetupAndroidResponse(w, "Internal server error", 0)
	}
	fmt.Println("Shots counted from backend: " + shotsCounted)
	fmt.Println("Shots counted Int: " + string(shotsCountedInt))
	fmt.Println("shots counted for id: " + countData.AccelerometerData.Email + "; count: " + shotsCounted + "\nTHIS NEEDS TO BE THE SAME AS THE NEXT ONE")

	countData.ShotsCounted = shotsCountedInt
	fmt.Println(resp)
	w.WriteHeader(http.StatusOK)
	types.SetupAndroidResponse(w, "", shotsCountedInt)
	// We need to be extra careful here that we're not going to end up setting the count again before we submit the data
	go submitDataToDB(&countData)
	if countData.ShotsCounted > 0 {
		go incrementCountInDB(&countData)
	} else {
		fmt.Println("Not incrementing count because we didn't find any shots")
	}

	return
}

func submitDataToDB(countData *data.CountObject) {
	fmt.Println(&countData.AccelerometerData.Email)
	fmt.Println("shots counted for email: " + countData.AccelerometerData.Email + "; count: " + strconv.Itoa(countData.ShotsCounted) + "\nTHIS IS THE NEXT ONE")

	dataToSubmit := types.AccelerometerData{}
	dataToSubmit.Email = countData.AccelerometerData.Email
	dataToSubmit.X = countData.AccelerometerData.X
	dataToSubmit.Y = countData.AccelerometerData.Y
	dataToSubmit.Z = countData.AccelerometerData.Z
	dataToSubmit.Date = time.Now()

	// Create DB session
	session, err := mgo.Dial("mongodb://main_admin:abc123@mongodb-service")
	defer session.Close()
	if err != nil {
		// TODO: What do we return to the user here?
		fmt.Println("Error dialing mongodb: " + err.Error())
	}
	// Error check here?? TODO: Stop using test database
	accelDataCollection := session.DB("test").C("accelData")
	if err != nil {
		fmt.Println("Error getting collection from db: " + err.Error())
	} else {
		err = accelDataCollection.Insert(&dataToSubmit)
		if err != nil {
			fmt.Println("Error inserting data into DB: " + err.Error())
		}
	}

}

func incrementCountInDB(countData *data.CountObject) {
	// Create DB session
	session, err := mgo.Dial("mongodb://main_admin:abc123@mongodb-service")
	defer session.Close()
	if err != nil {
		// TODO: What do we return to the user here?
		fmt.Println("Error dialing mongodb: " + err.Error())
	}

	colQuerier := bson.M{"email": countData.AccelerometerData.Email}
	fmt.Println(colQuerier)
	user := types.User{}
	// Error check here?? TODO: Stop using test database

	usersCollection := session.DB("test").C("users")
	if err != nil {
		fmt.Println("Error getting collection from db: " + err.Error())
	} else {
		err = usersCollection.Find(colQuerier).One(&user)
		if err != nil {
			fmt.Println("Error finding user: " + err.Error())
			return
		}
		change := bson.M{"$inc": bson.M{"daily_count": countData.ShotsCounted, "monthly_count": countData.ShotsCounted}}

		err = usersCollection.Update(colQuerier, change)
		if err != nil {
			fmt.Println("Error incrementing count: " + err.Error())
		}
		fmt.Println(user.Email)
	}

}
