package gota

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWSMA(t *testing.T) {
	list := []float64{1,2,3,4,5,6}

	var actList []float64
	wsma := NewWSMA(3, WarmSMA)
	for _, v := range list {
		if outV := wsma.Add(v); wsma.Warmed() {
			actList = append(actList, outV)
		}
	}

	expList := []float64{2,2.6666667,3.4444444,4.2962963}

	assert.InDeltaSlice(t, expList, actList, 1E-7, "Expected: %v\nActual: %v", expList, actList)
}
