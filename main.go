package main

import (
	"flag"

	"example.com/set1"
)

func main() {
	challengePtr := flag.Int("c", 1, "choice of challenge to run")

	flag.Parse()

	switch *challengePtr {
	case 2:
		set1.Challenge2()
		break
	case 3:
		set1.Challenge3()
		break
	case 4:
		set1.Challenge4()
		break
	case 5:
		set1.Challenge5()
		break
	case 6:
		set1.Challenge6()
		break
	case 7:
		set1.Challenge7()
		break
	case 8:
		set1.Challenge8()
		break
	default:
		// case 1
		set1.Challenge1()
	}
}
