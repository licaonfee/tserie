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
