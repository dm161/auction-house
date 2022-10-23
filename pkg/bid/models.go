package bid

const (
	timestamp = iota
	userID
	action
	item
	bidAmount
)

const (
	bidFields = 5
	bidAction = "BID"
)

// Bid represents a Bid from a user for an item
type Bid struct {
	Timestamp uint64
	Amount    uint64
	UserID    string
	Item      string
}
