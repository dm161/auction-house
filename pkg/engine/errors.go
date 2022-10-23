package engine

import (
	"errors"
)

var (
	errListingAlreadyExists = errors.New("listing already exists for this item")
	errItemNotForSale       = errors.New("item not listed for sale")
)
