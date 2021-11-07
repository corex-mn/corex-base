package peerleecher

import (
	"time"

	"github.com/corex-mn/corex-base/inter/dag"
)

type EpochDownloaderConfig struct {
	RecheckInterval        time.Duration
	DefaultChunkSize       dag.Metric
	ParallelChunksDownload int
}
