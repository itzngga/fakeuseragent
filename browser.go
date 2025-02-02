package useragent

import (
	"fmt"
	"math/rand"
	"time"
)

type Browser int

const (
	Chrome Browser = iota
	Opera
	Firefox
	Safari
	Edge
)

func (b Browser) String() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	switch b {
	case Chrome:
		return fmt.Sprintf("%s Chrome/%s", APPLE_WEB_KIT[0], CHROME[rng.Intn(len(CHROME))])
	case Opera:
		return fmt.Sprintf("%s %s", APPLE_WEB_KIT[0], OPERA[rng.Intn(len(OPERA))])
	case Firefox:
		return fmt.Sprintf("Gecko/20100101 %s", FIREFOX[rng.Intn(len(FIREFOX))])
	case Safari:
		return fmt.Sprintf("%s %s", APPLE_WEB_KIT[0], SAFARI[rng.Intn(len(SAFARI))])
	case Edge:
		return fmt.Sprintf("%s %s", APPLE_WEB_KIT[0], EDGE[rng.Intn(len(EDGE))])
	default:
		return ""
	}
}

func RandomBrowser(rng *rand.Rand) Browser {
	n := rng.Intn(120)
	switch n {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29:
		return Chrome
	case 30, 31, 32, 33, 34, 35, 36, 37, 38, 39:
		return Opera
	case 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69:
		return Firefox
	case 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89:
		return Safari
	default:
		return Edge
	}
}
