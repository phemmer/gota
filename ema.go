package gota

import (
	"bytes"
	"fmt"
)

type WarmupType uint8

const (
	// No alternate warmup algorithm.
	// WarmNone is highly discouraged as it excessively weights the initial results towards the first values.
	WarmNone WarmupType = iota

	// Simple Moving Average
	WarmSMA

	// Exponential Moving Average
	// WarmEMA scales the EMA's alpha value during the warmup period.
	// Meaning it acts as if the algorithms's configured period for the first sample is 1, second sample 2, etc.
	// Some algorithms have a shorter warmup period when this type is used.
	WarmEMA
)

func (wt *WarmupType) UnmarshalJSON(bs []byte) error {
	if bytes.Equal([]byte(`"EMA"`), bs) {
		*wt = WarmEMA
	}
	if bytes.Equal([]byte(`"SMA"`), bs) {
		*wt = WarmSMA
	}
	return fmt.Errorf("unknown warmup type %s", bs)
}
func (wt WarmupType) MarshalJSON() ([]byte, error) {
	switch wt {
	case WarmEMA:
		return []byte(`"EMA"`), nil
	case WarmSMA:
		return []byte(`"SMA"`), nil
	}
	return nil, fmt.Errorf("unknown warmup type %d", wt)
}

// EMA - Exponential Moving Average (http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:moving_averages#exponential_moving_average_calculation)
type EMA struct {
	inTimePeriod int
	last         float64
	count        int
	alpha        float64
	warmType     WarmupType
}

// EMAConstructor is for passing to algorithms with swappable moving average implementations.
type EMAConstructor struct {
	WarmupType WarmupType
}

func (c EMAConstructor) New(inTimePeriod int) AlgSingle {
	return NewEMA(inTimePeriod, c.WarmupType)
}

// NewEMA constructs a new EMA.
//
// When warmed with WarmSMA the first inTimePeriod samples will result in a simple average, switching to exponential moving average after warmup is complete.
//
// When warmed with WarmEMA the algorithm immediately starts using an exponential moving average for the output values. During the warmup period the alpha value is scaled to prevent unbalanced weighting on initial values.
func NewEMA(inTimePeriod int, warmType WarmupType) *EMA {
	return &EMA{
		inTimePeriod: inTimePeriod,
		alpha:        2 / float64(inTimePeriod+1),
		warmType:     warmType,
	}
}

// WarmCount returns the number of samples that must be provided for the algorithm to be fully "warmed".
func (ema *EMA) WarmCount() int {
	return ema.inTimePeriod - 1
}

// Warmed indicates whether the algorithm has enough data to generate accurate results.
func (ema *EMA) Warmed() bool {
	return ema.count == ema.inTimePeriod
}

// Last returns the last output value.
func (ema *EMA) Last() float64 {
	return ema.last
}

// Add adds a new sample value to the algorithm and returns the computed value.
func (ema *EMA) Add(v float64) float64 {
	var avg float64
	if ema.count == 0 {
		avg = v
	} else {
		lastAvg := ema.Last()
		if !ema.Warmed() && ema.warmType != WarmNone {
			if ema.warmType == WarmSMA {
				avg = (lastAvg*float64(ema.count) + v) / float64(ema.count+1)
			} else { // ema.warmType == WarmEMA
				// Scale the alpha so that we don't excessively weight the result towards the first value.
				// We adjust using the existing alpha instead of calculating from scratch (which would be simpler/faster) as
				// WSMA overrides the alpha with a custom numerator.
				alpha := ema.alpha * float64(ema.inTimePeriod+1) / float64(ema.count+2)
				avg = (v-lastAvg)*alpha + lastAvg
			}
		} else {
			avg = (v-lastAvg)*ema.alpha + lastAvg
		}
	}

	ema.last = avg
	if ema.count < ema.inTimePeriod {
		// don't just keep incrementing to prevent potential overflow
		ema.count++
	}
	return avg
}

// DEMA - Double Exponential Moving Average (https://en.wikipedia.org/wiki/Double_exponential_moving_average)
type DEMA struct {
	ema1 EMA
	ema2 EMA
}

