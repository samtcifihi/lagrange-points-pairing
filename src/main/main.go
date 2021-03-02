package main

import (
	"fmt"
	"strconv"
)

func main() {
	// slice of entrants
	entrants := []entrant{}

	// List of Entrants
	addEntrant("test", 1283, 3, &entrants)         // 10k
	addEntrant("testChinitsu", 1994, 3, &entrants) // 1d

	// List of Matches
	fmt.Println(matchStr(entrants[0], entrants[1]))
}

func addEntrant(username string, rating float64, gameRegCount uint8, entrants *[]entrant) {
	e := newEntrant(username, rating, gameRegCount)
	*entrants = append(*entrants, *e)
}

// pass entrant.internalRank for both parameters
func rkomi(bRank int, wRank int) int {
	return 7 + bRank - wRank
}

func matchStr(black entrant, white entrant) string {
	return black.username + " [" + black.displayRank + "] vs. " +
		white.username + " [" + white.displayRank + "] (" +
		strconv.Itoa(rkomi(black.internalRank, white.internalRank)) + " komi)"
}
