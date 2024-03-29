# tserie

Time Series data generator

[![Go Reference](https://pkg.go.dev/badge/github.com/licaonfee/tserie.svg)](https://pkg.go.dev/github.com/licaonfee/tserie)

Example

```go
package main

import (
    "fmt"
    "time"

    "github.com/licaonfee/tserie"
)

func main() {
    start := time.Date(2022, 6, 26, 0, 0, 0, 0, time.UTC)
    stop := time.Date(2022, 6, 26, 1, 0, 0, 0, time.UTC)
    step := time.Minute
    value := tserie.Sine(time.Hour, 1, 0)
    ts := tserie.MakeTS(start, stop, step, value)
    for _, t := range ts {
        fmt.Printf("%02d = %.02f\n", t.Time.Minute(), t.Value)
    }
}

// This will generate 
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
```
