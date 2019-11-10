package gota

import (
	"testing"

	"github.com/phemmer/talib"
)

func TestEMA(t *testing.T) {
	ema := NewEMA(10, WarmSMA)
	testTALibSimple(t, 10, ema, talib.Ema)
}

func TestDEMA(t *testing.T) {
	dema := NewDEMA(10, WarmSMA)
	testTALibSimple(t, 10, dema, talib.Dema)
}

func TestTEMA(t *testing.T) {
	tema := NewTEMA(4, WarmSMA)
	testTALibSimple(t, 4, tema, talib.Tema)
}
