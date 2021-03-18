package main

import "fmt"

type tv struct {
	isOn bool
}

func (t *tv) on() {
	t.isOn = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isOn = false
	fmt.Println("Turning tv off")
}
