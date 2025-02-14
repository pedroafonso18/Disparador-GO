package services

import (
	"fmt"
	"time"
)

func IsCurrentTimeBetween(start, end string) (bool, error) {
	now := time.Now()

	today := now.Format("2006-01-02")
	loc := now.Location()
	startTime, err := time.ParseInLocation("2006-01-02 15:04", today+" "+start, loc)
	if err != nil {
		return false, fmt.Errorf("failed to parse start time: %w", err)
	}
	endTime, err := time.ParseInLocation("2006-01-02 15:04", today+" "+end, loc)
	if err != nil {
		return false, fmt.Errorf("failed to parse end time: %w", err)
	}

	return now.After(startTime) && now.Before(endTime), nil
}

func ReturnTime() bool {
	dt := time.Now()
	fmt.Println(dt.Format("01-02-2006 15:04"))

	inRange, err := IsCurrentTimeBetween("08:00", "17:00")
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return inRange
}
