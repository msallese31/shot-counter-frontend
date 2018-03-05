package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/counting-frontend/backend"
	"github.com/counting-frontend/data"
	"github.com/counting-frontend/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	} else {
		// TODO: Create error message -> error code mapping
		w.WriteHeader(http.StatusBadRequest)
	}
}

func handleCountRequest(w http.ResponseWriter, r *http.Request) {

	countData := data.CountObject{}
	countData.Request = *r
	countData.Writer = w
	backend.CallShotCounter(countData)
	return
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
	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&signInRequest)
	if err != nil {
		fmt.Println("Something went wrong reading bytes during sign in request: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = signIn(&signInRequest)
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

func signIn(signInRequest *types.SignInRequest) error {
	fmt.Println("handling sign in request")

	// TODO: Stop using mock sign in request
	userFound, err := lookupUser(signInRequest)
	if err != nil {
		errorString := err.Error()
		switch errorString {
		case "not found":
			fmt.Println("User not found with id: " + signInRequest.IDToken)
			// Create User
			err = createUser(signInRequest)
			if err != nil {
				fmt.Println("Unable to create new user: " + err.Error())
				return err
			}
			fmt.Println("Created new user!")
			return nil
		default:
			fmt.Println("Unseen DB error: \n" + errorString)
			return err
		}
	} else {
		fmt.Println("Found user with name: " + userFound.Name)
		return nil
	}

	// Condition on whether or not the user exists
	// We tried with Bo, now lets try with a unknown user.

}

func lookupUser(signInRequest *types.SignInRequest) (types.User, error) {

	lookupUser := types.User{}

	// Error check here?? TODO: Stop using test database
	usersCollection, err := getUsersCollectionFromDB()
	if err != nil {
		fmt.Println(err)
	} else {
		err = usersCollection.Find(bson.M{"_id": signInRequest.IDToken}).One(&lookupUser)
	}
	return lookupUser, err
}

func createUser(signInRequest *types.SignInRequest) error {
	// Create DB session
	userToInsert := types.User{}
	userToInsert.ID = signInRequest.IDToken
	userToInsert.Name = signInRequest.Name
	userToInsert.Email = signInRequest.Email
	userToInsert.DailyCount = 0
	userToInsert.MonthlyCount = 0
	userToInsert.DailyRequests = 0
	userToInsert.MonthlyRequests = 0

	usersCollection, err := getUsersCollectionFromDB()
	if err != nil {
		return err
	}

	err = usersCollection.Insert(userToInsert)
	return err
}

func getUsersCollectionFromDB() (*mgo.Collection, error) {
	// TODO: Take in mongo url from configuration

	// Create DB session
	session, err := mgo.Dial("mongodb://main_admin:abc123@mongodb-service")
	if err != nil {
		fmt.Println("Error dialing mongodb: " + err.Error())
		return &mgo.Collection{}, err
	}
	// Error check here?? TODO: Stop using test database
	usersCollection := session.DB("test").C("users")
	return usersCollection, nil
}

var _ http.Handler = &RequestHandler{}
