package bid

type entry struct {
	item     Bid
	priority uint64
}

// Heap keeps bids ordered in a desc/asc fashion based on `priority` (this case it will be ordered by bid amount)
type Heap struct {
	list []entry
	min bool
}

func NewMinHeap(size int) *Heap {
	return &Heap{
		min:  true,
		list: make([]entry, 0, size),
	}
}

func NewMaxHeap(size int) *Heap {
	return &Heap{
		min:  false,
		list: make([]entry, 0, size),
	}
}

// How many items in the heap?
func (h *Heap) Size() int {
	return len(h.list)
}

// Push an item with a priority
func (h *Heap) Push(item Bid, priority uint64) {
	h.list = append(h.list, entry{item, priority})
	h.fixUp(len(h.list) - 1)
}

// Pop the highest (or lowest) priority item
func (h *Heap) Pop() *Bid {
	if len(h.list) == 0 {
		return nil
	}
	i := h.list[0].item
	h.list[0] = h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]
	h.fixDown(0)
	return &i
}

// Peek at the highest (or lowest) priority item
func (h *Heap) Peek() *Bid {
	if len(h.list) == 0 {
		return nil
	}
	return &h.list[0].item
}

type Iterator func() (item interface{}, next Iterator)

func (h *Heap) Items() (it Iterator) {
	i := 0
	return func() (item interface{}, next Iterator) {
		var e entry
		if i < len(h.list) {
			e = h.list[i]
			i++
			return e.item, it
		}
		return nil, nil
	}
}

func (h *Heap) fixUp(k int) {
	parent := (k+1)/2 - 1
	for k > 0 {
		if h.gte(parent, k) {
			return
		}
		h.list[parent], h.list[k] = h.list[k], h.list[parent]
		k = parent
		parent = (k+1)/2 - 1
	}
}

func (h *Heap) fixDown(k int) {
	kid := (k+1)*2 - 1
	for kid < len(h.list) {
		if kid+1 < len(h.list) && h.lt(kid, kid+1) {
			kid++
		}
		if h.gte(k, kid) {
			break
		}
		h.list[kid], h.list[k] = h.list[k], h.list[kid]
		k = kid
		kid = (k+1)*2 - 1
	}
}

func (h *Heap) gte(i, j int) bool {
	if h.min {
		return h.list[i].priority <= h.list[j].priority
	} else {
		return h.list[i].priority >= h.list[j].priority
	}
}

func (h *Heap) lt(i, j int) bool {
	if h.min {
		return h.list[i].priority > h.list[j].priority
	} else {
		return h.list[i].priority < h.list[j].priority
	}
}
