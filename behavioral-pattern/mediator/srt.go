package main

import "fmt"

type srt struct {
	code     string
	mediator mediator
}

func (s *srt) arrive() {
	if !s.mediator.canArrive(s) {
		fmt.Println("SRT code", s.code, "wait for arrival")
		return
	}
	fmt.Println("SRT code", s.code, "arrive")
}

func (s *srt) depart() {
	fmt.Println("SRT code", s.code, "leaving")
	s.mediator.notifyAboutDeparture()
}

func (s *srt) permitArrival() {
	fmt.Println("SRT code", s.code, "get permit")
	s.arrive()
}
