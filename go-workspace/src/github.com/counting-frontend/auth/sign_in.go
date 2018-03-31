package auth

import (
	"fmt"

	"github.com/counting-frontend/database"
	"github.com/counting-frontend/types"
	"gopkg.in/mgo.v2/bson"
)

func CreateUser(signInRequest *types.SignInRequest) error {
	// Create DB session
	userToInsert := types.User{}
	userToInsert.ID = signInRequest.IDToken
	userToInsert.Name = signInRequest.Name
	userToInsert.Email = signInRequest.Email
	userToInsert.DailyCount = 0
	userToInsert.MonthlyCount = 0
	userToInsert.DailyRequests = 0
	userToInsert.MonthlyRequests = 0
	userToInsert.AccountType = "free"

	usersCollection, err := database.GetUsersCollection()
	if err != nil {
		return err
	}

	err = usersCollection.Insert(userToInsert)
	return err
}

func SignIn(signInRequest *types.SignInRequest) error {
	fmt.Println("handling sign in request")

	userFound, err := LookupUser(signInRequest)
	if err != nil {
		errorString := err.Error()
		switch errorString {
		case "not found":
			fmt.Println("User not found with id: " + signInRequest.IDToken + " Name: " + signInRequest.Name + " Email: " + signInRequest.Email)
			// Create User
			err = CreateUser(signInRequest)
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

func LookupUser(signInRequest *types.SignInRequest) (types.User, error) {

	lookupUser := types.User{}

	// Error check here?? TODO: Stop using test database
	usersCollection, err := database.GetUsersCollection()
	if err != nil {
		fmt.Println(err)
	} else {
		err = usersCollection.Find(bson.M{"email": signInRequest.Email}).One(&lookupUser)
	}
	return lookupUser, err
}
