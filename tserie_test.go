// nolint: dupl
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

func TestTimeIterator(t *testing.T) {
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
			ts := tserie.NewTimeIterator(tt.start, tt.stop, tt.step, tt.getValue)
			var got []tserie.Point
			for ts.Next() {
				got = append(got, ts.Item())
			}
			if !equalSerie(got, tt.want) {
				t.Errorf("Iterator got %v , want %v", got, tt.want)
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

// ExampleSin generates a sinusoidal function with period of one hour
// and samples every one minute
func ExampleSin() {
	start := time.Date(2022, 6, 26, 0, 0, 0, 0, time.UTC)
	stop := time.Date(2022, 6, 26, 1, 0, 0, 0, time.UTC)
	step := time.Minute
	value := tserie.Sin(time.Hour, 1, 0)
	ts := tserie.MakeTS(start, stop, step, value)
	for _, t := range ts {
		fmt.Printf("%02d = %.02f\n", t.Time.Minute(), t.Value)
	}

	//Output:
	// 00 = 0.00
	// 01 = 0.10
	// 02 = 0.21
	// 03 = 0.31
	// 04 = 0.41
	// 05 = 0.50
	// 06 = 0.59
	// 07 = 0.67
	// 08 = 0.74
	// 09 = 0.81
	// 10 = 0.87
	// 11 = 0.91
	// 12 = 0.95
	// 13 = 0.98
	// 14 = 0.99
	// 15 = 1.00
	// 16 = 0.99
	// 17 = 0.98
	// 18 = 0.95
	// 19 = 0.91
	// 20 = 0.87
	// 21 = 0.81
	// 22 = 0.74
	// 23 = 0.67
	// 24 = 0.59
	// 25 = 0.50
	// 26 = 0.41
	// 27 = 0.31
	// 28 = 0.21
	// 29 = 0.10
	// 30 = 0.00
	// 31 = -0.10
	// 32 = -0.21
	// 33 = -0.31
	// 34 = -0.41
	// 35 = -0.50
	// 36 = -0.59
	// 37 = -0.67
	// 38 = -0.74
	// 39 = -0.81
	// 40 = -0.87
	// 41 = -0.91
	// 42 = -0.95
	// 43 = -0.98
	// 44 = -0.99
	// 45 = -1.00
	// 46 = -0.99
	// 47 = -0.98
	// 48 = -0.95
	// 49 = -0.91
	// 50 = -0.87
	// 51 = -0.81
	// 52 = -0.74
	// 53 = -0.67
	// 54 = -0.59
	// 55 = -0.50
	// 56 = -0.41
	// 57 = -0.31
	// 58 = -0.21
	// 59 = -0.10
	// 00 = 0.00
}

// ExampleCos generates a cosinusoidal function with period of one hour
// and samples every one minute
func ExampleCos() {
	start := time.Date(2022, 6, 26, 0, 0, 0, 0, time.UTC)
	stop := time.Date(2022, 6, 26, 1, 0, 0, 0, time.UTC)
	step := time.Minute
	value := tserie.Sin(time.Hour, 1, 0)
	ts := tserie.MakeTS(start, stop, step, value)
	for _, t := range ts {
		fmt.Printf("%02d = %.02f\n", t.Time.Minute(), t.Value)
	}

	//Output:
	// 00 = 0.00
	// 01 = 0.10
	// 02 = 0.21
	// 03 = 0.31
	// 04 = 0.41
	// 05 = 0.50
	// 06 = 0.59
	// 07 = 0.67
	// 08 = 0.74
	// 09 = 0.81
	// 10 = 0.87
	// 11 = 0.91
	// 12 = 0.95
	// 13 = 0.98
	// 14 = 0.99
	// 15 = 1.00
	// 16 = 0.99
	// 17 = 0.98
	// 18 = 0.95
	// 19 = 0.91
	// 20 = 0.87
	// 21 = 0.81
	// 22 = 0.74
	// 23 = 0.67
	// 24 = 0.59
	// 25 = 0.50
	// 26 = 0.41
	// 27 = 0.31
	// 28 = 0.21
	// 29 = 0.10
	// 30 = 0.00
	// 31 = -0.10
	// 32 = -0.21
	// 33 = -0.31
	// 34 = -0.41
	// 35 = -0.50
	// 36 = -0.59
	// 37 = -0.67
	// 38 = -0.74
	// 39 = -0.81
	// 40 = -0.87
	// 41 = -0.91
	// 42 = -0.95
	// 43 = -0.98
	// 44 = -0.99
	// 45 = -1.00
	// 46 = -0.99
	// 47 = -0.98
	// 48 = -0.95
	// 49 = -0.91
	// 50 = -0.87
	// 51 = -0.81
	// 52 = -0.74
	// 53 = -0.67
	// 54 = -0.59
	// 55 = -0.50
	// 56 = -0.41
	// 57 = -0.31
	// 58 = -0.21
	// 59 = -0.10
	// 00 = 0.00
}
