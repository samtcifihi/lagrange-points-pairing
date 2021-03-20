package main

import (
	"math"
	"strconv"
)

type entrant struct {
	username     string
	rating       float64
	gameRegCount uint8
	internalRank int // 14 times larger than the 0-30 OGS scale
	displayRank  string
}

func newEntrant(username string, rating float64, gameRegCount uint8) *entrant {
	e := new(entrant)

	e.username = username
	e.rating = rating
	e.gameRegCount = gameRegCount

	// Calculate and set internalRank
	// This is the OGS rating-rank conversion formula multiplied
	// by 14 to gradate ranks down to rkomi points.
	e.internalRank = (int)(math.Round(math.Log(rating/525) * 23.15 * 14))

	// Generate and set displayRank
	ogsRank := (math.Log(rating/525) * 23.15) - 30
	if ogsRank < 0 { // kyu
		ogsRank = math.Ceil(ogsRank * -1)
		kyudan := (int)(ogsRank)
		e.displayRank = strconv.Itoa(kyudan) + "k"
	} else { // dan
		ogsRank = math.Floor(ogsRank) + 1
		kyudan := (int)(ogsRank)
		e.displayRank = strconv.Itoa(kyudan) + "d"
	}

	return e
}
