package gota

import (
	"math"
)

// ATR - Average True Range (https://en.wikipedia.org/wiki/Average_true_range)
type ATR struct {
	prevClose float64
	ma AlgSimple

	init bool
}

// NewATRMA constructs a new ATR with the given moving average algorithm.
func NewATRMA(inTimePeriod int, maNew AlgSimpleConstructor) *ATR {
	return &ATR{
		ma: maNew(inTimePeriod),
	}
}

// NewATR constructs a new ATR with the default Wilder's Smoothing.
func NewATR(inTimePeriod int) *ATR {
	return NewATRMA(inTimePeriod, WSMAConstructor{WarmSMA}.New)
}

// Add adds a new sample value to the algorithm and returns the computed value.
func (atr *ATR) Add(high, low, close float64) float64 {
	var catr float64
	if atr.init {
		tr := high - low
		tr = math.Max(tr, math.Abs(high - atr.prevClose))
		tr = math.Max(tr, math.Abs(low - atr.prevClose))

		catr = atr.ma.Add(tr)
	} else {
		tr := high - low
		catr = tr
		atr.init = true
	}

	atr.prevClose = close
	return catr
}

// Warmed indicates whether the algorithm has enough data to generate accurate results.
func (atr ATR) Warmed() bool {
	return atr.ma.Warmed()
}

// WarmCount returns the number of samples that must be provided for the algorithm to be fully "warmed".
func (atr ATR) WarmCount() int {
	return atr.ma.WarmCount() + 1
}
