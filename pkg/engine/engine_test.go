package engine

import (
	"testing"

	"diego.pizza/auction-house/pkg/hearthbeat"
	"diego.pizza/auction-house/pkg/listing"
)

func TestExpiredListings(t *testing.T) {
	for name, tc := range map[string]struct {
		listings        map[string]*listing.Listing
		hearthbeat      hearthbeat.Hearthbeat
		expiredListings []*listing.Listing
	}{
		"none is expired": {
			listings: map[string]*listing.Listing{
				"1": {CloseTime: uint64(10)},
				"2": {CloseTime: uint64(12)},
				"3": {CloseTime: uint64(15)},
			},
			hearthbeat:      hearthbeat.Hearthbeat{Timestamp: 5},
			expiredListings: []*listing.Listing{},
		},
		"one is expired": {
			listings: map[string]*listing.Listing{
				"1": {CloseTime: uint64(10)},
				"2": {CloseTime: uint64(12)},
				"3": {CloseTime: uint64(15)},
			},
			hearthbeat:      hearthbeat.Hearthbeat{Timestamp: 11},
			expiredListings: []*listing.Listing{{CloseTime: uint64(10)}},
		},
		"all are expired": {
			listings: map[string]*listing.Listing{
				"1": {CloseTime: uint64(10)},
				"2": {CloseTime: uint64(12)},
				"3": {CloseTime: uint64(15)},
			},
			hearthbeat:      hearthbeat.Hearthbeat{Timestamp: 50},
			expiredListings: []*listing.Listing{{CloseTime: uint64(10)}, {CloseTime: uint64(12)}, {CloseTime: uint64(15)}},
		},
	} {
		t.Run(name, func(t *testing.T) {
			e := engine{listings: tc.listings}
			expired := e.expiredListings(tc.hearthbeat)
			if len(expired) != len(tc.expiredListings) {
				t.Errorf("expected len(`%v`) got len(`%v`)", len(tc.expiredListings), len(expired))
			}
		})
	}
}

func TestList(t *testing.T) {
	for name, tc := range map[string]struct {
		listings            map[string]*listing.Listing
		listingBids         map[string]*listingBids
		expectedListingBids map[string]*listingBids
		listing             listing.Listing
		expectedListings    []*listing.Listing
		err                 error
	}{
		"err listing already exists": {
			listings: map[string]*listing.Listing{
				"1": {Item: "1"},
			},
			listingBids: map[string]*listingBids{
				"1": newListingBids(),
			},
			listing: listing.Listing{Item: "1"},
			expectedListings: []*listing.Listing{
				{Item: "1"},
			},
			err: errListingAlreadyExists,
		},
		"listing added successfully": {
			listings: map[string]*listing.Listing{
				"1": {Item: "1"},
			},
			listingBids: map[string]*listingBids{
				"1": newListingBids(),
			},
			listing: listing.Listing{Item: "2"},
			expectedListings: []*listing.Listing{
				{Item: "1"},
				{Item: "2"},
			},
			expectedListingBids: map[string]*listingBids{
				"1": newListingBids(),
				"2": newListingBids(),
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			e := engine{listings: tc.listings, listingsBids: tc.listingBids}
			err := e.List(tc.listing)
			if err != tc.err {
				t.Errorf("expected `%v` got `%v`", tc.err, err)
			}
			if len(e.listings) != len(tc.expectedListings) {
				t.Errorf("expected len(`%v`) got len(`%v`)", len(tc.expectedListings), len(e.listings))
			}
			if tc.expectedListingBids != nil && e.listingsBids[tc.listing.Item] == nil {
				t.Error("listingBids heaps should be initialised when adding a new listing")
			}
		})
	}
}
