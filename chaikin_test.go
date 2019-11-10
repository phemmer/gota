package gota

import (
	"testing"

	"github.com/phemmer/talib"
)

func TestADL(t *testing.T) {
	testTALibQuad0Per(t, talib.Ad, NewADL())
}

func TestADO(t *testing.T) {
	ado := NewADO(3, 10, WarmNone)
	testTALibQuad2Per(t, 3, 10, talib.AdOsc, ado)
}
