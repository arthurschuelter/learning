package main

import (
	"time"
)

type Cinema struct {
	ID    string
	Title string
}

func main() {
	dates := getToday()

	cinemas := []Cinema{
		{ID: "146", Title: "GNC Mueller"},
		{ID: "851", Title: "GNC Garten Shopping"},
	}

	ScrapeGNC(cinemas, dates)
}

func getToday() []string {
	var dates []string
	dates = append(dates, time.Now().Format("2006-01-02"))
	return dates
}

// func getWeek() []string {
// 	var dates []string
// 	today := time.Now()

// 	for i := 0; i < 7; i++ {
// 		date := today.AddDate(0, 0, i).Format("2006-01-02")
// 		dates = append(dates, date)
// 	}

// 	return dates
// }
