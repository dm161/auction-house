package hearthbeat

const (
	hearthbeat = iota

	hearthbeatFields = 1
)

type Hearthbeat struct {
	Timestamp uint64
}
