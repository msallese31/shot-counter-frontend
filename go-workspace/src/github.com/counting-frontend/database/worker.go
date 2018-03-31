package database

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

// GetUsersCollection sets up a session with our mnogo DB and returns the "users" collection
func GetUsersCollection() (*mgo.Collection, error) {
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

// GetDailyHistoryCollection sets up a session with our mnogo DB and returns the "dailyHistory" collection
func GetDailyHistoryCollection() (*mgo.Collection, error) {
	// TODO: Take in mongo url from configuration

	// Create DB session
	session, err := mgo.Dial("mongodb://main_admin:abc123@mongodb-service")
	if err != nil {
		fmt.Println("Error dialing mongodb: " + err.Error())
		return &mgo.Collection{}, err
	}
	// Error check here?? TODO: Stop using test database
	dailyHistoryCollection := session.DB("test").C("dailyHistory")
	return dailyHistoryCollection, nil
}

// GetMonthlyHistoryCollection sets up a session with our mongo DB and returns the "monthlyHistory" collection
func GetMonthlyHistoryCollection() (*mgo.Collection, error) {
	// TODO: Take in mongo url from configuration

	// Create DB session
	session, err := mgo.Dial("mongodb://main_admin:abc123@mongodb-service")
	if err != nil {
		fmt.Println("Error dialing mongodb: " + err.Error())
		return &mgo.Collection{}, err
	}
	// Error check here?? TODO: Stop using test database
	monthlyHistoryCollection := session.DB("test").C("monthlyHistory")
	return monthlyHistoryCollection, nil
}

func PerformDailyBackup() error {
	fmt.Println("stub for now")
	return nil
}
