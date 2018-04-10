package database

import (
	"fmt"
	"time"

	"github.com/counting-frontend/types"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func InsertDailyHistory() error {
	fmt.Println("In InsertDailyHistory")

	// 	// Create DB session
	session, err := mgo.Dial("mongodb://main_admin:abc123@mongodb-service")
	defer session.Close()
	if err != nil {
		fmt.Println("Error dialing mongodb: " + err.Error())
		return err
	}
	// Error check here?? TODO: Stop using test database
	dailyHistoryCollection := session.DB("test").C("dailyHistory")

	// Error check here?? TODO: Stop using test database
	usersCollection := session.DB("test").C("users")

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

func InsertMonthlyHistory() error {
	fmt.Println("In InsertMonthlyHistory")

	// Create DB session
	session, err := mgo.Dial("mongodb://main_admin:abc123@mongodb-service")
	defer session.Close()
	if err != nil {
		fmt.Println("Error dialing mongodb: " + err.Error())
		return err
	}

	monthlyHistoryCollection := session.DB("test").C("monthlyHistory")
	usersCollection := session.DB("test").C("users")

	user := &types.User{}
	monthlyHistory := &types.MonthlyHistory{}

	t := time.Now()
	monthlyHistory.Date = t.Format("01-2006")

	findUsers := usersCollection.Find(bson.M{})
	users := findUsers.Iter()
	for users.Next(&user) {
		monthlyHistory.Email = user.Email
		monthlyHistory.MonthlyCount = user.MonthlyCount
		monthlyHistoryCollection.Insert(&monthlyHistory)
	}
	fmt.Println("Successfully inserted monthly history")
	return nil
}
