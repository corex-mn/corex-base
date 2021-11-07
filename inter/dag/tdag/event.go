package tdag

import (
	"github.com/corex-mn/corex-base/hash"
	"github.com/corex-mn/corex-base/inter/dag"
)

type TestEvent struct {
	dag.MutableBaseEvent
	Name string
}

func (e *TestEvent) AddParent(id hash.Event) {
	parents := e.Parents()
	parents.Add(id)
	e.SetParents(parents)
}
