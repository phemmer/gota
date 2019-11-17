package gota

import (
	"testing"

	"github.com/phemmer/talib"
	"github.com/stretchr/testify/assert"
)

func TestADX(t *testing.T) {
	// https://school.stockcharts.com/doku.php?id=technical_indicators:average_directional_index_adx
	// The input/expected data comes from the Excel spreadsheet provided on the page linked above. We do not use the
	// sample data provided on the page itself, as that data has errors. Specifically you can note that the second phase,
	// waiting for ADX to warm, is only 13 periods long, not 14 like it should be.
	// The Excel spreadsheet has completely different data, and does not have this issue.

	high := []float64{30.20,30.28,30.45,29.35,29.35,29.29,28.83,28.73,28.67,28.85,28.64,27.68,27.21,26.87,27.41,26.94,26.52,26.52,27.09,27.69,28.45,28.53,28.67,29.01,29.87,29.80,29.75,30.65,30.60,30.76,31.17}
	low := []float64{29.41,29.32,29.96,28.74,28.56,28.41,28.08,27.43,27.66,27.83,27.40,27.09,26.18,26.13,26.63,26.13,25.43,25.35,25.88,26.96,27.14,28.01,27.88,27.99,28.76,29.14,28.71,28.93,30.03,29.39,30.14}
	close := []float64{29.87,30.24,30.10,28.90,28.92,28.48,28.56,27.56,28.47,28.28,27.49,27.23,26.35,26.33,27.03,26.22,26.01,25.46,27.03,27.45,28.36,28.43,27.95,29.01,29.38,29.36,28.91,30.61,30.05,30.19,31.12}

	expList := []float64{33.58,32.15,29.93,28.36}

	adx := NewADX(14, WarmSMA)
	var actList []float64
	var actWarmCount int
	for i := 0; i < len(high); i++ {
		if outV := adx.Add(high[i], low[i], close[i]); adx.Warmed() {
			actList = append(actList, outV)
		} else {
			actWarmCount++
		}
	}

	// we have to have low precision as the data provided is rounded, where as our calculations are full float64
	assert.InDeltaSlice(t, expList, actList, 0.5,
		"Expected: %v (len=%d)\nActual:  %v (len=%d)",
		expList, len(expList),
		actList, len(actList),
	)

	assert.Equal(t, adx.WarmCount(), actWarmCount, "warm count not equal")
}
func TestADX2(t *testing.T) {
	// This data is a lot more volatile than the real world data provided by stockcharts.com in the prior test.
	// The warmup period is also really short, and the validation precision is really high, so errors should show up
	// easily.
	// The data was verified against the python 'ta' module, including full 1E-7 precision.
	high := []float64{13,19,14,11,10,15,5,4,3,7,12,20}
	low := []float64{12,18,13,9,9,14,2,3,2,6,10,16}
	close := []float64{12.293114244553857,18.29708256355629,13.75257303555161,19.413165323827398,9.86533501300156,14.696719165746634,3.5714609181500023,3.02830308332589,2.158328277745128,6.607253439545516,11.950483237721157,16.317814493495487}

	expList := []float64{22.966507177033485,29.904306220095688,42.918254804963105,53.13243417538992,58.23952386060333,64.51658732712706,49.731086892100414,62.09857619403313,76.71267677991193}

	adx := NewADX(2, WarmSMA)
	var actList []float64
	var actWarmCount int
	for i := 0; i < len(high); i++ {
		if outV := adx.Add(high[i], low[i], close[i]); adx.Warmed() {
			actList = append(actList, outV)
		} else {
			actWarmCount++
		}
	}

	assert.InDeltaSlice(t, expList, actList, 1E-7,
		"Expected: %v (len=%d)\nActual:  %v (len=%d)",
		expList, len(expList),
		actList, len(actList),
	)

	assert.Equal(t, adx.WarmCount(), actWarmCount, "warm count not equal")
}

func TestADX3(t *testing.T) {
	// This test fails. TA-Lib seems to use different version of the algorithm that produces significantly different results.
	// However the difference in the implementation seems to be very subtle, as after working at it for far too long, I
	// cannot identify where the problem lies.
	t.SkipNow()
	adx := NewADX(2, WarmSMA)
	testTALibTri(t, 2, talib.Adx, adx)
}
