package bid

import (
	"testing"
)

func TestParse(t *testing.T) {
	for name, tc := range map[string]struct {
		inStr string
		out   Bid
		err   error
	}{
		"returns error on invalid number of fields": {
			inStr: "1|2",
			out:   Bid{},
			err:   errInvalidBidFormat,
		},
		"returns error on invalid bid action": {
			inStr: "12|8|SELL|toaster_1|7.50",
			out:   Bid{},
			err:   errInvalidBidFormat,
		},
		"returns error on invalid timestamp": {
			inStr: "foobar|8|BID|toaster_1|7.50",
			out:   Bid{},
			err:   errInvalidBidTimestamp,
		},
		"returns error on invalid bid amount": {
			inStr: "12|8|BID|toaster_1|foobar",
			out:   Bid{},
			err:   errInvalidBidAmount,
		},
		"returns valid bid": {
			inStr: "12|8|BID|toaster_1|7.50",
			out: Bid{
				Timestamp: uint64(12),
				UserID:    "8",
				Item:      "toaster_1",
				Amount:    uint64(750),
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			bid, err := New().Parse(tc.inStr)
			if err != tc.err {
				t.Errorf("got `%v` expected `%v`", err, tc.err)
			}
			if bid != tc.out {
				t.Errorf("got `%v` expected `%v`", bid, tc.out)
			}
		})
	}
}
