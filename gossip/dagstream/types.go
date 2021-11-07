package dagstream

import (
	"github.com/corex-mn/corex-base/hash"
	"github.com/corex-mn/corex-base/inter/dag"
)

type Request struct {
	Session   Session
	Limit     dag.Metric
	Type      RequestType
	MaxChunks uint32
}

type Response struct {
	SessionID uint32
	Done      bool
	IDs       hash.Events
	Events    []interface{}
}

type Session struct {
	ID    uint32
	Start []byte
	Stop  []byte
}

type RequestType uint8

const (
	RequestIDs    RequestType = 0
	RequestEvents RequestType = 2
)
