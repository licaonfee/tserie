package tserie

import (
	"math/rand"
	"time"
)

// Point contains a pair time,value
type Point struct {
	Time  time.Time
	Value float64
}

// MakeTS create a new time serie between start and stop with all the time values given by step
func MakeTS(start, stop time.Time, step time.Duration, getValue func(time.Time) float64) []Point {
	r := stop.Sub(start) / step
	ts := make([]Point, 0, int(r))
	for t := start; t.Before(stop); t = t.Add(step) {
		p := Point{
			Time:  t,
			Value: getValue(t),
		}
		ts = append(ts, p)
	}
	return ts
}

// Normal returns a normally distributed float64 as
// rand.NormFloat64() * std + mean
func Normal(std, mean float64) func(time.Time) float64 {
	return func(time.Time) float64 {
		return (rand.NormFloat64() * std) + mean
	}
}

// TimeIterator generates datapoints in the same way as MakeTS but is suitable for big datasets
type TimeIterator struct {
	start time.Time
	stop  time.Time
	step  time.Duration
	curr  Point
	gen   func(time.Time) float64
}

// Next returns true while current time is before stop
func (it *TimeIterator) Next() bool {
	if !it.curr.Time.Before(it.stop) {
		return false
	}
	t := it.curr.Time.Add(it.step)
	it.curr = Point{
		Time:  t,
		Value: it.gen(t),
	}
	return true
}

//  Item returns the current data point
func (it *TimeIterator) Item() Point {
	return it.curr
}

// NewTimeIterator creates a new TimeIterator
func NewTimeIterator(start, stop time.Time, step time.Duration, gen func(time.Time) float64) *TimeIterator {
	return &TimeIterator{
		start: start,
		stop:  stop,
		step:  step,
		gen:   gen,
	}
}