// DEMAConstructor is for passing to algorithms with swappable moving average implementations.
type DEMAConstructor struct {
	WarmupType WarmupType
}

func (c DEMAConstructor) New(inTimePeriod int) AlgSingle {
	return NewDEMA(inTimePeriod, c.WarmupType)
}

// NewDEMA constructs a new DEMA.
//
// When warmed with WarmSMA the first inTimePeriod samples will result in a simple average, switching to exponential moving average after warmup is complete.
//
// When warmed with WarmEMA the algorithm immediately starts using an exponential moving average for the output values. During the warmup period the alpha value is scaled to prevent unbalanced weighting on initial values.
func NewDEMA(inTimePeriod int, warmType WarmupType) *DEMA {
	return &DEMA{
		ema1: *NewEMA(inTimePeriod, warmType),
		ema2: *NewEMA(inTimePeriod, warmType),
	}
}

// WarmCount returns the number of samples that must be provided for the algorithm to be fully "warmed".
func (dema *DEMA) WarmCount() int {
	if dema.ema1.warmType == WarmEMA {
		return dema.ema1.WarmCount()
	}
	return dema.ema1.WarmCount() + dema.ema2.WarmCount()
}

// Add adds a new sample value to the algorithm and returns the computed value.
func (dema *DEMA) Add(v float64) float64 {
	avg1 := dema.ema1.Add(v)
	var avg2 float64
	if dema.ema1.Warmed() || dema.ema1.warmType == WarmEMA {
		avg2 = dema.ema2.Add(avg1)
	} else {
		avg2 = avg1
	}
	return 2*avg1 - avg2
}

// Warmed indicates whether the algorithm has enough data to generate accurate results.
func (dema *DEMA) Warmed() bool {
	return dema.ema2.Warmed()
}

// TEMA - Triple Exponential Moving Average (https://en.wikipedia.org/wiki/Triple_exponential_moving_average)
type TEMA struct {
	ema1 EMA
	ema2 EMA
	ema3 EMA
}

// TEMAConstructor is for passing to algorithms with swappable moving average implementations.
type TEMAConstructor struct {
	WarmupType WarmupType
}

func (c TEMAConstructor) New(inTimePeriod int) AlgSingle {
	return NewEMA(inTimePeriod, c.WarmupType)
}

// NewTEMA constructs a new TEMA.
//
// When warmed with WarmSMA the first inTimePeriod samples will result in a simple average, switching to exponential moving average after warmup is complete.
//
// When warmed with WarmEMA the algorithm immediately starts using an exponential moving average for the output values. During the warmup period the alpha value is scaled to prevent unbalanced weighting on initial values.
func NewTEMA(inTimePeriod int, warmType WarmupType) *TEMA {
	return &TEMA{
		ema1: *NewEMA(inTimePeriod, warmType),
		ema2: *NewEMA(inTimePeriod, warmType),
		ema3: *NewEMA(inTimePeriod, warmType),
	}
}

// WarmCount returns the number of samples that must be provided for the algorithm to be fully "warmed".
func (tema *TEMA) WarmCount() int {
	if tema.ema1.warmType == WarmEMA {
		return tema.ema1.WarmCount()
	}
	return tema.ema1.WarmCount() + tema.ema2.WarmCount() + tema.ema3.WarmCount()
}

// Add adds a new sample value to the algorithm and returns the computed value.
func (tema *TEMA) Add(v float64) float64 {
	avg1 := tema.ema1.Add(v)
	var avg2 float64
	if tema.ema1.Warmed() || tema.ema1.warmType == WarmEMA {
		avg2 = tema.ema2.Add(avg1)
	} else {
		avg2 = avg1
	}
	var avg3 float64
	if tema.ema2.Warmed() || tema.ema2.warmType == WarmEMA {
		avg3 = tema.ema3.Add(avg2)
	} else {
		avg3 = avg2
	}
	return 3*avg1 - 3*avg2 + avg3
}

// Warmed indicates whether the algorithm has enough data to generate accurate results.
func (tema *TEMA) Warmed() bool {
	return tema.ema3.Warmed()
}
