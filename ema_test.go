package gota

import (
	"testing"

	"github.com/phemmer/talib"
	"github.com/stretchr/testify/assert"
)

func TestEMA(t *testing.T) {
	ema := NewEMA(10, WarmSMA)
	testTALibSimple(t, 10, ema, talib.Ema)
}

func TestEMAPreWarmSMA(t *testing.T) {
	ema := NewEMA(5, WarmSMA)
	ema.Add(1)
	ema.Add(1)
	ema.Add(1)
	v := ema.Add(0)
	assert.InDelta(t, 0.75, v, 1E-7)
}

func TestDEMA(t *testing.T) {
	dema := NewDEMA(10, WarmSMA)
	testTALibSimple(t, 10, dema, talib.Dema)
}

func TestTEMA(t *testing.T) {
	tema := NewTEMA(4, WarmSMA)
	testTALibSimple(t, 4, tema, talib.Tema)
}
