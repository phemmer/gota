package gota

import (
	"testing"

	"github.com/phemmer/talib"
	"github.com/stretchr/testify/assert"
)

func TestTRIX(t *testing.T) {
	list := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	expList, _ := talib.Trix(list, 4, nil)

	trix := NewTRIX(4, WarmSMA)
	var actList []float64
	for _, v := range list {
		if vOut := trix.Add(v); trix.Warmed() {
			actList = append(actList, vOut)
		}
	}

	assert.InDeltaSlice(t, expList, actList, 1e-7)
}
