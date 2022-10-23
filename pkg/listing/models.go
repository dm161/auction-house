package listing

const (
	timestamp = iota
	userID
	action
	item
	reservePrice
	closeTime
)

const (
	listingFields = 6
	listingAction = "SELL"
)

// Listing represents an item listed for sale by a user
type Listing struct {
	Timestamp    uint64
	ReservePrice uint64
	CloseTime    uint64
	UserID       string
	Item         string
}
