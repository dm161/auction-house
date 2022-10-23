package listing

import (
	"strconv"
	"strings"
)

type Parser interface {
	Parse(line string) (Listing, error)
}

type parser struct{}

func New() Parser {
	return parser{}
}

// Parse returns a Listing from a raw string
func (p parser) Parse(line string) (Listing, error) {
	parts := strings.Split(line, "|")
	if len(parts) != listingFields {
		return Listing{}, errInvalidListingFormat
	}
	if parts[action] != listingAction {
		return Listing{}, errInvalidListingFormat
	}

	intTimestamp, err := strconv.Atoi(parts[timestamp])
	if err != nil {
		return Listing{}, errInvalidListingTimestamp
	}
	intCloseTime, err := strconv.Atoi(parts[closeTime])
	if err != nil {
		return Listing{}, errInvalidListingCloseTime
	}
	intReservePrice, err := strconv.Atoi(strings.ReplaceAll(parts[reservePrice], ".", ""))
	if err != nil {
		return Listing{}, errInvalidListingReservePrice
	}

	return Listing{
		Timestamp:    uint64(intTimestamp),
		ReservePrice: uint64(intReservePrice),
		CloseTime:    uint64(intCloseTime),
		UserID:       parts[userID],
		Item:         parts[item],
	}, nil
}
