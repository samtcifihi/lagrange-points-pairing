package main

import (
	"fmt"
	"strconv"
)

func main() {
	// slice of entrants
	entrants := []entrant{}

	// List of Entrants
	addEntrant("Samraku", 1763, 3, &entrants) // 1d
	addEntrant("riiia", 1600, 3, &entrants)         // 10k
	addEntrant("vyzhael", 1305, 3, &entrants)         // 10k
	addEntrant("anoek", 1347, 3, &entrants)         // 10k
	addEntrant("KyTb", 1429, 3, &entrants)         // 4
	addEntrant("Go_Michael", 993, 3, &entrants)         // 5

	// List of Matches
	fmt.Println(matchStr(entrants[0], entrants[5]))
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
	return black.username + " [" + black.displayRank + "] (B) vs. " +
		white.username + " [" + white.displayRank + "] (W; " +
		strconv.Itoa(rkomi(black.internalRank, white.internalRank)) + " komi)" +
		" ([result](link))"
}
