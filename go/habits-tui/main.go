package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"slices"
	"time"
)

type Habit struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Entries     []HabitEntry `json:"entries"`
	Streak      int          `json:"streak"`
}

type HabitStreak struct {
	HabitEntryDate string
	IsActive       bool
}

type HabitEntry struct {
	EntryDate string `json:"entry_date"`
	HabitId   string `json:"habit_id"`
	Duration  int    `json:"duration"`
}

type EntryRequest struct {
	HabitId   string `json:"habit_id"`
	EntryDate string `json:"entry_date"`
	Duration  int    `json:"duration"`
}

func main() {
	habits := GetHabits()
	for {

		var command int
		PrintMenu()

		_, err := fmt.Scan(&command)
		if err != nil {
			fmt.Println("Bye bye")
			return
		}

		if command == 1 {
			fmt.Println("Fetching habits...")
			habits := GetHabits()
			printFetchHabits(habits)
		}

		if command == 2 {
			HabitEntryOptionsMenu(habits)
		}

		if command >= 4 {
			fmt.Println("Bye bye")
			return
		}
	}
}

func PrintMenu() {
	fmt.Println("=======================")
	fmt.Println("(1) Fetch habits")
	fmt.Println("(2) Add habit entry")
	fmt.Println("(4) Exit")
	fmt.Print("Enter command: ")
}

func GetHabits() []Habit {
	base_url := "http://localhost:8000/api/habits"
	response, err := http.Get(base_url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	var habits []Habit
	err = json.Unmarshal(body, &habits)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	return habits
}

func PostHabitEntry(entryRequest EntryRequest) {
	base_url := "http://localhost:8000/api/habit_entries"
	jsonData, err := json.Marshal(entryRequest)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	response, err := http.Post(base_url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()
	fmt.Println("Habit successfully logged!")
}

func PrintHabitsOptions(habits []Habit) {
	for i, habit := range habits {
		fmt.Printf("(%d) %s\n", i+1, habit.Name)
	}
}

func HabitEntryOptionsMenu(habits []Habit) {
	var habitChoice int
	clear()
	fmt.Println("=======================")
	PrintHabitsOptions(habits)
	fmt.Print("Enter command: ")
	_, err := fmt.Scan(&habitChoice)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	if habitChoice < 1 || habitChoice > len(habits) {
		fmt.Println("Returning to main menu...")
		return
	}

	var entryDate string
	fmt.Print("Enter date (YYYY-MM-DD) or press Enter for today: ")

	_, err = fmt.Scanln(&entryDate)
	if err != nil || entryDate == "" {
		entryDate = time.Now().Format("2006-01-02")
	}

	fmt.Println("Entry date:", entryDate)

	_, err = time.Parse("2006-01-02", entryDate)
	if err != nil {
		fmt.Println("Invalid date:", entryDate)
		return
	}

	var duration int
	fmt.Print("Enter duration in minutes (or Press Enter for 15 minutes): ")
	_, err = fmt.Scanln(&duration)
	if err != nil || duration == 0 {
		duration = 15
	}

	var requestBody = EntryRequest{
		HabitId:   habits[habitChoice-1].Id,
		EntryDate: entryDate,
		Duration:  duration,
	}

	PostHabitEntry(requestBody)
}

func printFetchHabits(habits []Habit) {
	clear()
	clear()
	for _, habit := range habits {
		fmt.Printf("%s (%s)\n", habit.Name, habit.Id)
		fmt.Printf("Description: %s\n", habit.Description)
		fmt.Printf("Streak: %d\n", habit.Streak)
		fmt.Println("Entries (last 30 days):")
		PrintStreak(habit)
		PrintTotalDuration(habit)
		fmt.Println("")
	}
}

func PrintStreak(habit Habit) {
	n := 30
	var habitStreak []HabitStreak
	for i := 0; i < n; i++ {
		date := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		habitStreak = append(habitStreak, HabitStreak{
			HabitEntryDate: date,
			IsActive:       false,
		})
	}
	entries := habit.Entries
	for i := 0; i < n; i++ {
		for _, entry := range entries {
			if CompareDates(habitStreak[i].HabitEntryDate, entry.EntryDate) {
				habitStreak[i].IsActive = true
			}
		}
	}
	// fmt.Println(habitStreak)
	slices.Reverse(habitStreak)
	for i := 0; i < len(habitStreak); i++ {
		if habitStreak[i].IsActive {
			fmt.Print("■ ")
		} else {
			fmt.Print("□ ")
		}
	}
	fmt.Println("")
}

func PrintTotalDuration(habit Habit) {
	totalDuration := 0.0
	for _, entry := range habit.Entries {
		totalDuration += float64(entry.Duration)
	}
	fmt.Printf("Total duration: %.2f hours\n", totalDuration/60.0)
}

func CompareDates(date1 string, date2 string) bool {
	layout := "2006-01-02"
	t1, err := time.Parse(layout, date1)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	layout2 := "2006-01-02T15:04:05"
	t2, err := time.Parse(layout2, date2)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	return t1.Equal(t2)
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
