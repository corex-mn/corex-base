package abft

import (
	"github.com/corex-mn/corex-base/hash"
	"github.com/corex-mn/corex-base/inter/dag"
)

// EventSource is a callback for getting events from an external storage.
type EventSource interface {
	HasEvent(hash.Event) bool
	GetEvent(hash.Event) dag.Event
}
