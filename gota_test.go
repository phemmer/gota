package gota

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func randList(size int, min, max float64) []float64 {
	list := make([]float64, size)
	for i := 0; i < size; i++ {
		list[i] = math.Floor(rand.Float64() * (max - min) + min)
	}
	return list
}

func dataHighLowCloseVolume(size int) ([]float64, []float64, []float64, []float64) {
	low := randList(size, 1, 20)

	high := make([]float64, len(low))
	for i := 0; i < size; i++ {
		high[i] = low[i] + math.Floor(rand.Float64() * 5)
	}

	close := make([]float64, size)
	for i := 0; i < size; i++ {
		close[i] = rand.Float64() * (high[i] - low[i]) + low[i]
	}

	volume := randList(size, 1, 5)

	return high, low, close, volume
}

type TALibQuad func([]float64, []float64, []float64, []float64, []float64) ([]float64, int)
func testTALibQuad(t *testing.T, taAlg TALibQuad, alg AlgQuad) {
	high, low, close, volume := dataHighLowCloseVolume(20)

	expList, _ := taAlg(high, low, close, volume, nil)

	var actList []float64
	for i := 0; i < len(high); i++ {
		if vOut := alg.Add(high[i], low[i], close[i], volume[i]); alg.Warmed() {
			actList = append(actList, vOut)
		}
	}

	assert.InDeltaSlice(t, expList, actList, 1E-7)
}

type TALibPerQuad func([]float64, []float64, []float64, []float64, int, []float64) ([]float64, int)

type TALibSingle func([]float64, []float64) ([]float64, int)

type TALibPerSingle func([]float64, int, []float64) ([]float64, int)
func testTALibPerSingle(t *testing.T, inTimePeriod int, alg AlgSimple, taAlg TALibPerSingle) {
	list, _, _, _ := dataHighLowCloseVolume(alg.WarmCount() + 3)

	expList, _ := taAlg(list, inTimePeriod, nil)

	var actWarmCount int
	var actList []float64
	for i := 0; i < len(list); i++ {
		if vOut := alg.Add(list[i]); alg.Warmed() {
			actList = append(actList, vOut)
		} else {
			actWarmCount++
		}
	}

	if !alg.Warmed() {
		t.Errorf("algorithm did not warm up within time period")
		t.FailNow()
	}

	assert.InDeltaSlice(t, expList, actList, 1E-7)

	assert.Equal(t, alg.WarmCount(), actWarmCount, "warm count not equal")
}
