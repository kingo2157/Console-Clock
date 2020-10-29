package main

import (
	"fmt"
	"strings"
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
	s1a := draw(s1)
	s2a := draw(s2)
	r := "0"
	for i := 0; i < len(s1a); i++ {
		// Replace all linebreaks in s1a[i] with s2a[i], all
		r = strings.NewReplacer("\n", s2a[i]).Replace(s1a[i])
		fmt.Print(r)
	}
}

func main() {
	doEvery(100*time.Millisecond, clock) // Change milisecond time to flag, possibly
}
