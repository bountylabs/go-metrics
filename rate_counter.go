package metrics

import (
	"sync"
	"sync/atomic"

	"github.com/bountylabs/go-metrics/clock"
)

//go:generate counterfeiter . RateCounter
type RateCounter interface {
	Mark(int64)
	// Rate1() returns the rate up to the last full sampling period, so at time 5.3s it will only return the rate on [0, 5].
	// Count() returns the total count ever, including the current sampling period.
	Count() int64
	Rate1() float64
	Clear()
	Snapshot() RateCounter
}

// A port of https://cgit.twitter.biz/source/tree/src/java/com/twitter/search/common/metrics/SearchRateCounter.java

// A counter that tells you the rate per second that something happened during the past 60 seconds
// (excluding the most recent fractional second).
type StandardRateCounter struct {
	clock clock.Clock

	// lastCount "lags behind" by a sample period by design, this one really counts all events so far.
	// Note that lastCount is at par with lastRate, so that in MonViz rate(meter.lastCount) = meter.lastRate
	counter        int64
	samplePeriodMs int64
	windowSizeMs   int64

	lock sync.RWMutex

	// These values should only be used while holding the lock
	timestampsMs     []int64
	counts           []int64
	headIndex        int // Array index to most recent written value.
	tailIndex        int // Array index to oldest written value.
	lastSampleTimeMs int64
	lastRate         float64
	lastCount        int64
}

// GetOrRegisterRateCounter returns an existing RateCounter or constructs and registers a
// new StandardRateCounter.
func GetOrRegisterRateCounter(name string, r Registry) RateCounter {
	if nil == r {
		r = DefaultRegistry
	}
	return r.GetOrRegister(name, func() RateCounter { return NewStandardRateCounter(60, 1000, clock.New()) }).(RateCounter)
}

func NewRateCounter() RateCounter {
	if UseNilMetrics {
		return NilRateCounter{}
	}
	return NewStandardRateCounter(60, 1000, clock.New())
}

// NewRegisteredRateCounter constructs and registers a new StandardRateCounter.
func NewRegisteredRateCounter(name string, r Registry, clock clock.Clock) RateCounter {
	c := NewStandardRateCounter(60, 1000, clock)

	if nil == r {
		r = DefaultRegistry
	}
	r.Register(name, c)
	return c
}

func NewStandardRateCounter(numSamples int64, samplePeriodMs int64, clock clock.Clock) RateCounter {
	rc := &StandardRateCounter{
		samplePeriodMs: samplePeriodMs,
		windowSizeMs:   numSamples * samplePeriodMs,
		timestampsMs:   make([]int64, numSamples+1),
		counts:         make([]int64, numSamples+1),
		clock:          clock,
	}

	rc.Clear()

	return rc
}

func (srCounter *StandardRateCounter) Clear() {
	srCounter.lock.Lock()
	defer srCounter.lock.Unlock()

	atomic.StoreInt64(&srCounter.counter, 0)

	resetTimeMs := srCounter.clock.Now().UnixNano() / 1e6
	for i, _ := range srCounter.timestampsMs {
		srCounter.timestampsMs[i] = resetTimeMs
		srCounter.counts[i] = 0
	}

	srCounter.lastSampleTimeMs = resetTimeMs
	srCounter.lastRate = 0.0
	srCounter.lastCount = 0

	// Head and tail never point to the same index.
	srCounter.headIndex = 0
	srCounter.tailIndex = len(srCounter.timestampsMs) - 1
}

func (srCounter *StandardRateCounter) Mark(n int64) {
	atomic.AddInt64(&srCounter.counter, n)
	srCounter.maybeSampleCount()
}

func (srCounter *StandardRateCounter) Rate1() float64 {
	srCounter.maybeSampleCount()
	srCounter.lock.RLock()
	defer srCounter.lock.RUnlock()
	return srCounter.lastRate
}

func (srCounter *StandardRateCounter) Count() int64 {
	return atomic.LoadInt64(&srCounter.counter)
}

func (srCounter *StandardRateCounter) Snapshot() RateCounter {
	srCounter.maybeSampleCount()
	srCounter.lock.RLock()
	defer srCounter.lock.RUnlock()

	return &RateCounterSnapshot{
		count: atomic.LoadInt64(&srCounter.counter),
		rate:  srCounter.lastRate,
	}
}

