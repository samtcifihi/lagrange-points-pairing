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

	// currentPeriod := 0 // const

	roundGenMacro(roster, &entries, &pairings)
	// debugMacro(roster, &entries, &pairings)
}

func roundGenMacro(roster *prs.Roster, entries *[][]int, pairings *[][]int) {
	// 2021-1 (1)
	currentPeriod := 0
	newPlayer("Samraku", 1771, "OGS", currentPeriod, roster)                 // 0
	newPlayer("illusory_deceit", 1278, "OGS", currentPeriod, roster)         // 1
	newPlayer("KoBa", 1808, "OGS", currentPeriod, roster)                    // 2 https://online-go.com/user/view/85719
	newPlayer("teapoweredrobot", 1488, "OGS", currentPeriod, roster)         // 3
	newPlayer("He Who Walks in Shadows", 1775, "OGS", currentPeriod, roster) // 4
	newPlayer("Kaworu Nagisa", 2368, "OGS", currentPeriod, roster)           // 5
	newPlayer("pdg137", 1391, "OGS", currentPeriod, roster)                  // 6
	newPlayer("wurfmau3", 2279, "OGS", currentPeriod, roster)                // 7
	newPlayer("DashaTabasco", 1126, "OGS", currentPeriod, roster)            // 8
	newPlayer("riiia", 1610, "OGS", currentPeriod, roster)                   // 9
	newPlayer("vyzhael", 1337, "OGS", currentPeriod, roster)                 // 10
	newPlayer("LittlePebble", 785, "OGS", currentPeriod, roster)             // 11
	newPlayer("sbk96", 1341, "OGS", currentPeriod, roster)                   // 12

	// TODO: Rating Updating

	// TODO: Print Player Info (rating/record)

	// 2021-2 (2)
	newEntry(0, 4, roster, entries) // Samraku
	newEntry(3, 2, roster, entries) // teapoweredrobot

	fmt.Println()
	*pairings = randPair(*roster, *entries)
	printPairings(roster, *pairings)
}

func debugMacro(roster *prs.Roster, entries *[][]int, pairings *[][]int) {
	/*
		// 2021-1 (1)
		currentPeriod := 0
		newEntry("Samraku", 1771, "OGS", 4, currentPeriod, roster, entries)                 // 0
		newEntry("illusory_deceit", 1278, "OGS", 1, currentPeriod, roster, entries)         // 1
		newEntry("KoBa", 1808, "OGS", 3, currentPeriod, roster, entries)                    // 2 https://online-go.com/user/view/85719
		newEntry("teapoweredrobot", 1488, "OGS", 2, currentPeriod, roster, entries)         // 3
		newEntry("He Who Walks in Shadows", 1775, "OGS", 3, currentPeriod, roster, entries) // 4
		newEntry("Kaworu Nagisa", 2368, "OGS", 1, currentPeriod, roster, entries)           // 5
		newEntry("pdg137", 1391, "OGS", 1, currentPeriod, roster, entries)                  // 6
		newEntry("wurfmau3", 2279, "OGS", 7, currentPeriod, roster, entries)                // 7
		newEntry("DashaTabasco", 1126, "OGS", 3, currentPeriod, roster, entries)            // 8
		newEntry("riiia", 1610, "OGS", 3, currentPeriod, roster, entries)                   // 9
		newEntry("vyzhael", 1337, "OGS", 2, currentPeriod, roster, entries)                 // \u218a
		newEntry("LittlePebble", 785, "OGS", 1, currentPeriod, roster, entries)             // \u218b
		newEntry("sbk96", 1341, "OGS", 1, currentPeriod, roster, entries)                   // 10
	*/
}

func newPlayer(username string, xr float64, xro string, currentPeriod int, roster *prs.Roster) {
	roster.AddCard(username, xr, xro, currentPeriod)
}

func newEntry(i int, gamesRegged int, roster *prs.Roster, entries *[][]int) {
	*entries = append(*entries, []int{i, gamesRegged})
}

// pairs random players together until only one player has games remaining
// and returns a pairings list
// Restrictions:
// * rating gap <= 120 (8 stones 8 points; 1/3 board)
// * avoid duplicate matchups
//   * strong prohibition
//   * weak prohibition
func randPair(roster prs.Roster, entries [][]int) [][]int {
	fmt.Println("ENTERING randPair") // DEBUG

	// initialize seed and pairings slice
	rand.Seed(time.Now().UnixNano())
	pairings := [][]int{}           // [[cardB, cardW], [cardB, cardW]]
	fmt.Println("CREATED pairings") // DEBUG

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
							for _, x := range pairings {
								if ((x[0] == entries[v][0] &&
									x[1] == entries[w][0]) == false) &&
									((x[0] == entries[w][0] &&
										x[1] == entries[v][0]) == false) {
									// pair
									pairings = append(pairings, []int{randPlayers[v], randPlayers[w]})
									entries[v][1]--
									entries[w][1]--
									findingMatch = false
								} else {
									break // continue finding a valid match
								}
							}
						case 1: // Weak Prohibition
							for _, x := range pairings {
								if (x[0] == entries[v][0] &&
									x[1] == entries[w][0]) == false {
									// pair
									pairings = append(pairings, []int{randPlayers[v], randPlayers[w]})
									entries[v][1]--
									entries[w][1]--
									findingMatch = false
								} else {
									break // continue finding a valid match
								}
							}
						case 0: // No Prohibition
							// pair
							pairings = append(pairings, []int{randPlayers[0], randPlayers[1]})
							entries[randPlayers[0]][1]--
							entries[randPlayers[1]][1]--
							findingMatch = false
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
		randPlayers = rand.Perm(len(entries))

		// check if all gamesRegged <= 0
		for _, v := range entries {
			if v[1] > 0 {
				playersLeft++
				gamesLeft += v[1]
				lastPlayer = roster.GetName(v[0])
			}
		}
		if playersLeft <= 1 {
			fmt.Printf("\"%v\" is missing %v games\n", lastPlayer, gamesLeft)
			genningPairings = false
		}

		playersLeft, gamesLeft = 0, 0
	} // pairings genned

	fmt.Printf("pairings: %v\n\n", pairings)
	return pairings
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
