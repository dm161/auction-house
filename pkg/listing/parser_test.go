package listing

import (
	"testing"
)

func TestParse(t *testing.T) {
	for name, tc := range map[string]struct {
		inStr string
		out   Listing
		err   error
	}{
		"returns error on invalid number of fields": {
			inStr: "1|2",
			out:   Listing{},
			err:   errInvalidListingFormat,
		},
		"returns error on invalid listing action": {
			inStr: "15|8|BID|tv_1|250.00|20",
			out:   Listing{},
			err:   errInvalidListingFormat,
		},
		"returns error on invalid timestamp": {
			inStr: "foobar|8|SELL|tv_1|250.00|20",
			out:   Listing{},
			err:   errInvalidListingTimestamp,
		},
		"returns error on invalid listing reserve price": {
			inStr: "15|8|SELL|tv_1|foobar|20",
			out:   Listing{},
			err:   errInvalidListingReservePrice,
		},
		"returns valid bid": {
			inStr: "15|8|SELL|tv_1|250.00|20",
			out: Listing{
				Timestamp:    uint64(15),
				UserID:       "8",
				Item:         "tv_1",
				ReservePrice: uint64(25000),
				CloseTime:    uint64(20),
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			listing, err := New().Parse(tc.inStr)
			if err != tc.err {
				t.Errorf("got `%v` expected `%v`", err, tc.err)
			}
			if listing != tc.out {
				t.Errorf("got `%v` expected `%v`", listing, tc.out)
			}
		})
	}
}
