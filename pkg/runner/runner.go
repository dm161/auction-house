package runner

import (
	"errors"

	"diego.pizza/auction-house/pkg/bid"
	"diego.pizza/auction-house/pkg/engine"
	"diego.pizza/auction-house/pkg/hearthbeat"
	"diego.pizza/auction-house/pkg/listing"
)

var (
	errUnableToParseMessage = errors.New("unable to parse message")
)

// StringIter is an iterator type used to return string(s) to the [runner's] caller
type StringIter func() (string, bool)

// EmptyIter used to return an empty iterator
var EmptyIter StringIter = func() (string, bool) { return "", false }

// IterFromStringSlice used to return an iterator from a []string response
var IterFromStringSlice = func(res []string) StringIter {
	var idx int = 0
	if len(res) == 0 {
		return EmptyIter
	}

	return func() (string, bool) {
		prevIdx := idx
		idx++
		if prevIdx >= len(res) {
			return "", false
		}

		return res[prevIdx], (prevIdx < len(res))
	}
}

// Runner holds reference of all message parsing and
// engine which is responsible for calulating auction results
type Runner struct {
	bidParser        bid.Parser
	listingParser    listing.Parser
	hearthbeatParser hearthbeat.Parser
	engine           engine.Engine
}

// New returns a new Runner instance
func New(
	bidParser bid.Parser,
	listingParser listing.Parser,
	hearthbeatParser hearthbeat.Parser,
	engine engine.Engine,
) Runner {
	return Runner{bidParser, listingParser, hearthbeatParser, engine}
}

// Run tries to parse incoming an message and when it encounters an hearthbeat
// attempts to return an engine response as an iterator
func (r Runner) Run(msg string) (StringIter, error) {
	listing, err := r.listingParser.Parse(msg)
	if err != nil {
		bid, err := r.bidParser.Parse(msg)
		if err != nil {
			hearthbeat, err := r.hearthbeatParser.Parse(msg)
			if err != nil {
				return EmptyIter, errUnableToParseMessage
			}

			return IterFromStringSlice(r.engine.Hearthbeat(hearthbeat)), nil
		}

		return EmptyIter, r.engine.Bid(bid)
	}

	return EmptyIter, r.engine.List(listing)
}
