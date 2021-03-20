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
	addEntrant("Samraku", 1771, 4, &entrants)                 // 0
	addEntrant("illusory_deceit", 1278, 1, &entrants)         // 1
	addEntrant("KoBa", 1808, 3, &entrants)                    // 2 https://online-go.com/user/view/85719
	addEntrant("teapoweredrobot", 1488, 2, &entrants)         // 3
	addEntrant("He Who Walks in Shadows", 1775, 3, &entrants) // 4
	addEntrant("Kaworu Nagisa", 2368, 1, &entrants)           // 5
	addEntrant("pdg137", 1391, 1, &entrants)                  // 6
	addEntrant("wurfmau3", 2279, 7, &entrants)                // 7
	addEntrant("DashaTabasco", 1126, 3, &entrants)            // 8
	addEntrant("riiia", 1610, 3, &entrants)                   // 9
	addEntrant("vyzhael", 1337, 2, &entrants)                 // \u218a
	addEntrant("LittlePebble", 785, 1, &entrants)             // \u218b
	addEntrant("sbk96", 1341, 1, &entrants)                   // 10

	// Pair Entrants
	fmt.Println()
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

	running := true // pairing players
	finding := true // looking for a valid pairing
	l := len(entrants)
	pl := rand.Perm(l) // player list
	c := 0             // count of players to pair
	n := ""            // name of last player to pair
	g := 0             // count of games to pair

	for running {
		// check if strong prohibition of duplicate matchups works
		for i := range pl {
			for j := range pl {
				if finding &&
					i != j &&
					entrants[pl[i]].gameRegCount > 0 &&
					entrants[pl[j]].gameRegCount > 0 {
					if conVal(*pairings, []int{pl[i], pl[j]}) <= 0 &&
						conVal(*pairings, []int{pl[j], pl[i]}) <= 0 {

						// pair
						*pairings = append(*pairings, []int{pl[i], pl[j]})
						entrants[pl[i]].gameRegCount--
						entrants[pl[j]].gameRegCount--
						finding = false
					}
				}
			}
		}

		// check if weak prohibition of duplicate matchups works
		if finding {
			for i := range pl {
				for j := range pl {
					if finding &&
						i != j &&
						entrants[pl[i]].gameRegCount > 0 &&
						entrants[pl[j]].gameRegCount > 0 &&
						!(conVal(*pairings, []int{pl[i], pl[j]}) > 0) {
						// pair
						*pairings = append(*pairings, []int{pl[i], pl[j]})
						entrants[pl[i]].gameRegCount--
						entrants[pl[j]].gameRegCount--
						finding = false
					}
				}
			}
		}

		// pair 0 with 1
		if finding {
			*pairings = append(*pairings, []int{pl[0], pl[1]})
			entrants[pl[0]].gameRegCount--
			entrants[pl[1]].gameRegCount--
		}

		// reinitialize variables
		finding = true
		pl = rand.Perm(l)

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

// Checks if a slice contains a given value and returns the number of hits
func conVal(slice [][]int, val []int) int {
	hits := 0

	// temp := []int{0, 1}

	for i := 0; i < len(slice); i++ {

		if slice[i][0] == val[0] && slice[i][1] == val[1] {
			hits++
		}
	}

	return hits
}
