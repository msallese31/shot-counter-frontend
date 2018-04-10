package logutil

import (
	"fmt"
	"time"
)

func Log(stringToLog string) {
	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05") + ": " + stringToLog)
}
