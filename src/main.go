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

func newEntry(username string, xr float64, xro string, gamesRegged int, currentPeriod int, roster *prs.Roster, entries *[][]int) {
	roster.AddCard(username, xr, xro, currentPeriod)

	*entries = append(*entries, []int{roster.Len(), gamesRegged})
}

func returningEntry(i int, gamesRegged int, roster *prs.Roster, entries *[][]int) {
	*entries = append(*entries, []int{i, gamesRegged})
}

// pairs random players together until only one player has games remaining
// Restrictions:
// * rating gap <= 120 (8 stones 8 points; 1/3 board)
// * avoid duplicate matchups
//   * strong prohibition
//   * weak prohibition
func randPair(roster prs.Roster, entries [][]int, pairings *[][]int) {
	// initialize seed and pairings slice
	rand.Seed(time.Now().UnixNano())
	pairings = nil

	genningPairings := true // pairing players
	findingMatch := true    // looking for a valid matchup
	dupProhib := 2          // 2: strong duplication avoidance; 1: weak duplication avoidance; 0: no duplication avoidance

	randPlayers := rand.Perm(len(entries)) // player list

	playersLeft := 0 // count of players to pair
	gamesLeft := 0   // count of games to pair
	lastPlayer := "" // name of last player to pair

	for genningPairings {
		for findingMatch {
			for _, v := range randPlayers {
				for _, w := range randPlayers {
					if findingMatch &&
						v != w && // No self-matching
						roster.GetRatingGap(entries[v][0], entries[w][0]) <= 120 && // <= 8 stones 8 points
						entries[v][1] > 0 && // No more games than signed up for
						entries[w][1] > 0 { // No more games than signed up for
						switch dupProhib {
						case 2: // Strong Prohibition
							for _, x := range *pairings {
								if ((x[0] == entries[v][0] &&
									x[1] == entries[w][0]) == false) &&
									((x[0] == entries[w][0] &&
										x[1] == entries[v][0]) == false) {
									// pair
									*pairings = append(*pairings, []int{randPlayers[v], randPlayers[w]})
									entries[v][1]--
									entries[w][1]--
									findingMatch == false
								} else {
									break // continue finding a valid match
								}
							}
						case 1: // Weak Prohibition
							for _, x := range *pairings {
								if (x[0] == entries[v][0] &&
									x[1] == entries[w][0]) == false {
									// pair
									*pairings = append(*pairings, []int{randPlayers[v], randPlayers[w]})
									entries[v][1]--
									entries[w][1]--
									findingMatch == false
								} else {
									break // continue finding a valid match
								}
							}
						case 0: // No Prohibition
							// pair
							*pairings = append(*pairings, []int{randPlayers[0], randPlayers[1]})
							entries[randPlayers[0]][1]--
							entries[randPlayers[1]][1]--
							findingMatch == false
						default:
							fmt.Println("Achievment Unlocked: -1 Restrictions on Duplicate Matchups!")
						}
					}
				}
			}
			if findingMatch {
				dupProhib--
			}
		}

		// reinitialize variables
		findingMatch = true
		pl = rand.Perm(l)

		// check if all gamesRegged <= 0
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
	} // pairings genned

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
