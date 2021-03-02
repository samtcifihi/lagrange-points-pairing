package main

func main() {
	// slice of entrants
	entrants := []entrant{}
	// pass to addEntrant()
	addEntrant("test", 1283, 3, &entrants) // 10k
}

func addEntrant(username string, rating float64, gameRegCount uint8, entrants *[]entrant) {
	e := newEntrant(username, rating, gameRegCount)
	*entrants = append(*entrants, *e)
}
