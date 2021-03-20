package pointsratingsystem

// Roster holds a collection of player cards in a given prs system
type Roster []*Card

// NewRoster creates a new Roster to hold player cards
func NewRoster() *Roster {
	return new(Roster)
}

// AddCard adds a new player card to the Roster
func (r Roster) AddCard(name string, xrating float64, lastPeriod int) {
	c := NewCard(name, xrating, lastPeriod)

	r = append(r, c)
}

// ListCards returns data for all cards for output
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
