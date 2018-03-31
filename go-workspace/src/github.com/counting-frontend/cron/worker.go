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

	database.InsertDailyHistory()
	database.ResetDailyUserCounts()
	database.PerformDailyBackup()

	return nil
}

// DoDailyWork handles any work that needs to happen on a montly basis
// Right now, this entails capturing monthly history and clearing monthly count.
func DoMonthlyWork(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("In DoDailyWork")

	database.InsertMonthlyHistory()
	database.ResetMonthlyUserCounts()

	return nil
}