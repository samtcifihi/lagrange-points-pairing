package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// slice of entrants
	entrants := []entrant{}

	// List of Entrants
	addEntrant("Samraku", 1763, 3, &entrants)                 // 0
	addEntrant("illusory_deceit", 1298, 1, &entrants)         // 1
	addEntrant("KoBa", 1784, 3, &entrants)                    // 2
	addEntrant("teapoweredrobot", 1494, 2, &entrants)         // 3
	addEntrant("He Who Walks in Shadows", 1755, 3, &entrants) // 4
	addEntrant("Kaworu Nagisa", 2368, 1, &entrants)           // 5
	addEntrant("pdg137", 1358, 1, &entrants)                  // 6

	// Pair Entrants
	pairings := [][]int{}
	randPair(entrants, &pairings)

	// List Matches
	printPairings(entrants, pairings)
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

// pairs random players together until only one player has games remaining
func randPair(entrants []entrant, pairings *[][]int) {

	// initialize seed
	rand.Seed(time.Now().UnixNano())

	running := true
	l := len(entrants)
	p := []int{}
	p = append(p, rand.Intn(l))
	p = append(p, rand.Intn(l))
	c := 0  // count of players to pair
	n := "" // name of last player to pair
	g := 0  // count of games to pair

	for running {
		// randomly select entrant in entrants with gameRegCount > 0
		// randomly select different entrant in entrants with gameRegCount > 0
		for entrants[p[0]].gameRegCount <= 0 {
			p[0] = rand.Intn(l)
		}
		for entrants[p[1]].gameRegCount <= 0 || p[0] == p[1] {
			p[1] = rand.Intn(l)
		}

		// add pairing to pairings
		*pairings = append(*pairings, []int{p[0], p[1]})

		// decrement both gameRegCount values
		entrants[p[0]].gameRegCount--
		entrants[p[1]].gameRegCount--

		// re-initialize p
		p[0] = rand.Intn(l)
		p[1] = rand.Intn(l)

		// check if all gameRegCount <= 0
		for _, v := range entrants {
			if v.gameRegCount > 0 {
				c++
				g += int(v.gameRegCount)
				n = v.username
			}
		}
		if c <= 1 {
			fmt.Printf("\"%v\" is missing %v games\n", n, g)
			running = false
		}
		c, g = 0, 0
	} // end running loop

	fmt.Printf("pairings: %v\n\n", *pairings)
}

func printPairings(entrants []entrant, pairings [][]int) {
	// loop through pairings
	for _, v := range pairings {
		fmt.Println(matchStr(entrants[v[0]], entrants[v[1]]))
	}
}
