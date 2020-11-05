package main

import (
	"fmt"
	"time"
)

func main() {
	doEvery(250*time.Millisecond, clock)
}

func doEvery(td time.Duration, t func(time.Time)) {
	for x := range time.Tick(td) {
		t(x)
	}
}

func clock(t time.Time) {
	sec := t.Second()
	sec1, sec2 := 0, sec%10
	if sec >= 10 {
		sec1 = sec / 10
	}
	sec1a, sec2a := draw(sec1), draw(sec2)

	min := t.Minute()
	min1, min2 := 0, min%10
	if min >= 10 {
		min1 = min / 10
	}
	min1a, min2a := draw(min1), draw(min2)

	hour := t.Hour()
	hour1, hour2 := 0, hour%10
	if hour >= 10 {
		hour1 = hour / 10
	}
	hour1a, hour2a := draw(hour1), draw(hour2)

	semicln := draw(58)

	fmt.Print("\033[H\033[2J") // Clear the console

	for i := 0; i < len(sec1a); i++ {
		// Print Clock
		fmt.Print(hour1a[i], hour2a[i], semicln[i], min1a[i], min2a[i], semicln[i], sec1a[i], sec2a[i], "\n")
	}
}
