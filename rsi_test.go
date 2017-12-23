package gota

import (
	"testing"

	"github.com/phemmer/talib"
	"github.com/stretchr/testify/assert"
)

func TestRSI(t *testing.T) {
	list := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	expList, _ := talib.Rsi(list, 10, nil)

	rsi := NewRSI(10, WarmSMA)
	var actList []float64
	for _, v := range list {
		if vOut := rsi.Add(v); rsi.Warmed() {
			actList = append(actList, vOut)
		}
	}

	assert.InDeltaSlice(t, expList, actList, 0.0000001)
}
