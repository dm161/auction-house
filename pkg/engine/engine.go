package engine

import (
	"fmt"

	"diego.pizza/auction-house/pkg/bid"
	"diego.pizza/auction-house/pkg/hearthbeat"
	"diego.pizza/auction-house/pkg/listing"
)

// heap size for storing asc/desc sorted bids
const heapSize = 1 << 5

type Engine interface {
	BiddingEngine
	ListingEngine
	HearthbeatEngine
}

type BiddingEngine interface {
	Bid(bid.Bid) error
}

type ListingEngine interface {
	List(listing.Listing) error
}

type HearthbeatEngine interface {
	Hearthbeat(hearthbeat.Hearthbeat) []string
}

type listingBids struct {
	anyAsc  *bid.Heap // min heap
	anyDesc *bid.Heap // max heap

	atReservePrice *bid.Heap // max heap
}

func newListingBids() *listingBids {
	lb := &listingBids{
		anyAsc:         bid.NewMinHeap(heapSize),
		anyDesc:        bid.NewMaxHeap(heapSize),
		atReservePrice: bid.NewMaxHeap(heapSize),
	}

	return lb
}

func (lb *listingBids) Add(l listing.Listing, b bid.Bid) {
	lb.anyAsc.Push(b, b.Amount)
	lb.anyDesc.Push(b, b.Amount)
	if l.ReservePrice <= b.Amount {
		lb.atReservePrice.Push(b, b.Amount)
	}
}

func (lb *listingBids) Unsold() bool {
	return lb.atReservePrice.Size() == 0
}

type engine struct {
	listingsBids map[string]*listingBids
	listings     map[string]*listing.Listing
	userBids     map[string]*bid.Heap // min heap
}

func New() Engine {
	return &engine{
		listingsBids: make(map[string]*listingBids, 0),
		listings:     make(map[string]*listing.Listing, 0),
		userBids:     make(map[string]*bid.Heap, 0),
	}
}

func (e *engine) Bid(b bid.Bid) error {
	l, ok := e.listings[b.Item]
	if !ok {
		return errItemNotForSale
	}

	_, exists := e.userBids[b.UserID]
	if !exists {
		userBids := bid.NewMinHeap(heapSize)
		userBids.Push(b, b.Amount)
		e.userBids[b.UserID] = userBids
	}
	userBids, exists := e.userBids[b.UserID]

	// skip if user has already an higher bid
	if lastBid := userBids.Peek(); lastBid.Amount < b.Amount {
		return nil
	}

	e.userBids[b.UserID].Push(b, b.Amount)
	e.listingsBids[b.Item].Add(*l, b)

	return nil
}

func (e *engine) List(listing listing.Listing) error {
	_, exists := e.listings[listing.Item]
	if exists {
		return errListingAlreadyExists
	}
	e.listings[listing.Item] = &listing
	e.listingsBids[listing.Item] = newListingBids()

	return nil
}

func (e *engine) Hearthbeat(hearthbeat hearthbeat.Hearthbeat) []string {
	var res []string

	for _, l := range e.expiredListings(hearthbeat) {

		// if no bids return unsold
		if e.listingsBids[l.Item].Unsold() {
			res = append(res,
				fmt.Sprintf(
					"%d|%s||UNSOLD|0.00|%d|%d|%d",
					hearthbeat.Timestamp,
					l.Item,
					e.listingsBids[l.Item].anyDesc.Size(),
					e.listingsBids[l.Item].anyDesc.Peek().Amount,
					e.listingsBids[l.Item].anyAsc.Peek().Amount,
				),
			)
			continue
		}

		// there was only one valid bid for this item
		if e.listingsBids[l.Item].atReservePrice.Size() == 1 {
			res = append(res,
				fmt.Sprintf(
					"%d|%s|%s|SOLD|%d|%d|%d|%d",
					hearthbeat.Timestamp,
					l.Item,
					e.listingsBids[l.Item].atReservePrice.Peek().UserID,
					l.ReservePrice,
					e.listingsBids[l.Item].anyDesc.Size(),
					e.listingsBids[l.Item].anyDesc.Peek().Amount,
					e.listingsBids[l.Item].anyAsc.Peek().Amount,
				),
			)
			continue
		}

		// there were at least 2 valid bids
		highestBid := e.listingsBids[l.Item].atReservePrice.Pop()
		secondHighestBid := e.listingsBids[l.Item].atReservePrice.Pop()
		res = append(res,
			fmt.Sprintf(
				"%d|%s|%s|SOLD|%d|%d|%d|%d",
				hearthbeat.Timestamp,
				l.Item,
				highestBid.UserID,
				secondHighestBid.Amount,
				e.listingsBids[l.Item].anyDesc.Size(),
				e.listingsBids[l.Item].anyDesc.Peek().Amount,
				e.listingsBids[l.Item].anyAsc.Peek().Amount,
			),
		)
	}

	return res
}

func (e *engine) expiredListings(hearthbeat hearthbeat.Hearthbeat) []*listing.Listing {
	var ls []*listing.Listing
	for _, l := range e.listings {
		if l.CloseTime <= hearthbeat.Timestamp {
			ls = append(ls, l)
		}
	}

	return ls
}
