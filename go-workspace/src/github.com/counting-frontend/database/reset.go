package database

import (
	"fmt"

	"github.com/counting-frontend/types"
	"gopkg.in/mgo.v2/bson"
)

// ResetMonthlyUserCounts sets the montly_count field in the users collection to 0
func ResetMonthlyUserCounts() error {
	fmt.Println("In ResetMonthlyUserCounts")
	err := resetFieldForAllUsers("monthly_count")
	return err
}

// ResetDailyUserCounts sets the daily_count field in the users collection to 0
func ResetDailyUserCounts() error {
	fmt.Println("In ResetDailyUserCounts")
	err := resetFieldForAllUsers("daily_count")
	return err
}

func resetFieldForAllUsers(field string) error {
	usersCollection, err := GetUsersCollection()
	if err != nil {
		fmt.Println("ERROR: Couldn't get user collection:\n" + err.Error())
		return err
	}

	findUsers := usersCollection.Find(bson.M{})
	user := &types.User{}
	users := findUsers.Iter()
	change := bson.M{"$set": bson.M{field: 0}}
	for users.Next(&user) {
		err := usersCollection.Update(bson.M{"email": user.Email}, change)
		if err != nil {
			// TODO: Keep going if it's just one user
			fmt.Println("ERROR: Couldn't reset" + field + " for user: " + user.Email + ". Giving up. \n" + err.Error())
			return err
		}
	}
	fmt.Println("Successfully reset " + field + " for all users.")
	return nil
}
