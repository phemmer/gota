package gota

import (
	"testing"

	"github.com/phemmer/talib"
)

func TestATR(t *testing.T) {
	atr := NewATR(5)
	testTALibTri(t, 5, talib.Atr, atr)
}
