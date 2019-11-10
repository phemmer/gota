package gota

import (
	"testing"

	"github.com/phemmer/talib"
)

func TestRSI(t *testing.T) {
	rsi := NewRSI(10, WarmSMA)
	testTALibSimple(t, 10, rsi, talib.Rsi)
}
