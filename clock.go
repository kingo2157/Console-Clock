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

func clock(timr time.Time) {
	t := time.Now()
	s := t.Second()
	var s1 int
	s2 := s % 10
	if s >= 10 {
		s1 = s / 10
	}

	fmt.Print("\033[H\033[2J") // Clear the console
	// fmt.Print(draw(s1), draw(s2))
	s1a := draw(s1)
	s2a := draw(s2)
	sa := append(s1a, s2a...)
	// fmt.Print(s1a[0], s2a[0])
	fmt.Print(sa)
}

func main() {
	doEvery(100*time.Millisecond, clock) // Change milisecond time to flag, possibly
}
