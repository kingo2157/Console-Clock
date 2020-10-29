package main

func draw(num int /*, pass int*/) []string {
	if num == 0 {
		return []string{"   XXXXXX  \n",
			"  XX    XX \n",
			" XX    XXXX\n",
			" XX   XX XX\n",
			" XX  XX  XX\n",
			" XX XX   XX\n",
			" XXXX    XX\n",
			"  XX    XX \n",
			"   XXXXXX  \n"}
	} else if num == 1 {
		return []string{"     XX    \n",
			"    XXX    \n",
			"  XXXXX    \n",
			"     XX    \n",
			"     XX    \n",
			"     XX    \n",
			"     XX    \n",
			"     XX    \n",
			"  XXXXXXXX \n"}
	} else if num == 2 {
		return []string{"  XXXXXXXX \n",
			" XX      XX\n",
			"         XX\n",
			"       XXX \n",
			"     XXX   \n",
			"   XXX     \n",
			"  XX       \n",
			" XX        \n",
			" XXXXXXXXXX\n"}
	} else if num == 3 {
		return []string{" XXXXXXXXXX\n",
			"         XX\n",
			"       XXX \n",
			"    XXXX   \n",
			"       XXX \n",
			"         XX\n",
			"         XX\n",
			" XX      XX\n",
			"  XXXXXXXX \n"}
	} else if num == 4 {
		return []string{"      XXX  \n",
			"     XXXX  \n",
			"    XX XX  \n",
			"   XX  XX  \n",
			"  XX   XX  \n",
			" XXXXXXXXXX\n",
			"       XX  \n",
			"       XX  \n",
			"       XX  \n"}
	} else if num == 5 {
		return []string{" XXXXXXXXXX\n",
			" XX        \n",
			" XX        \n",
			" XXXXXXXX  \n",
			"        XX \n",
			"         XX\n",
			"         XX\n",
			" XX     XX \n",
			"  XXXXXXX  \n"}
	} else if num == 6 {
		return []string{"     XXXX  \n",
			"   XXX     \n",
			"  XX       \n",
			" XX        \n",
			" XXXXXXXXX \n",
			" XX      XX\n",
			" XX      XX\n",
			" XX      XX\n",
			"  XXXXXXXX \n"}
	} else if num == 7 {
		return []string{" XXXXXXXXXX\n",
			"         XX\n",
			"        XX \n",
			"       XX  \n",
			"      XX   \n",
			"     XX    \n",
			"    XX     \n",
			"   XX      \n",
			"   XX      \n"}
	} else if num == 8 {
		return []string{"  XXXXXXXX \n",
			" XX      XX\n",
			" XX      XX\n",
			"  XX    XX \n",
			"   XXXXXX  \n",
			"  XX    XX \n",
			" XX      XX\n",
			" XX      XX\n",
			"  XXXXXXXX \n"}
	} else if num == 9 {
		return []string{"  XXXXXXXX \n",
			" XX      XX\n",
			" XX      XX\n",
			" XX      XX\n",
			"  XXXXXXXXX\n",
			"         XX\n",
			"        XX \n",
			"      XXX  \n",
			"   XXXX    \n"}
	} else if num == 58 {
		return []string{"   \n",
			" XX\n",
			" XX\n",
			"   \n",
			"   \n",
			"   \n",
			" XX\n",
			" XX\n"}
	} else {
		return []string{}
	}
}
