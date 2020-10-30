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
	s1a := draw(s1)
	s2a := draw(s2)

	m := t.Minute()
	var m1 int
	m2 := m % 10
	if m >= 10 {
		m1 = m / 10
	}
	m1a := draw(m1)
	m2a := draw(m2)

	h := t.Hour()
	var h1 int
	h2 := h % 10
	if h >= 10 {
		h1 = h / 10
	}
	h1a := draw(h1)
	h2a := draw(h2)

	semicln := draw(58)

	fmt.Print("\033[H\033[2J") // Clear the console

	for i := 0; i < len(s1a); i++ {
		fmt.Print(h1a[i], h2a[i], semicln[i], m1a[i], m2a[i], semicln[i], s1a[i], s2a[i], "\n")
	}
}

func main() {
	doEvery(250*time.Millisecond, clock)
}
