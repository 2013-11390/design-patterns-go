package main

func main() {
	stationManager := &stationManager{isPlatformFree: true}

	srt1 := &srt{code: "A1", mediator: stationManager}
	srt2 := &srt{code: "B1", mediator: stationManager}
	srt3 := &srt{code: "C1", mediator: stationManager}

	srt1.arrive()
	srt3.arrive()
	srt2.arrive()
	srt1.depart()
	srt3.depart()
}
