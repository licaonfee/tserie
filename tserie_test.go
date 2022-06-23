package tserie_test

import (
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/licaonfee/tserie"
)

func TestMakeTS(t *testing.T) {
	tests := map[string]struct {
		start    time.Time
		stop     time.Time
		step     time.Duration
		getValue func(time.Time) float64
		want     []tserie.Point
	}{
		"hourly": {
			start:    time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			stop:     time.Date(2022, 1, 1, 0, 59, 59, 0, time.UTC),
			step:     10 * time.Minute,
			getValue: func(t time.Time) float64 { return 0.6 },
			want: []tserie.Point{
				{Time: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), Value: 0.6},
				{Time: time.Date(2022, 1, 1, 0, 10, 0, 0, time.UTC), Value: 0.6},
				{Time: time.Date(2022, 1, 1, 0, 20, 0, 0, time.UTC), Value: 0.6},
				{Time: time.Date(2022, 1, 1, 0, 30, 0, 0, time.UTC), Value: 0.6},
				{Time: time.Date(2022, 1, 1, 0, 40, 0, 0, time.UTC), Value: 0.6},
				{Time: time.Date(2022, 1, 1, 0, 50, 0, 0, time.UTC), Value: 0.6},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tserie.MakeTS(tt.start, tt.stop, tt.step, tt.getValue)
			if !equalSerie(got, tt.want) {
				t.Errorf("MakeTS got %v , want %v", got, tt.want)
			}
		})
	}

}

// test if a serie is equal to another, with precision of nanoseconds
func equalSerie(a, b []tserie.Point) bool {
	if a == nil && b == nil || len(a) == 0 && len(b) == 0 {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	min := int(math.Min(float64(len(a)), float64(len(b))))

	for i := 0; i < min; i++ {
		if a[i].Time.UnixNano() != b[i].Time.UnixNano() || a[i].Value != b[i].Value {
			fmt.Printf("[%d] %d , %d \n", i, a[i].Time.Unix(), b[i].Time.Unix())
			return false
		}
	}
	return true
}

func BenchmarkMakeTS(b *testing.B) {
	// one day full
	start := time.Date(2022, 6, 26, 0, 0, 0, 0, time.UTC)
	stop := time.Date(2022, 6, 27, 0, 0, 0, 0, time.UTC)
	step := time.Second
	value := func(time.Time) float64 {
		return 0.0
	}
	for i := 0; i < b.N; i++ {
		_ = tserie.MakeTS(start, stop, step, value)
	}
}
