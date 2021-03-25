package pointsratingsystem

import (
	"math"
	"strconv"
)

// Card holds vital information on a player in the Points Rating System
type Card struct {
	name       string
	rating     int     // 14 points to a stone; 0 == 1d
	volatility float64 // (0, 84]
	lastPeriod int     // last period played in
}

// NewCard creates a new playercard with default volatility
func NewCard(name string, xrating float64, ratingOrigin string, lastPeriod int) *Card {
	c := new(Card)

	c.name = name
	c.rating = Xrtor(xrating, ratingOrigin)
	c.lastPeriod = lastPeriod
	c.volatility = 7.0

	return c
}

// UpdateCard processes the W-L record for a given rating period
// for a given card
// volatility gives a triangular numbers approach
func (c Card) UpdateCard(wins int, losses int, draws int, missedPeriods int) {
	c.volatility = math.Min(c.volatility+(float64(missedPeriods)*7), 84) // Update volatility for inactivity

	// Update c.rating based on Bayesian reasoning and the volatility
	probUnderated := Underated(wins, losses, draws)
	c.rating = int(math.Round(float64(c.rating) + (probUnderated * c.volatility)))

	c.volatility = c.volatility * 0.5 // enforce downward volatility bias

	// c.volatility = c.volatility + ((inverse from above - 2) * 2) // Max 84
	if probUnderated <= 0.5 {
		c.volatility = math.Min(c.volatility+(((1.0/probUnderated)-2.0)*2.0), 84.0)
	} else {
		c.volatility = math.Min(c.volatility+(((1.0/(1.0-probUnderated))-2.0)*2.0), 84.0)
	}
}

// Underated returns the probability that the player is underated
// given the results of the last rating period with naive assumptions
func Underated(wins int, losses int, draws int) float64 {
	if draws == 0 {
		/*
			A == P(being underated)
			B == WLD record

			P(A | B) = (P(A) * P(B | A)) / (P(B))
			P(A | B) = (0.5 * P(B | 2/3 winning coin)) / ((P(B | 2/3 winning coin) * 0.5) + (P(B | 1/3 winning coin) * 0.5))
			P(A | B) = P(B | 2/3 winning coin) / (P(B | 2/3 winning coin) + P(B | 1/3 winning coin))
		*/

		underated, overrated := biasCoin(wins, wins+losses)

		return underated / (underated + overrated)
	}

	return (Underated(wins+draws, losses, 0) +
		Underated(wins, losses+draws, 0)/2)
}

func biasCoin(hitTarget int, coins int) (float64, float64) {
	spreads := []int{}
	for i := 0; i < coins; i++ {
		spreads = append(spreads, 0)
	}

	head := 0
	carry := false

	zeroes := coins
	winningBiasHits, losingBiasHits, trials := 0, 0, 0

	if (coins - zeroes) >= hitTarget {
		winningBiasHits++
	}
	if zeroes >= hitTarget {
		losingBiasHits++
	}
	trials++

	for {
		switch spreads[head] {
		case 0:
			spreads[head]++
			zeroes--
			if (coins - zeroes) >= hitTarget {
				winningBiasHits++
			}
			if zeroes >= hitTarget {
				losingBiasHits++
			}
			trials++
			carry = false
		case 1:
			spreads[head]++
			if (coins - zeroes) >= hitTarget {
				winningBiasHits++
			}
			if zeroes >= hitTarget {
				losingBiasHits++
			}
			trials++
			carry = false
		case 2:
			spreads[head] = 0
			zeroes++
			head++
			carry = true
		}

		if head >= coins {
			break
		}

		if carry == false {
			head = 0
		}
	}

	return float64(winningBiasHits) / float64(trials),
		float64(losingBiasHits) / float64(trials)
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
		r = int(math.Round((math.Log(xr/525.0)*23.15)-30.0) * 14.0)
	case "OGS-12":
		// 12 points per stone in conversion
		// [0, 13] == 1d
		r = int(math.Round((math.Log(xr/525.0)*23.15)-30.0) * 12.0)
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
		rf64 = math.Ceil(rf64 / -14.0)
		kdstr = strconv.Itoa(int(rf64)) + "k"
	} else { // dan
		rf64 = math.Floor(rf64/14.0) + 1.0
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
