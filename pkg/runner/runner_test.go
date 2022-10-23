package runner

import (
	"errors"
	"testing"

	"diego.pizza/auction-house/pkg/bid"
	"diego.pizza/auction-house/pkg/engine"
	"diego.pizza/auction-house/pkg/hearthbeat"
	"diego.pizza/auction-house/pkg/listing"
)

type bidParserSuccess struct{}

func (b bidParserSuccess) Parse(msg string) (bid.Bid, error) {
	return bid.Bid{}, nil
}

type bidParserFail struct{}

var errBidParser = errors.New("fix me")

func (p bidParserFail) Parse(msg string) (bid.Bid, error) {
	return bid.Bid{}, errBidParser
}

type listingParserSuccess struct{}

func (p listingParserSuccess) Parse(msg string) (listing.Listing, error) {
	return listing.Listing{}, nil
}

type listingParserFail struct{}

var errListingParser = errors.New("fix me")

func (p listingParserFail) Parse(msg string) (listing.Listing, error) {
	return listing.Listing{}, errListingParser
}

type hearthbeatParserSuccess struct{}

func (b hearthbeatParserSuccess) Parse(msg string) (hearthbeat.Hearthbeat, error) {
	return hearthbeat.Hearthbeat{}, nil
}

type hearthbeatParserFail struct{}

var errHearthbeatParser = errors.New("fix me")

func (b hearthbeatParserFail) Parse(msg string) (hearthbeat.Hearthbeat, error) {
	return hearthbeat.Hearthbeat{}, errHearthbeatParser
}

type engineMock struct {
	res []string
}

func (e engineMock) Bid(b bid.Bid) error {
	return nil
}
func (e engineMock) List(l listing.Listing) error {
	return nil
}
func (e engineMock) Hearthbeat(b hearthbeat.Hearthbeat) []string {
	return e.res
}

func TestRun(t *testing.T) {
	for name, tc := range map[string]struct {
		bidParser        bid.Parser
		listingParser    listing.Parser
		hearthbeatParser hearthbeat.Parser
		engine           engine.Engine
		in               string
		out              []string
		err              error
	}{
		"listing parser returns successfully": {
			listingParser: listingParserSuccess{},
			engine:        engineMock{},
			out:           []string{},
		},
		"bid parser returns successfully": {
			listingParser: listingParserFail{},
			bidParser:     bidParserSuccess{},
			engine:        engineMock{},
			out:           []string{},
		},
		"hearthbeat parser returns successfully": {
			listingParser:    listingParserFail{},
			bidParser:        bidParserFail{},
			hearthbeatParser: hearthbeatParserSuccess{},
			engine:           engineMock{[]string{"ok", "ok1", "ok2"}},
			out:              []string{"ok", "ok1", "ok2"},
		},
		"returns error on all invalid parsers": {
			listingParser:    listingParserFail{},
			bidParser:        bidParserFail{},
			hearthbeatParser: hearthbeatParserFail{},
			out:              []string{},
			err:              errUnableToParseMessage,
		},
	} {
		t.Run(name, func(t *testing.T) {
			r := New(tc.bidParser, tc.listingParser, tc.hearthbeatParser, tc.engine)
			iter, err := r.Run("foo")
			if err != tc.err {
				t.Errorf("got `%v` expected `%v`", err, tc.err)
			}
			var i int
			for line, next := iter(); next; line, next = iter() {
				if line != tc.out[i] {
					t.Errorf("got `%v` expected `%v`", line, tc.out[i])
				}
				i++
			}
			if i != len(tc.out) {
				t.Errorf("got `%v` expected `%v`", i, len(tc.out))
			}
		})
	}
}
