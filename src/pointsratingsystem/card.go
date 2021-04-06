package pointsratingsystem

import (
	"fmt"
	"math"
	"strconv"
)

// Card holds vital information on a player in the Points Rating System
type Card struct {
	name       string
	rating     int     // 14 points to a stone; 0 == 1d
	volatility float64 // (2, 84]
	lastPeriod int     // last period played in
}

// NewCard creates a new playercard with default volatility
func NewCard(name string, xrating float64, ratingOrigin string, lastPeriod int) *Card {
	c := new(Card)

	c.name = name
	c.rating = Xrtor(xrating, ratingOrigin)
	c.lastPeriod = lastPeriod
	c.volatility = 84.0

	return c
}

// GetName returns c.name
func (c Card) GetName() string {
	return c.name
}

// UpdateCard processes the W-L record for a given rating period
// for a given card
// volatility gives a triangular numbers approach
func (c *Card) UpdateCard(wins int, losses int, draws int, missedPeriods int) {

	fmt.Printf("%v: %v-%v-%v\n", c.name, wins, losses, draws)
	fmt.Printf("prior rating: %v (%v) (vol = %v)\n", Rtodr(c.rating), Rtokd(c.rating), strconv.FormatFloat(c.volatility, 'f', 2, 64))

	// If the player had any results this period
	if !(wins == 0 &&
		losses == 0 &&
		draws == 0) {

		c.volatility = math.Min(c.volatility+(float64(missedPeriods)*3.5), 84.0) // Update volatility for inactivity

		// Update c.rating based on Bayesian reasoning and the volatility
		probUnderated := Underated(wins, losses, draws)

		c.rating = int(math.Round(float64(c.rating) + (probUnderated * c.volatility) - ((1.0 - probUnderated) * c.volatility))) // prior rating + points for underated chance + points for overrated chance
		c.volatility = (c.volatility * 0.5) + 1.5                                                                               // enforce downward volatility bias and make this step make volatility approach 3.0, where rating will change by 1 point on a (1, 0) record

		// c.volatility = c.volatility + ((inverse from above - 2) * 1) // Max 84
		if probUnderated <= 0.5 {
			c.volatility = math.Min(c.volatility+(((1.0/probUnderated)-2.0)*1.0), 84.0) // 1.0 is a free parameter; may need slight decreasing
		} else {
			c.volatility = math.Min(c.volatility+(((1.0/(1.0-probUnderated))-2.0)*1.0), 84.0) // 1.0 is a free parameter; may need slight decreasing
		}
	}

	fmt.Printf("updated rating: %v (%v) (vol = %v)\n\n", Rtodr(c.rating), Rtokd(c.rating), strconv.FormatFloat(c.volatility, 'f', 2, 64))
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

		probGivenUnderated, probGivenOverrated := resultExtremity(wins, wins+losses)

		return probGivenUnderated / (probGivenUnderated + probGivenOverrated)
	}

	return (Underated(wins+draws, losses, 0) +
		Underated(wins, losses+draws, 0)) / 2
}

// resultExtremity calculates how likely a result at least as extreme as
// provided is, for both a 2/3 chance of hitting (winning), and a 1/3 chance of hitting.
func resultExtremity(hitTarget int, sackSize int) (float64, float64) {
	likelyHitProb, unlikelyHitProb := 0.0, 0.0

	if hitTarget == sackSize-hitTarget {
		likelyHitProb, unlikelyHitProb = 0.5, 0.5 // Not true, but gives the correct result in this case.
	} else if hitTarget < sackSize-hitTarget { // hitTarget lower half (including middle) of sackSize
		for n := 0; n <= hitTarget; n++ {
			likelyHitProb += float64(combos(n, sackSize)) * math.Pow((2.0/3.0), float64(n)) * math.Pow((1.0/3.0), float64(sackSize-n))
			unlikelyHitProb += float64(combos(n, sackSize)) * math.Pow((1.0/3.0), float64(n)) * math.Pow((2.0/3.0), float64(sackSize-n))
		}
	} else { // hitTarget upper half (excluding middle) of sackSize
		for n := hitTarget; n <= sackSize; n++ {
			likelyHitProb += float64(combos(n, sackSize)) * math.Pow((2.0/3.0), float64(n)) * math.Pow((1.0/3.0), float64(sackSize-n))
			unlikelyHitProb += float64(combos(n, sackSize)) * math.Pow((1.0/3.0), float64(n)) * math.Pow((2.0/3.0), float64(sackSize-n))
		}
	}

	return likelyHitProb, unlikelyHitProb
}

func combos(hits int, space int) int {
	if (hits <= 0) || (hits == space) { // base case
		return 1
	}

	// combos when the first index of space is not a hit,
	// plus combos when the first index of space is a hit.
	return combos(hits, (space-1)) + combos((hits-1), (space-1))
}

func biasCoin(hitTarget int, coins int) (float64, float64) {
	// winHitTarget, lossHitTarget := wins, coins-wins
	// hitTarget = wins

	spreads := []int{}
	for i := 0; i < coins; i++ {
		spreads = append(spreads, 0)
	}

	head := 0
	carry := false

	zeroes := coins
	winningBiasHits, losingBiasHits, trials := 0, 0, 0

	if hitTarget < coins-hitTarget {
		if (coins - zeroes) <= hitTarget {
			winningBiasHits++
		}
	} else {
		if (coins - zeroes) >= hitTarget {
			winningBiasHits++
		}
	}
	if hitTarget < coins-hitTarget {
		if zeroes <= hitTarget {
			losingBiasHits++
		}
	} else {
		if zeroes >= hitTarget {
			losingBiasHits++
		}
	}
	trials++

	fmt.Println("winningBiasHits", winningBiasHits, // DEBUG
		"; losingBiasHits", losingBiasHits,
		"; trials", trials)

	for {
		switch spreads[head] {
		case 0:
			spreads[head]++
			zeroes--
			if hitTarget < coins-hitTarget {
				if (coins - zeroes) <= hitTarget {
					winningBiasHits++
				}
			} else {
				if (coins - zeroes) >= hitTarget {
					winningBiasHits++
				}
			}
			if hitTarget < coins-hitTarget {
				if zeroes <= hitTarget {
					losingBiasHits++
				}
			} else {
				if zeroes >= hitTarget {
					losingBiasHits++
				}
			}
			trials++
			fmt.Println("spreads", spreads, // DEBUG
				"; winningBiasHits", winningBiasHits,
				"; losingBiasHits", losingBiasHits,
				"; trials", trials)

			carry = false
		case 1:
			spreads[head]++
			if hitTarget < coins-hitTarget {
				if (coins - zeroes) <= hitTarget {
					winningBiasHits++
				}
			} else {
				if (coins - zeroes) >= hitTarget {
					winningBiasHits++
				}
			}
			if hitTarget < coins-hitTarget {
				if zeroes <= hitTarget {
					losingBiasHits++
				}
			} else {
				if zeroes >= hitTarget {
					losingBiasHits++
				}
			}
			trials++
			fmt.Println("spreads", spreads, // DEBUG
				"; winningBiasHits", winningBiasHits,
				"; losingBiasHits", losingBiasHits,
				"; trials", trials)

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
