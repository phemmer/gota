package gota

import (
	"math"
)

// ADX - Average Directional Movement Index (https://en.wikipedia.org/wiki/Average_directional_movement_index)
type ADX struct {
	high  float64
	low   float64
	pdmMA AlgSimple
	ndmMA AlgSimple
	dmMA  AlgSimple
	atr   AlgTri
	init  bool
}

// NewADXMA constructs a new ADX with the given moving average algorithm.
func NewADXMA(inTimePeriod int, maNew AlgSimpleConstructor) *ADX {
	return &ADX{
		pdmMA: maNew(inTimePeriod),
		ndmMA: maNew(inTimePeriod),
		dmMA:  maNew(inTimePeriod),
		atr:   NewATRMA(inTimePeriod, maNew),
	}
}

// NewADX constructs a new ADX using a WSMA as the moving average algorithm.
// Note that this does not produce the same values as Ta-Lib. It instead complies with the seemingly more common
// algorithm which can be found at:
//  https://school.stockcharts.com/doku.php?id=technical_indicators:average_directional_index_adx
// and
//  https://github.com/bukosabino/ta
// Both claim to be compliant with Wilder's implementation of ADX, but that can't be, as ta-lib produces different
// results.
//
// Note that the current implementation of this algorithm does not return meaningful results until after warmup is
// completed. This may be changed in the future to produce semi-meaningful results, similar to the other algorithms.
func NewADX(inTimePeriod int, warmupType WarmupType) *ADX {
	return NewADXMA(inTimePeriod, WSMAConstructor{warmupType}.New)
}

// Add adds a new sample value to the algorithm and returns the computed value.
func (adx *ADX) Add(high, low, close float64) float64 {
	atr := adx.atr.Add(high, low, close)

	umove := high - adx.high
	dmove := adx.low - low
	adx.high = high
	adx.low = low
	if !adx.init {
		adx.init = true
		return 0
	}

	var pdm, ndm float64
	if umove > 0 && umove > dmove {
		pdm = umove
	}
	if dmove > 0 && dmove > umove {
		ndm = dmove
	}

	pdma := adx.pdmMA.Add(pdm)
	ndma := adx.ndmMA.Add(ndm)

	if !adx.pdmMA.Warmed() /* || !adx.atr.Warmed() */ {
		// The atr check shouldn't be necessary as it should have warmed before pdmMA & ndmMA
		return 0
	}

	pdi := 100 * pdma / atr
	ndi := 100 * ndma / atr

	var dx float64
	if pdi+ndi != 0 {
		dx = math.Abs((pdi - ndi) / (pdi + ndi))
	}
	cadx := 100 * adx.dmMA.Add(dx)
	return cadx
}

func (adx ADX) Warmed() bool {
	return adx.dmMA.Warmed()
}

func (adx ADX) WarmCount() int {
	return adx.dmMA.WarmCount() + adx.pdmMA.WarmCount() + 1
}

/*
type AvgSum struct {
	warmCount int
	count     int
	last      float64
}

func NewAvgSum(inTimePeriod int) *AvgSum {
	return &AvgSum{
		warmCount: inTimePeriod,
	}
}

func (as *AvgSum) Add(v float64) float64 {
	if !as.Warmed() {
		as.last += v
		as.count++
		return as.last / float64(as.count) * float64(as.warmCount)
	}

	as.last = as.last - as.last/float64(as.count) + v
	return as.last
}

func (as *AvgSum) Warmed() bool {
	return as.count == as.warmCount
}

func (as *AvgSum) WarmCount() int {
	return as.warmCount - 1
}

type adxTR struct {
	count     int
	warmCount int
	prevClose float64
	avgsum    float64
}

func newADXTR(inTimePeriod int) *adxTR {
	return &adxTR{
		warmCount: inTimePeriod + 1,
	}
}

func (tr *adxTR) Add(high, low, close float64) float64 {
	v := high - low
	if tr.count != 0 {
		v = math.Max(v, math.Abs(high-tr.prevClose))
		v = math.Max(v, math.Abs(low-tr.prevClose))
	}

	tr.prevClose = close
	if tr.count != tr.warmCount {
		tr.count++
		if tr.count == 1 {
			return v * float64(tr.warmCount - 1)
		}
		tr.avgsum += v
		return tr.avgsum / float64(tr.count-1) * float64(tr.warmCount-1)
	}
	tr.avgsum = tr.avgsum - tr.avgsum/float64(tr.warmCount-1) + v
	return tr.avgsum
}

func (tr adxTR) Warmed() bool {
	return tr.count == tr.warmCount
}

func (tr adxTR) WarmCount() int {
	return 2
}
*/
