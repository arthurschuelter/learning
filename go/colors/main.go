package main

import "fmt"

type Colors int

const (
	Reset Colors = iota
	RedText
	GreenText
	YellowText
	BlueText
	BoldText
)

func GetColors() map[Colors]string {
	return map[Colors]string{
		Reset:      "\033[0m",
		RedText:    "\033[31m",
		GreenText:  "\033[32m",
		YellowText: "\033[33m",
		BlueText:   "\033[34m",
		BoldText:   "\033[1m",
	}
}

func main() {
	fmt.Println(Red("This is red"))
	fmt.Println(Red(Bold("This is bold red")))
	fmt.Println(Green("This is green"))
	fmt.Println(Green(Bold("This is bold green")))
	fmt.Println(Yellow("This is yellow"))
	fmt.Println(Yellow(Bold("This is bold yellow")))
	fmt.Println(Blue("This is blue"))
	fmt.Println(Blue(Bold("This is bold blue")))
	fmt.Println("This is plain text")
}

func Red(s string) string {
	colors := GetColors()
	return colors[RedText] + s + colors[Reset]
}

func Green(s string) string {
	colors := GetColors()
	return colors[GreenText] + s + colors[Reset]
}

func Yellow(s string) string {
	colors := GetColors()
	return colors[YellowText] + s + colors[Reset]
}

func Blue(s string) string {
	colors := GetColors()
	return colors[BlueText] + s + colors[Reset]
}

func Bold(s string) string {
	colors := GetColors()
	return colors[BoldText] + s + colors[Reset]
}
