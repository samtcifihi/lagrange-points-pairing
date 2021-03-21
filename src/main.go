package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	prs "github.com/samtcifihi/lagrange-points-pairing/src/pointsratingsystem"
)

func main() {
	roster := prs.NewRoster()
	entries := [][]int{}  // [[card, gamereg], [card, gamereg]]
	pairings := [][]int{} // [[cardB, cardW], [cardB, cardW]]

	currentPeriod := 0 // const

	// List of Entrants
	newEntry("Samraku", 1771, "OGS", 4, currentPeriod, roster, &entries)                 // 0
	newEntry("illusory_deceit", 1278, "OGS", 1, currentPeriod, roster, &entries)         // 1
	newEntry("KoBa", 1808, "OGS", 3, currentPeriod, roster, &entries)                    // 2 https://online-go.com/user/view/85719
	newEntry("teapoweredrobot", 1488, "OGS", 2, currentPeriod, roster, &entries)         // 3
	newEntry("He Who Walks in Shadows", 1775, "OGS", 3, currentPeriod, roster, &entries) // 4
	newEntry("Kaworu Nagisa", 2368, "OGS", 1, currentPeriod, roster, &entries)           // 5
	newEntry("pdg137", 1391, "OGS", 1, currentPeriod, roster, &entries)                  // 6
	newEntry("wurfmau3", 2279, "OGS", 7, currentPeriod, roster, &entries)                // 7
	newEntry("DashaTabasco", 1126, "OGS", 3, currentPeriod, roster, &entries)            // 8
	newEntry("riiia", 1610, "OGS", 3, currentPeriod, roster, &entries)                   // 9
	newEntry("vyzhael", 1337, "OGS", 2, currentPeriod, roster, &entries)                 // \u218a
	newEntry("LittlePebble", 785, "OGS", 1, currentPeriod, roster, &entries)             // \u218b
	newEntry("sbk96", 1341, "OGS", 1, currentPeriod, roster, &entries)                   // 10

	// Pair Entrants
	fmt.Println()
	randPair(*roster, entries, &pairings)

	// List Matches
	printPairings(roster, pairings)
}

func newEntry(username string, xrating float64, ratingOrigin string, gameRegCount int, currentPeriod int, roster *prs.Roster, entries *[][]int) {
	roster.AddCard(username, xrating, ratingOrigin, currentPeriod)

	*entries = append(*entries, []int{roster.Len(), gameRegCount})
}

func returningEntry(i int, gameRegCount int, roster *prs.Roster, entries *[][]int) {
	*entries = append(*entries, []int{i, gameRegCount})
}

// pairs random players together until only one player has games remaining
func randPair(roster prs.Roster, entries [][]int, pairings *[][]int) {
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

func printPairings(roster *prs.Roster, pairings [][]int) {
	fmt.Println("B | W | komi | result")
	fmt.Println("- | - | ---- | ------")

	// loop through pairings
	for _, v := range pairings {
		fmt.Println(matchStr(v[0], v[1], roster))
	}
}

func matchStr(black int, white int, roster *prs.Roster) string {
	return roster.GetName(black) + " [" + roster.DisplayRank(black) + "] | " +
		roster.GetName(white) + " [" + roster.DisplayRank(white) + "] | " +
		strconv.Itoa(rkomi(roster.GetRating(black), roster.GetRating(white))) +
		" | [result](link)"
}

// pass entrant.internalRank for both parameters
func rkomi(bRank int, wRank int) int {
	return 7 + bRank - wRank
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
