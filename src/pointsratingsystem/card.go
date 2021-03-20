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

// UpdateCard processes the W-L record for a given rating period
// for a given card
// volatility gives a triangular numbers approach
func (c Card) UpdateCard(wins int, losses int, draws int, period int) {
	// Update Volatility for periods elapsed
	c.volatility = c.volatility + (period - c.lastPeriod) - 1
	// put volatility in [1, 7]
	c.volatility = int(math.Min(math.Max(float64(c.volatility), 1), 7))

	// Update Volatility for W-L pairs (and draws)
	for draws > 0 {
		c.volatility--
		draws--
	}

	for wins > 0 && losses > 0 {
		c.volatility = int(math.Max(float64(c.volatility)-2, 1))
		wins--
		losses--
	}

	// Update rating and Volatility for remaining W/Ls
	for wins > 0 {
		c.rating = c.rating + c.volatility
		c.volatility = int(math.Max(float64(c.volatility)-1, 1))
	}

	for losses > 0 {
		c.rating = c.rating - c.volatility
		c.volatility = int(math.Max(float64(c.volatility)-1, 1))
	}

	c.lastPeriod = period
}

// Xrtor converts an external rating to prs rating
func Xrtor(xrating float64) int {
	// 12 points per stone in conversion
	// [0, 13] == 1d
	return int(math.Round((math.Log(xrating/525)*23.15)-30) * 12)
}

// DisplayRank returns prs rating in terms of kyu-dan
// stones are 14 points apart
// 0 = shodan
func (c Card) DisplayRank() string {
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
