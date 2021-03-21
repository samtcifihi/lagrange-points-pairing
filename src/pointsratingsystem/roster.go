package pointsratingsystem

// Roster holds a collection of player cards in a given prs system
type Roster struct {
	list []*Card
}

// NewRoster creates a new Roster to hold player cards
func NewRoster() *Roster {
	return new(Roster)
}

// AddCard adds a new player card to the Roster
func (r Roster) AddCard(name string, xrating float64, ratingOrigin string, lastPeriod int) {
	c := NewCard(name, xrating, ratingOrigin, lastPeriod)

	r.list = append(r.list, c)
}

// RetrieveLast returns the most recently added card
func (r Roster) RetrieveLast() *Card {
	return r.list[len(r.list)]
}

// Len returns the index of the most recently added card
func (r Roster) Len() int {
	return len(r.list)
}

// GetName returns the name for the card at list[i]
func (r Roster) GetName(i int) string {
	return r.list[i].name
}

// GetRating returns the rating for the card at list[i]
func (r Roster) GetRating(i int) int {
	return r.list[i].rating
}

// DisplayRank returns the rank of card i in kyu-dan format for printing
func (r Roster) DisplayRank(i int) string {
	return r.list[i].DisplayRank()
}

// ListCards returns data for all cards for output
// Currently null
func (r Roster) ListCards(i ...int) string {
	var out string

	// if len(i) == 0 {
	// // List all
	// for j := range i {
	// out = out + r[]
	// }
	// } else {
	// // List r[i] for all i
	// }

	return out
}
