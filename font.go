package main

func draw(num int) []string {
	switch num {
	case 0:
		return []string{"   XXXXXX  ",
			"  XX    XX ",
			" XX    XXXX",
			" XX   XX XX",
			" XX  XX  XX",
			" XX XX   XX",
			" XXXX    XX",
			"  XX    XX ",
			"   XXXXXX  "}
	case 1:
		return []string{"     XX    ",
			"    XXX    ",
			"  XXXXX    ",
			"     XX    ",
			"     XX    ",
			"     XX    ",
			"     XX    ",
			"     XX    ",
			"  XXXXXXXX "}
	case 2:
		return []string{"  XXXXXXXX ",
			" XX      XX",
			"         XX",
			"       XXX ",
			"     XXX   ",
			"   XXX     ",
			"  XX       ",
			" XX        ",
			" XXXXXXXXXX"}
	case 3:
		return []string{" XXXXXXXXXX",
			"         XX",
			"       XXX ",
			"    XXXX   ",
			"       XXX ",
			"         XX",
			"         XX",
			" XX      XX",
			"  XXXXXXXX "}
	case 4:
		return []string{"      XXX  ",
			"     XXXX  ",
			"    XX XX  ",
			"   XX  XX  ",
			"  XX   XX  ",
			" XXXXXXXXXX",
			"       XX  ",
			"       XX  ",
			"       XX  "}
	case 5:
		return []string{" XXXXXXXXXX",
			" XX        ",
			" XX        ",
			" XXXXXXXX  ",
			"        XX ",
			"         XX",
			"         XX",
			" XX     XX ",
			"  XXXXXXX  "}
	case 6:
		return []string{"     XXXX  ",
			"   XXX     ",
			"  XX       ",
			" XX        ",
			" XXXXXXXXX ",
			" XX      XX",
			" XX      XX",
			" XX      XX",
			"  XXXXXXXX "}
	case 7:
		return []string{" XXXXXXXXXX",
			"         XX",
			"        XX ",
			"       XX  ",
			"      XX   ",
			"     XX    ",
			"    XX     ",
			"   XX      ",
			"   XX      "}
	case 8:
		return []string{"  XXXXXXXX ",
			" XX      XX",
			" XX      XX",
			"  XX    XX ",
			"   XXXXXX  ",
			"  XX    XX ",
			" XX      XX",
			" XX      XX",
			"  XXXXXXXX "}
	case 9:
		return []string{"  XXXXXXXX ",
			" XX      XX",
			" XX      XX",
			" XX      XX",
			"  XXXXXXXXX",
			"         XX",
			"        XX ",
			"      XXX  ",
			"   XXXX    "}
	case 58:
		return []string{"   ",
			" XX",
			" XX",
			"   ",
			"   ",
			"   ",
			" XX",
			" XX",
			"   "}
	default:
		return []string{}
	}
}
