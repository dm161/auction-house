package listing

import (
	"errors"
)

var (
	errInvalidListingFormat       = errors.New("invalid listing format")
	errInvalidListingTimestamp    = errors.New("invalid listing timestamp")
	errInvalidListingCloseTime    = errors.New("invalid listing close time")
	errInvalidListingReservePrice = errors.New("invalid listing reserve price")
)
