package database

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

func DeleteADEntries() error {
	fmt.Println("In DeleteADEntries")

	// 	// Create DB session
	session, err := mgo.Dial("mongodb://main_admin:abc123@mongodb-service")
	defer session.Close()
	if err != nil {
		fmt.Println("Error dialing mongodb: " + err.Error())
	}
	accelDataCollection := session.DB("test").C("accelData")
	if err != nil {
		fmt.Println("Error getting collection from db: " + err.Error())
	}
	accelDataCollection.RemoveAll(nil)
	return err
}

func GetADCount() (int, error) {
	fmt.Println("In GetADCount")

	// Create DB session
	session, err := mgo.Dial("mongodb://main_admin:abc123@mongodb-service")
	defer session.Close()
	if err != nil {
		fmt.Println("Error dialing mongodb: " + err.Error())
		return -1, err
	}

	// Error check here?? TODO: Stop using test database
	accelDataCollection := session.DB("test").C("accelData")
	if err != nil {
		fmt.Println("Error getting collection from db: " + err.Error())
		return -1, err
	}
	return accelDataCollection.Count()
}
