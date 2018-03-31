package cron

import (
	"fmt"
	"net/http"

	"github.com/counting-frontend/database"
)

// DoDailyWork handles any work that needs to happen on a daily basis
// Right now, this entails capturing daily history, clearing daily count,
// and performing daily DB backups
func DoDailyWork(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("In DoDailyWork")

	err := database.InsertDailyHistory()
	if err != nil {
		fmt.Println("ERROR: Can't insert daily history: \n" + err.Error())
		return err
	}
	err = database.ResetDailyUserCounts()
	if err != nil {
		fmt.Println("ERROR: Can't reset daily counts: \n" + err.Error())
		return err
	}
	err = database.PerformDailyBackup()
	if err != nil {
		fmt.Println("ERROR: Can't perform daily db backup: \n" + err.Error())
		return err
	}

	return nil
}

// DoMonthlyWork handles any work that needs to happen on a montly basis
// Right now, this entails capturing monthly history and clearing monthly count.
func DoMonthlyWork(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("In DoMonthlyWork")

	database.InsertMonthlyHistory()
	database.ResetMonthlyUserCounts()

	return nil
}
