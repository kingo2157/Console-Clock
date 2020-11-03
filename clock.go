package main

import (
	"fmt"
	"time"
)

func doEvery(td time.Duration, t func(time.Time)) {
	for x := range time.Tick(td) {
		t(x)
	}
}

func clock(t time.Time) {
	sec := t.Second()
	var sec1 int
	sec2 := sec % 10
	if sec >= 10 {
		sec1 = sec / 10
	}
	sec1a := draw(sec1)
	sec2a := draw(sec2)

	min := t.Minute()
	var min1 int
	min2 := min % 10
	if min >= 10 {
		min1 = min / 10
	}
	min1a := draw(min1)
	min2a := draw(min2)

	hour := t.Hour()
	var hour1 int
	hour2 := hour % 10
	if hour >= 10 {
		hour1 = hour / 10
	}
	hour1a := draw(hour1)
	hour2a := draw(hour2)

	semicln := draw(58)

	fmt.Print("\033[H\033[2J") // Clear the console

	for i := 0; i < len(sec1a); i++ {
		fmt.Print(hour1a[i], hour2a[i], semicln[i], min1a[i], min2a[i], semicln[i], sec1a[i], sec2a[i], "\n")
	}
}

func main() {
	doEvery(250*time.Millisecond, clock)
}
