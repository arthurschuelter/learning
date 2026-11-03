package models

import (
	"fmt"
)

type Repo struct {
	Id           int64  `json:"id"`
	FullName     string `json:"full_name"`
	LanguagesURL string `json:"languages_url"`
	Stars        int    `json:"stargazers_count"`
}

type Ranking struct {
	Language string
	Total    int
	Ratio    float32
}

type User struct {
	Name      string `json:"name"`
	Login     string `json:"login"`
	Location  string `json:"location"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
}

func (ranking *Ranking) PrintRanking(i int) {
	fmt.Printf("  (%2d) %s: %.2f kB (%.2f%%)\n", i+1, ranking.Language, float32(ranking.Total)/1024, ranking.Ratio*100)
}

func (user *User) PrintData() {
	fmt.Printf("%s (%s) @ %s\n", user.Name, user.Login, user.Location)
	// fmt.Printf("(%s)\n", user.Login)
	fmt.Printf("Followers: %d | Following: %d\n", user.Followers, user.Following)
}
