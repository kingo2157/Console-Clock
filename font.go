package main

func draw(num int) []string {
	if num == 0 {
		return []string{"   XXXXXX  ",
			"  XX    XX ",
			" XX    XXXX",
			" XX   XX XX",
			" XX  XX  XX",
			" XX XX   XX",
			" XXXX    XX",
			"  XX    XX ",
			"   XXXXXX  "}
	} else if num == 1 {
		return []string{"     XX    ",
			"    XXX    ",
			"  XXXXX    ",
			"     XX    ",
			"     XX    ",
			"     XX    ",
			"     XX    ",
			"     XX    ",
			"  XXXXXXXX "}
	} else if num == 2 {
		return []string{"  XXXXXXXX ",
			" XX      XX",
			"         XX",
			"       XXX ",
			"     XXX   ",
			"   XXX     ",
			"  XX       ",
			" XX        ",
			" XXXXXXXXXX"}
	} else if num == 3 {
		return []string{" XXXXXXXXXX",
			"         XX",
			"       XXX ",
			"    XXXX   ",
			"       XXX ",
			"         XX",
			"         XX",
			" XX      XX",
			"  XXXXXXXX "}
	} else if num == 4 {
		return []string{"      XXX  ",
			"     XXXX  ",
			"    XX XX  ",
			"   XX  XX  ",
			"  XX   XX  ",
			" XXXXXXXXXX",
			"       XX  ",
			"       XX  ",
			"       XX  "}
	} else if num == 5 {
		return []string{" XXXXXXXXXX",
			" XX        ",
			" XX        ",
			" XXXXXXXX  ",
			"        XX ",
			"         XX",
			"         XX",
			" XX     XX ",
			"  XXXXXXX  "}
	} else if num == 6 {
		return []string{"     XXXX  ",
			"   XXX     ",
			"  XX       ",
			" XX        ",
			" XXXXXXXXX ",
			" XX      XX",
			" XX      XX",
			" XX      XX",
			"  XXXXXXXX "}
	} else if num == 7 {
		return []string{" XXXXXXXXXX",
			"         XX",
			"        XX ",
			"       XX  ",
			"      XX   ",
			"     XX    ",
			"    XX     ",
			"   XX      ",
			"   XX      "}
	} else if num == 8 {
		return []string{"  XXXXXXXX ",
			" XX      XX",
			" XX      XX",
			"  XX    XX ",
			"   XXXXXX  ",
			"  XX    XX ",
			" XX      XX",
			" XX      XX",
			"  XXXXXXXX "}
	} else if num == 9 {
		return []string{"  XXXXXXXX ",
			" XX      XX",
			" XX      XX",
			" XX      XX",
			"  XXXXXXXXX",
			"         XX",
			"        XX ",
			"      XXX  ",
			"   XXXX    "}
	} else if num == 58 {
		return []string{"   ",
			" XX",
			" XX",
			"   ",
			"   ",
			"   ",
			" XX",
			" XX",
			"   "}
	} else {
		return []string{}
	}
}
