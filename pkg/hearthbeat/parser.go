package hearthbeat

import (
	"strconv"
	"strings"
)

type Parser interface {
	Parse(line string) (Hearthbeat, error)
}

type parser struct{}

func New() Parser {
	return parser{}
}

// Parse returns a Hearthbeat from a raw string
func (p parser) Parse(line string) (Hearthbeat, error) {
	parts := strings.Split(line, "|")
	if len(parts) != hearthbeatFields {
		return Hearthbeat{}, errInvalidHearthbeatFormat
	}

	intHearthbeat, err := strconv.Atoi(parts[hearthbeat])
	if err != nil {
		return Hearthbeat{}, errInvalidHearthbeatValue
	}

	return Hearthbeat{Timestamp: uint64(intHearthbeat)}, nil
}
