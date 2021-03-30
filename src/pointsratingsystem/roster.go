package pointsratingsystem

// Roster holds a collection of player cards in a given prs system
type Roster []*Card

// NewRoster creates a new Roster to hold player cards
func NewRoster() *Roster {
	return new(Roster)
}

// AddCard adds a new player card to the Roster
func (r *Roster) AddCard(name string, xr float64, xro string, lastPeriod int) {
	// Test for name already existing in the roster.

	c := NewCard(name, xr, xro, lastPeriod)

	*r = append(*r, c)
}

// RetrieveCard returns the first card in Roster r with the passed name
func (r Roster) RetrieveCard(name string) *Card {
	var index int

	for i, c := range r {
		if c.name == name {
			index = i
			break
		}
	}

	if index < len(r) {
		return r[index]
	}

	r.AddCard(name, 0.0, "default", 0)

	return r[0]
}

// RetrieveLast returns the most recently added card
func (r Roster) RetrieveLast() *Card {
	return r[len(r)-1]
}

// RetrieveLastIndex returns the index of the most recently added card
func (r Roster) RetrieveLastIndex() int {
	return len(r) - 1
}

// GetName returns the name for the card at list[i]
func (r Roster) GetName(i int) string {
	return r[i].name
}

// GetRating returns the rating for the card at list[i]
func (r Roster) GetRating(i int) int {
	return r[i].rating
}

// GetRatingGap returns the positive gap between the ratings of two cards
func (r Roster) GetRatingGap(c1 int, c2 int) int {
	return RatingGap(r[c1].rating, r[c2].rating)
}

// DisplayRank returns the rank of card i in kyu-dan format for printing
func (r Roster) DisplayRank(i int) string {
	return r[i].DisplayRank()
}

// Inject injects the specified number of (positive or negative)
// rating points into every card in Roster r
func (r Roster) Inject(i int) {
	for _, v := range r {
		v.rating = v.rating + i
	}
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
