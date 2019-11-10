package gota

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func randList(size int) []float64 {
	min := float64(1)
	max := float64(20)
	list := make([]float64, size)
	for i := 0; i < size; i++ {
		list[i] = math.Floor(rand.Float64() * (max - min) + min)
	}
	return list
}

func dataHighLowCloseVolume(size int) ([]float64, []float64, []float64, []float64) {
	low := randList(size)

	high := make([]float64, len(low))
	for i := 0; i < size; i++ {
		high[i] = low[i] + math.Floor(rand.Float64() * 5)
	}

	close := make([]float64, size)
	for i := 0; i < size; i++ {
		close[i] = rand.Float64() * (high[i] - low[i]) + low[i]
	}

	volume := randList(size)

	return high, low, close, volume
}

//type TALibSimple0Per func([]float64, []float64) ([]float64, int)

type TALibSimple func([]float64, int, []float64) ([]float64, int)
func testTALibSimple(t *testing.T, inTimePeriod int, alg AlgSimple, taAlg TALibSimple) {
	list := randList(alg.WarmCount() + 3)

	expList, _ := taAlg(list, inTimePeriod, nil)

	var actWarmCount int
	var warmList []float64
	var actList []float64
	for i := 0; i < len(list); i++ {
		if vOut := alg.Add(list[i]); alg.Warmed() {
			warmList = append(warmList, list[i])
			actList = append(actList, vOut)
		} else {
			actWarmCount++
		}
	}

	if !alg.Warmed() {
		t.Errorf("algorithm did not warm up within time period")
		t.FailNow()
	}

	assert.InDeltaSlice(t, expList, actList, 1E-7,
		"Input:    %v (len=%d)\nExpected: %v (len=%d)\nActual:  %v (len=%d)",
		warmList, len(warmList),
		expList, len(expList),
		actList, len(actList),
	)

	assert.Equal(t, alg.WarmCount(), actWarmCount, "warm count not equal")
}

type TALibQuad0Per func([]float64, []float64, []float64, []float64, []float64) ([]float64, int)
func testTALibQuad0Per(t *testing.T, taAlg TALibQuad0Per, alg AlgQuad) {
	high, low, close, volume := dataHighLowCloseVolume(alg.WarmCount()+3)

	expList, _ := taAlg(high, low, close, volume, nil)

	var actWarmCount int
	var actList []float64
	for i := 0; i < len(high); i++ {
		if vOut := alg.Add(high[i], low[i], close[i], volume[i]); alg.Warmed() {
			actList = append(actList, vOut)
		} else {
			actWarmCount++
		}
	}

	if !alg.Warmed() {
		t.Errorf("algorithm did not warm up within time period")
		t.FailNow()
	}

	assert.InDeltaSlice(t, expList, actList, 1E-7,
		"Expected: %v (len=%d)\nActual:  %v (len=%d)",
		expList, len(expList),
		actList, len(actList),
	)

	assert.Equal(t, alg.WarmCount(), actWarmCount, "warm count not equal")
}

type TALibQuad func([]float64, []float64, []float64, []float64, int, []float64) ([]float64, int)

type TALibQuad2Per func([]float64, []float64, []float64, []float64, int, int, []float64) ([]float64, int)
func testTALibQuad2Per(t *testing.T, inTimePeriodShort, inTimePeriodLong int, taAlg TALibQuad2Per, alg AlgQuad) {
	high, low, close, volume := dataHighLowCloseVolume(alg.WarmCount() + 9)

	expList, _ := taAlg(high, low, close, volume, inTimePeriodShort, inTimePeriodLong, nil)

	var actWarmCount int
	var actList []float64
	for i := 0; i < len(high); i++ {
		if vOut := alg.Add(high[i], low[i], close[i], volume[i]); alg.Warmed() {
			actList = append(actList, vOut)
		} else {
			actWarmCount++
		}
	}

	if !alg.Warmed() {
		t.Errorf("algorithm did not warm up within time period")
		t.FailNow()
	}

	assert.InDeltaSlice(t, expList, actList, 1E-7,
		"Expected: %v (len=%d)\nActual:  %v (len=%d)",
		expList, len(expList),
		actList, len(actList),
	)

	assert.Equal(t, alg.WarmCount(), actWarmCount, "warm count not equal")
}