func (srCounter *StandardRateCounter) roundTime(timeMs int64) int64 {
	return timeMs - (timeMs % srCounter.samplePeriodMs)
}

func (srCounter *StandardRateCounter) advance(index int) int {
	return (index + 1) % len(srCounter.counts)
}

/**
 * May sample the current count and timestamp.  Note that this is not an unbiased sampling
 * algorithm, but given that we are computing a rate over a ring buffer of 60 samples, it
 * should not matter in practice.
 */
func (srCounter *StandardRateCounter) maybeSampleCount() {
	currentTimeMs := srCounter.clock.Now().UnixNano() / 1e6
	currentSampleTimeMs := srCounter.roundTime(currentTimeMs)

	srCounter.lock.RLock()
	toSample := currentSampleTimeMs > srCounter.lastSampleTimeMs
	srCounter.lock.RUnlock()

	if !toSample {
		return
	}

	srCounter.lock.Lock()
	defer srCounter.lock.Unlock()

	if currentSampleTimeMs > srCounter.lastSampleTimeMs {
		srCounter.sampleCountAndUpdateRate(currentSampleTimeMs)
	}
}

/**
 * Records a new sample to the ring buffer, advances head and tail if necessary, and
 * recomputes the rate.
 */
func (srCounter *StandardRateCounter) sampleCountAndUpdateRate(currentSampleTimeMs int64) {
	// Record newest up to date second sample time.  Clear rate.
	srCounter.lastSampleTimeMs = currentSampleTimeMs

	// Advance head and write values.
	srCounter.headIndex = srCounter.advance(srCounter.headIndex)
	srCounter.timestampsMs[srCounter.headIndex] = currentSampleTimeMs

	srCounter.lastCount = atomic.LoadInt64(&srCounter.counter)
	srCounter.counts[srCounter.headIndex] = srCounter.lastCount

	// Ensure tail is always ahead of head.
	if srCounter.tailIndex == srCounter.headIndex {
		srCounter.tailIndex = srCounter.advance(srCounter.tailIndex)
	}

	// Advance the 'tail' to the newest sample which is at least windowTimeMs old.
	for {
		nextWindowStart := srCounter.advance(srCounter.tailIndex)
		if nextWindowStart == srCounter.headIndex ||
			srCounter.timestampsMs[srCounter.headIndex]-srCounter.timestampsMs[nextWindowStart] < srCounter.windowSizeMs {
			break
		}
		srCounter.tailIndex = nextWindowStart
	}

	timeDeltaMs := srCounter.timestampsMs[srCounter.headIndex] - srCounter.timestampsMs[srCounter.tailIndex]
	if timeDeltaMs == 0 {
		srCounter.lastRate = 0.0
	} else {
		if timeDeltaMs > srCounter.windowSizeMs {
			timeDeltaMs = srCounter.windowSizeMs
		}

		deltaTimeSecs := timeDeltaMs / 1000.0
		deltaCount := srCounter.counts[srCounter.headIndex] - srCounter.counts[srCounter.tailIndex]
		if deltaTimeSecs <= 0.0 {
			srCounter.lastRate = 0
		} else {
			srCounter.lastRate = float64(deltaCount) / float64(deltaTimeSecs)
		}
	}
}

type RateCounterSnapshot struct {
	rate  float64
	count int64
}

func (rcSnapshot *RateCounterSnapshot) Mark(n int64) {
	panic("Mark called on RateCounterSnapshot")
}

func (rcSnapshot *RateCounterSnapshot) Count() int64 {
	return rcSnapshot.count
}

func (rcSnapshot *RateCounterSnapshot) Rate1() float64 {
	return rcSnapshot.rate
}

func (rcSnapshot *RateCounterSnapshot) Clear() {
	panic("Clear called on RateCounterSnapshot")
}

func (rcSnapshot *RateCounterSnapshot) Snapshot() RateCounter {
	return rcSnapshot
}

type NilRateCounter struct{}

func (NilRateCounter) Mark(int64) {
}

func (NilRateCounter) Count() int64 {
	return 0
}

func (NilRateCounter) Rate1() float64 {
	return 0
}

func (NilRateCounter) Clear() {
}

func (NilRateCounter) Snapshot() RateCounter {
	return NilRateCounter{}
}
