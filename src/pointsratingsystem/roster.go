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
func (r Roster) AddCard(name string, xr float64, xro string, lastPeriod int) {
	// Test for name already existing in the roster.

	c := NewCard(name, xr, xro, lastPeriod)

	r.list = append(r.list, c)
}

// RetrieveCard returns the first card in Roster r with the passed name
func (r Roster) RetrieveCard(name string) *Card {
	var index int

	for i, c := range r.list {
		if c.name == name {
			index = i
			break
		}
	}

	return r.list[index]
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

// GetRatingGap returns the positive gap between the ratings of two cards
func (r Roster) GetRatingGap(c1 int, c2 int) int {
	return RatingGap(r.list[c1].rating, r.list[c2].rating)
}

// DisplayRank returns the rank of card i in kyu-dan format for printing
func (r Roster) DisplayRank(i int) string {
	return r.list[i].DisplayRank()
}

// Inject injects the specified number of (positive or negative)
// rating points into every card in Roster r
func (r Roster) Inject(i int) {
	for _, v := range r.list {
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
