package cmd

import (
	"fmt"

	"github.com/alpeb/go-finance/fin"
	"gonum.org/v1/gonum/stat"
)

// WIP
func stats() {
	xs := []float64{
		32.32, 56.98, 21.52, 44.32,
		55.63, 13.75, 43.47, 43.34,
		12.34,
	}

	mean := stat.Mean(xs, nil)
	fmt.Printf("mean=%v\n", mean)

	rate := float64(0.5)
	npv := fin.NetPresentValue(rate, xs)
	fmt.Printf("npv=%v\n", npv)
}
