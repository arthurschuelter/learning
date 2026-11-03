package main

import "fmt"

type Computer interface {
	InsertLightning()
}

type Mac struct {
}

func (m *Mac) InsertLightning() {
	fmt.Println("Lightning was plugged into Mac!")
}

type Windows struct {
}

func (w *Windows) InsertUSB() {
	fmt.Println("USB was plugged into Windows!")
}

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) InsertLightning() {
	fmt.Println("Adapter converting Lightning -> USB...")
	w.windowsMachine.InsertUSB()
}
