package hearthbeat

import (
	"errors"
)

var (
	errInvalidHearthbeatFormat = errors.New("invalid hearthbeat format")
	errInvalidHearthbeatValue  = errors.New("invalid hearthbeat value")
)
