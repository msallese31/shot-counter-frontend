package backend

import (
	"fmt"
	"net/http"

	"github.com/counting-frontend/data"
	"github.com/counting-frontend/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetShotCount retrieve's a user's shot count for today from the DB
func GetShotCount(countData data.CountObject) {

	// TODO: Stuff idToken in countData;  right now idToken is in AccelerometerData which makes this difficult
	email := countData.Request.URL.Query().Get("email")

	// Create DB session
	session, err := mgo.Dial("mongodb://main_admin:abc123@mongodb-service")
	defer session.Close()
	if err != nil {
		// TODO: What do we return to the user here?
		fmt.Println("Error dialing mongodb: " + err.Error())
	}

	colQuerier := bson.M{"email": email}
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
			countData.Writer.WriteHeader(http.StatusNotFound)
			types.SetupAndroidResponse(countData.Writer, "Error:  email either not supplied or not found", -1)
			return
		}

		countData.Writer.WriteHeader(http.StatusOK)
		types.SetupAndroidResponse(countData.Writer, "", user.DailyCount)

	}
}
