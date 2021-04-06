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

	// Rating Updating
	updateRating(0, 4, 0, 0, 0, roster)  // Samraku
	updateRating(1, 0, 0, 0, 0, roster)  // illusory_deceit
	updateRating(2, 1, 1, 0, 0, roster)  // KoBa
	updateRating(3, 0, 2, 0, 0, roster)  // teapoweredrobot
	updateRating(4, 1, 2, 0, 0, roster)  // He Who Walks in Shadows
	updateRating(5, 0, 0, 0, 0, roster)  // Kaworu Nagisa
	updateRating(6, 0, 1, 0, 0, roster)  // pdg137
	updateRating(7, 5, 2, 0, 0, roster)  // wurfmau3
	updateRating(8, 0, 1, 0, 0, roster)  // DashaTabasco
	updateRating(9, 2, 1, 0, 0, roster)  // riiia
	updateRating(10, 0, 2, 0, 0, roster) // vyzhael
	updateRating(11, 0, 0, 0, 0, roster) // LittlePebble
	updateRating(12, 0, 1, 0, 0, roster) // sbk96

	// TODO: Print Player Info (rating/record)
	// fmt.Printf("\nafter rating updates:\n") // DEBUG
	// fmt.Println(roster.ListCards())

	// 2021-2 (2)
	newEntry(0, 4, roster, entries) // Samraku
	newEntry(3, 2, roster, entries) // teapoweredrobot
	newEntry(6, 1, roster, entries) // pdg137
	newEntry(7, 5, roster, entries) // wurfmau3

	// *pairings = randPair(*roster, *entries) // Manually Pairing this round

	fmt.Println()
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
	randPlayers := rand.Perm(len(entries)) // player list

	// pairing variables
	nextPlayer := 0                                                                // next player to pair; card index
	oppFit := make([][]int, len(entries))                                          // [[card, fit], [card, fit]]
	fmt.Printf("\nentries address: %p\noppFit address: %p\n\n", &entries, &oppFit) // DEBUG
	copy(oppFit, entries)
	fmt.Printf("\nafter copy()\nentries address: %p\noppFit address: %p\n\n", &entries, &oppFit) // DEBUG
	fmt.Println("oppFit: ", oppFit)                                                              // DEBUG
	bestOpp := 0
	bestOppFitRating := 0
	// reverseColors := false
	pairings := [][]int{}           // [[cardB, cardW], [cardB, cardW]]
	fmt.Println("CREATED pairings") // DEBUG

	// for loops
	genningPairings := true // pairing players

	// stats
	playersLeft := 0 // count of players to pair
	gamesLeft := 0   // count of games to pair
	lastPlayer := "" // name of last player to pair

	fmt.Println("randPair entries: ", entries) // DEBUG
	fmt.Println()                              // DEBUG

	for genningPairings {
		fmt.Println("entered genningPairings for loop") // DEBUG

		// Select player to pair
		for p := range randPlayers {
			fmt.Println("ENTER select player loop")
			if entries[randPlayers[p]][1] > 0 { // gamesLeft !> 0
				// pair player p
				nextPlayer = entries[randPlayers[p]][0]

				fmt.Printf("nextPlayer (after set): %v\n", nextPlayer) // DEBUG

				break
			}
		}

		fmt.Println("Finished finding nextPlayer") // DEBUG

		// Rate opponents
		for o := range oppFit {
			fmt.Println("ENTER oppRatingLoop")                         // DEBUG
			fmt.Printf("entries: %v\noppFit: %v\n\n", entries, oppFit) // DEBUG

			oppFit[o][1] = 0 // re-initialize

			if oppFit[o][0] == entries[nextPlayer][0] { // Prevent Self-Matching
				fmt.Printf("before adding 0x1000 to oppFitValue:\nentries: %v\noppFit: %v\n\n", entries, oppFit) // DEBUG
				oppFit[o][1] += 0x1000
				fmt.Printf("after adding 0x1000 to oppFitValue:\nentries: %v\noppFit: %v\n\n", entries, oppFit) // DEBUG
			}

			if roster.GetRatingGap(nextPlayer, oppFit[o][0]) > 120 { // <= 8 stones 8 points
				fmt.Println("ENTER GetRatingGap if")
				fmt.Printf("before adding 0x1000 to oppFitValue:\nentries: %v\noppFit: %v\n\n", entries, oppFit) // DEBUG
				oppFit[o][1] += 0x1000
				fmt.Printf("after adding 0x1000 to oppFitValue:\nentries: %v\noppFit: %v\n\n", entries, oppFit) // DEBUG
			}

			// Duplicate Matches
			for p := range pairings {
				fmt.Println("ENTER duplicateMatchLoop")                    // DEBUG
				fmt.Printf("entries: %v\noppFit: %v\n\n", entries, oppFit) // DEBUG

				if pairings[p][0] == nextPlayer &&
					pairings[p][1] == oppFit[o][0] { // 0x800 for duplicate match
					oppFit[o][1] += 0x800
				}

				if pairings[p][0] == oppFit[o][0] &&
					pairings[p][1] == nextPlayer { // 0x400 for reversed duplicate match
					oppFit[o][1] += 0x400
				}
			}
		}

		// Select best opponent (lowest oppFit)
		bestOppFitRating = 0x4000
		for o := range randPlayers {
			if oppFit[o][1] <= bestOppFitRating {
				bestOpp = oppFit[o][0]
				bestOppFitRating = oppFit[o][1]
			}
		}
		fmt.Printf("\nbestOpp: %v\nFitRating: %v\n\n", bestOpp, bestOppFitRating) // DEBUG

		// Pair
		// TODO; check if reversing B-W gives better pairing
		pairings = append(pairings, []int{nextPlayer, bestOpp})

		// Update entries[p, o][1]
		for e := range entries {
			if entries[e][0] == nextPlayer ||
				entries[e][0] == bestOpp {
				entries[e][1]--
			}
		}

		fmt.Println("CREATED new pairing")         // DEBUG
		fmt.Println("randPair entries: ", entries) // DEBUG
		fmt.Println()                              // DEBUG

		// reinitialize variables
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

	return pairings
}

func updateRating(card int, wins int, losses int, draws int, missedPeriods int, roster *prs.Roster) {
	roster.UpdateCardFromRoster(card, wins, losses, draws, missedPeriods)
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
