package database

import (
	"fmt"

	"github.com/counting-frontend/types"
	"gopkg.in/mgo.v2/bson"
)

func ResetMonthlyUserCounts() error {
	fmt.Println("stub for now")
	return nil
}

func ResetDailyUserCounts() error {
	fmt.Println("In ResetDailyUserCounts")
	usersCollection, err := GetUsersCollection()
	if err != nil {
		fmt.Println("ERROR: Couldn't get user collection:\n" + err.Error())
		return err
	}

	findUsers := usersCollection.Find(bson.M{})
	user := &types.User{}
	users := findUsers.Iter()
	change := bson.M{"$set": bson.M{"daily_count": 0}}
	for users.Next(&user) {
		err := usersCollection.Update(bson.M{"email": user.Email}, change)
		if err != nil {
			// TODO: Keep going if it's just one user
			fmt.Println("ERROR: Couldn't reset daily count for user: " + user.Email + ". Giving up. \n" + err.Error())
			return err
		}
	}
	fmt.Println("Successfully reset daily count for all users.")
	return nil
}
