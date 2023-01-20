package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Clock struct {
	hour int
	min  int
	sec  int
}

var gen_clock Clock

func main() {
	fmt.Print("Inserire un timer (hh-mm-ss): ")
	fmt.Scan(&gen_clock.hour, &gen_clock.min, &gen_clock.sec)
	countdown()
}

func cancella() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func countdown() {
	for {
		if gen_clock.hour == 0 && gen_clock.min == 0 && gen_clock.sec == 0 {
			fmt.Println("Kabooom!")
			break
		} else if gen_clock.sec == 0 {
			cancella()
			fmt.Println(gen_clock)
			updateMin()
			gen_clock.sec = 59
			time.Sleep(time.Duration(1) * time.Second)
		} else if gen_clock.sec > 0 && gen_clock.sec <= 59 {
			cancella()
			fmt.Println(gen_clock)
			gen_clock.sec -= 1
			time.Sleep(time.Duration(1) * time.Second)
		} else {
			fmt.Println("Uh-oh: qualcosa è andato storto")
			break
		}
	}
}

func updateMin() {
	if gen_clock.sec == 0 {
		gen_clock.min -= 1
	}
}

func updateHour() {
	if gen_clock.min == 0 {
		gen_clock.hour -= 1
	}
}
