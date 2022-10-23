package bid

import (
	"errors"
)

var (
	errInvalidBidFormat    = errors.New("invalid bid format")
	errInvalidBidTimestamp = errors.New("invalid bid timestamp")
	errInvalidBidAmount    = errors.New("invalid bid amount")
)
