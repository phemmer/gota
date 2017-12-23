package gota

import (
	"testing"

	"github.com/phemmer/talib"
	"github.com/stretchr/testify/assert"
)

func TestEMA(t *testing.T) {
	list := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	expList, _ := talib.Ema(list, 10, nil)

	ema := NewEMA(10, WarmSMA)
	var actList []float64
	for _, v := range list {
		if vOut := ema.Add(v); ema.Warmed() {
			actList = append(actList, vOut)
		}
	}

	assert.InDeltaSlice(t, expList, actList, 0.0000001)
}

func TestDEMA(t *testing.T) {
	list := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	expList, _ := talib.Dema(list, 10, nil)

	dema := NewDEMA(10, WarmSMA)
	var actList []float64
	for _, v := range list {
		if vOut := dema.Add(v); dema.Warmed() {
			actList = append(actList, vOut)
		}
	}

	assert.InDeltaSlice(t, expList, actList, 0.0000001)
}

func TestTEMA(t *testing.T) {
	list := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	expList, _ := talib.Tema(list, 4, nil)

	tema := NewTEMA(4, WarmSMA)
	var actList []float64
	for _, v := range list {
		if vOut := tema.Add(v); tema.Warmed() {
			actList = append(actList, vOut)
		}
	}

	assert.InDeltaSlice(t, expList, actList, 0.0000001)
}

func TestEmaWarmCount(t *testing.T) {
	period := 9
	ema := NewEMA(period, WarmSMA)

	var i int
	for i = 0; i < period*10; i++ {
		ema.Add(float64(i))
		if ema.Warmed() {
			break
		}
	}

	assert.Equal(t, i, ema.WarmCount())
}

func TestDemaWarmCount(t *testing.T) {
	period := 9
	dema := NewDEMA(period, WarmSMA)

	var i int
	for i = 0; i < period*10; i++ {
		dema.Add(float64(i))
		if dema.Warmed() {
			break
		}
	}

	assert.Equal(t, i, dema.WarmCount())
}

func TestTemaWarmCount(t *testing.T) {
	period := 9
	tema := NewTEMA(period, WarmSMA)

	var i int
	for i = 0; i < period*10; i++ {
		tema.Add(float64(i))
		if tema.Warmed() {
			break
		}
	}

	assert.Equal(t, i, tema.WarmCount())
}
