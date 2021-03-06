package gota

import (
	"testing"

	"github.com/phemmer/talib"
	"github.com/stretchr/testify/assert"
)

func TestCMO(t *testing.T) {
	list := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	expList := []float64{100, 100, 100, 100, 100, 80, 60, 40, 20, 0, -20, -40, -60, -80, -100, -100, -100, -100, -100}

	cmo := NewCMO(10)
	var actList []float64
	for _, v := range list {
		if vOut := cmo.Add(v); cmo.Warmed() {
			actList = append(actList, vOut)
		}
	}

	assert.Equal(t, 11, cmo.WarmCount())
	assert.Equal(t, cmo.WarmCount(), len(list) - len(actList) + 1)

	assert.InDeltaSlice(t, expList, actList, 1E-7)
}

func TestCMOS(t *testing.T) {
	cmo := NewCMOS(10, WarmSMA)
	testTALibSimple(t, 10, cmo, talib.Cmo)
}
