package bid

import "testing"

func TestMaxHeap(t *testing.T) {
	for name, tc := range map[string]struct {
		in  []Bid
		out []Bid
	}{
		"empty heap": {
			in: []Bid{},
			out: []Bid{},
		},
		"single item": {
			in: []Bid{{Amount: 1}},
			out: []Bid{{Amount: 1}},
		},
		"many items": {
			in: []Bid{{Amount: 1}, {Amount: 9}, {Amount: 0}, {Amount: 4}},
			out: []Bid{{Amount: 9}, {Amount: 4}, {Amount: 1}, {Amount: 0}},
		},
	} {
		t.Run(name, func(t *testing.T) {
			h := NewMaxHeap(1<<4)
			for _, i :=  range tc.in {
				h.Push(i, i.Amount)
			}
			for _, j := range tc.out {
				i := h.Pop()
				if j.Amount != i.Amount {
					t.Errorf("expected `%v` got `%v`", j, i)
				}
			}
		})
	}
}

func TestMinHeap(t *testing.T) {
	for name, tc := range map[string]struct {
		in  []Bid
		out []Bid
	}{
		"empty heap": {
			in: []Bid{},
			out: []Bid{},
		},
		"single item": {
			in: []Bid{{Amount: 1}},
			out: []Bid{{Amount: 1}},
		},
		"many items": {
			in: []Bid{{Amount: 1}, {Amount: 9}, {Amount: 0}, {Amount: 4}},
			out: []Bid{{Amount: 0}, {Amount: 1}, {Amount: 4}, {Amount: 9}},
		},
	} {
		t.Run(name, func(t *testing.T) {
			h := NewMinHeap(1<<4)
			for _, i :=  range tc.in {
				h.Push(i, i.Amount)
			}
			for _, j := range tc.out {
				i := h.Pop()
				if j.Amount != i.Amount {
					t.Errorf("expected `%v` got `%v`", j, i)
				}
			}
		})
	}
}
