package main

func main() {
	// slice of entrants
	entrants := []entrant{}

	// List of Entrants
	addEntrant("test", 1283, 3, &entrants) // 10k
}

func addEntrant(username string, rating float64, gameRegCount uint8, entrants *[]entrant) {
	e := newEntrant(username, rating, gameRegCount)
	*entrants = append(*entrants, *e)
}

// pass entrant.internalRank for both parameters
func rkomi(bRank int, wRank int) int {
	return 7 + bRank - wRank
}
