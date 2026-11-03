package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type Day struct {
	Date          string  `json:"date"`
	DateFormatted string  `json:"dateFormatted"`
	DayOfWeek     string  `json:"dayOfWeek"`
	IsToday       bool    `json:"isToday"`
	Movies        []Movie `json:"movies"`
}

type Movie struct {
	ID            string        `json:"id"`
	Title         string        `json:"title"`
	OriginalTitle string        `json:"originalTitle"`
	Duration      string        `json:"duration"`
	ContentRating string        `json:"contentRating"`
	Genres        []string      `json:"genres"`
	SessionTypes  []SessionType `json:"sessionTypes"`
	Images        []Image       `json:"images"`
}

type SessionType struct {
	Type     []string  `json:"type"`
	Sessions []Session `json:"sessions"`
}

type Session struct {
	ID    string  `json:"id"`
	Room  string  `json:"room"`
	Price float64 `json:"price"`
	Time  string  `json:"time"`
	Date  struct {
		LocalDate   string `json:"localDate"`
		DayOfWeek   string `json:"dayOfWeek"`
		DayAndMonth string `json:"dayAndMonth"`
		Hour        string `json:"hour"`
	} `json:"date"`
}

type Image struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

var domains = []string{
	"www.ingresso.com",
	"api-content.ingresso.com",
}

func ScrapeGNC(cinemas []Cinema, dates []string) {
	movieDays := []Day{}

	c := colly.NewCollector(
		colly.AllowedDomains(domains...),
	)

	c.OnResponse(func(r *colly.Response) {
		days, err := ParseDays(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		for _, day := range days {
			fmt.Print(FormatMovieOutput(day))
			movieDays = append(movieDays, day)
		}
	})

	fetchCinemas(c, cinemas, dates)
	c.Wait()

}

func ParseDays(data []byte) ([]Day, error) {
	var days []Day
	err := json.Unmarshal(data, &days)
	return days, err
}

func BuildURL(cinemaID string, date string) string {
	return fmt.Sprintf("https://api-content.ingresso.com/v0/sessions/city/16/theater/%s/partnership/home/groupBy/sessionType?date=%s",
		cinemaID, date)
}

func fetchCinemas(c *colly.Collector, cinemas []Cinema, dates []string) {
	for _, date := range dates {
		for _, cine := range cinemas {
			url := BuildURL(cine.ID, date)
			fmt.Printf("\n=== %s (%s) ===\n", cine.Title, date)
			err := c.Visit(url)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func FormatMovieOutput(day Day) string {
	var sb strings.Builder
	for _, movie := range day.Movies {
		sb.WriteString(fmt.Sprintf("\nFilme: %s (%s)\n", movie.Title, movie.OriginalTitle))
		sb.WriteString(fmt.Sprintf("Duração: %s min | Classificação: %s\n", movie.Duration, movie.ContentRating))
		sb.WriteString(fmt.Sprintf("Gêneros: [%s]\n", strings.Join(movie.Genres, ", ")))

		for _, sessionType := range movie.SessionTypes {
			sb.WriteString(fmt.Sprintf("  Tipo: [%v]\n", strings.Join(sessionType.Type, ", ")))
			for _, session := range sessionType.Sessions {
				sb.WriteString(fmt.Sprintf("    - %s | Sala: %s | R$ %.2f\n", session.Time, session.Room, session.Price))

			}
		}
	}

	return sb.String()
}
