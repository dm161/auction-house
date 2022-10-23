package hearthbeat

import (
	"testing"
)

func TestParse(t *testing.T) {
	for name, tc := range map[string]struct {
		inStr string
		out   Hearthbeat
		err   error
	}{
		"returns error on invalid number of fields": {
			inStr: "1|2",
			out:   Hearthbeat{},
			err:   errInvalidHearthbeatFormat,
		},
		"returns error on invalid timestamp": {
			inStr: "foobar",
			out:   Hearthbeat{},
			err:   errInvalidHearthbeatValue,
		},
		"returns valid bid": {
			inStr: "20",
			out: Hearthbeat{
				Timestamp: uint64(20),
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
