package middleware

import (
	"sync/atomic"
	"time"
)

type Metrics struct {
	totalRequests     atomic.Uint64
	totalErrors       atomic.Uint64
	totalLatencyNanos atomic.Int64
}

func NewMetrics() *Metrics {
	return &Metrics{}
}

func (m *Metrics) Record(status int, duration time.Duration) {
	if m == nil {
		return
	}

	m.totalRequests.Add(1)
	if status >= 500 {
		m.totalErrors.Add(1)
	}
	m.totalLatencyNanos.Add(duration.Nanoseconds())
}

type MetricsSnapshot struct {
	TotalRequests uint64
	TotalErrors   uint64
	AvgLatencyMs  float64
}

func (m *Metrics) Snapshot() MetricsSnapshot {
	if m == nil {
		return MetricsSnapshot{}
	}

	req := m.totalRequests.Load()
	errs := m.totalErrors.Load()
	totalLatency := m.totalLatencyNanos.Load()

	var avg float64
	if req > 0 {
		avg = float64(totalLatency) / float64(req) / 1e6
	}

	return MetricsSnapshot{
		TotalRequests: req,
		TotalErrors:   errs,
		AvgLatencyMs:  avg,
	}
}
