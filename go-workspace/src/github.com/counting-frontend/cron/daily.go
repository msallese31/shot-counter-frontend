package cron

import (
	"fmt"
	"net/http"
)

func DoDailyWork(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("In DoDailyWork")
	return nil
}
