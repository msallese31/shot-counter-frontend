package database

import (
	"fmt"
	"time"

	"github.com/counting-frontend/types"
	"gopkg.in/mgo.v2/bson"
)

func InsertDailyHistory() error {
	fmt.Println("In InsertDailyHistory")

	dailyHistoryCollection, err := GetDailyHistoryCollection()
	if err != nil {
		// TODO: Come up with how to handle this
		fmt.Println(err)
		return err
	}

	//CHANGE THE WAY THIS IS STRUCTURED
	usersCollection, err := GetUsersCollection()
	if err != nil {
		// TODO: Come up with how to handle this
		fmt.Println(err)
		return err
	}

	user := &types.User{}
	dailyHistory := &types.DailyHistory{}

	t := time.Now()
	dailyHistory.Date = t.Format("01-02-2006")

	findUsers := usersCollection.Find(bson.M{})
	users := findUsers.Iter()
	for users.Next(&user) {
		dailyHistory.Email = user.Email
		dailyHistory.DailyCount = user.DailyCount
		dailyHistoryCollection.Insert(&dailyHistory)
	}
	fmt.Println("Successfully inserted daily history")
	return nil
}

func InsertMonthlyHistory() {
	fmt.Println("In InsertMonthlyHistory")
}
