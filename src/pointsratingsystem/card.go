package pointsratingsystem

import (
	"math"
	"strconv"
)

// Card holds vital information on a player in the Points Rating System
type Card struct {
	name       string
	rating     int // 14 points to a stone; 0 == 1d
	volatility int // [1, 7]
	lastPeriod int // last period played in
}

// NewCard creates a new playercard with default volatility
func NewCard(name string, xrating float64, lastPeriod int) *Card {
	c := new(Card)

	c.name = name
	c.rating = Xrtor(xrating)
	c.lastPeriod = lastPeriod
	c.volatility = 7

	return c
}

// Xrtor converts an external rating to prs rating
func Xrtor(xrating float64) int {
	// 12 points per stone in conversion
	// [0, 13] == 1d
	return int(math.Round((math.Log(xrating/525)*23.15)-30) * 12)
}

// DisplayRanking returns prs rating in terms of kyu-dan
// stones are 14 points apart
// 0 = shodan
func (c Card) DisplayRanking() string {
	kdf := float64(c.rating)
	var kda string

	if c.rating < 0 { // kyu
		kdf = math.Ceil(kdf / -14)
		kda = strconv.Itoa(int(kdf)) + "k"
	} else { // dan
		kdf = math.Floor(kdf/14) + 1
		kda = strconv.Itoa(int(kdf)) + "d"
	}

	return kda
}
