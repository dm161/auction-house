package bid

import (
	"strconv"
	"strings"
)

type Parser interface {
	Parse(line string) (Bid, error)
}

type parser struct{}

func New() Parser {
	return parser{}
}

// Parse returns a Bid from a raw string
func (p parser) Parse(line string) (Bid, error) {
	parts := strings.Split(line, "|")
	if len(parts) != bidFields {
		return Bid{}, errInvalidBidFormat
	}
	if parts[action] != bidAction {
		return Bid{}, errInvalidBidFormat
	}

	intTimestamp, err := strconv.Atoi(parts[timestamp])
	if err != nil {
		return Bid{}, errInvalidBidTimestamp
	}
	intBidAmount, err := strconv.Atoi(strings.ReplaceAll(parts[bidAmount], ".", ""))
	if err != nil {
		return Bid{}, errInvalidBidAmount
	}

	return Bid{
		Timestamp: uint64(intTimestamp),
		Amount:    uint64(intBidAmount),
		UserID:    parts[userID],
		Item:      parts[item],
	}, nil
}
