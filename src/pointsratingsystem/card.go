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
func NewCard(name string, xrating float64, ratingOrigin string, lastPeriod int) *Card {
	c := new(Card)

	c.name = name
	c.rating = Xrtor(xrating, ratingOrigin)
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
// xro (External Rating Origin) defines the behavior of the conversion
func Xrtor(xr float64, xro string) int {
	var r int

	switch xro {
	case "R":
		r = int(xr)
	case "DR":
		r = Drtor(int(xr))
	case "OGS":
		// 14 points per stone in conversion
		// [0, 13] == 1d
		r = int(math.Round((math.Log(xr/525)*23.15)-30) * 14)
	case "OGS-12":
		// 12 points per stone in conversion
		// [0, 13] == 1d
		r = int(math.Round((math.Log(xr/525)*23.15)-30) * 12)
	default:
		r = -126 // Should be 9k
	}

	return r
}

// Drtor converts a display rating to prs rating
func Drtor(dr int) int {
	return dr - 420
}

// Rtodr converts a prs rating to a display rating
func Rtodr(r int) int {
	return r + 420
}

// Rtokd converts a prs rating to one in terms of kyu-dan
// stones are 14 points apart
// 0 = shodan
func Rtokd(r int) string {
	var kdstr string
	rf64 := float64(r)

	if r < 0 { // kyu
		rf64 = math.Ceil(rf64 / -14)
		kdstr = strconv.Itoa(int(rf64)) + "k"
	} else { // dan
		rf64 = math.Floor(rf64/14) + 1
		kdstr = strconv.Itoa(int(rf64)) + "d"
	}

	return kdstr
}

// RatingGap returns the positive gap between two ratings
func RatingGap(r1 int, r2 int) int {
	return int(math.Abs(float64(r1 - r2)))
}

// DisplayRank returns prs rating in terms of kyu-dan
// stones are 14 points apart
// 0 = shodan
func (c Card) DisplayRank() string {
	return Rtokd(c.rating)
}
