package gota

import (
	"testing"

	"github.com/phemmer/talib"
	"github.com/stretchr/testify/assert"
)

func TestKER(t *testing.T) {
	list := []float64{20, 21, 22, 23, 22, 21}

	expList := []float64{1, 1.0 / 3, 1.0 / 3}

	ker := NewKER(3)
	var actList []float64
	for _, v := range list {
		if vOut := ker.Add(v); ker.Warmed() {
			actList = append(actList, vOut)
		}
	}

	assert.InDeltaSlice(t, expList, actList, 0.0000001)
}

func TestKAMA(t *testing.T) {
	list := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	expList, _ := talib.Kama(list, 10, nil)

	kama := NewKAMA(10)
	var actList []float64
	for _, v := range list {
		if vOut := kama.Add(v); kama.Warmed() {
			actList = append(actList, vOut)
		}
	}

	assert.InDeltaSlice(t, expList, actList, 0.0000001)
}

func TestKAMAWarmCount(t *testing.T) {
	period := 9
	kama := NewKAMA(period)

	var i int
	for i = 0; i < period*10; i++ {
		kama.Add(float64(i))
		if kama.Warmed() {
			break
		}
	}

	assert.Equal(t, i, kama.WarmCount())
}

var BenchmarkKAMAVal float64

func BenchmarkKAMA(b *testing.B) {
	list := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for n := 0; n < b.N; n++ {
		kama := NewKAMA(5)
		for _, v := range list {
			BenchmarkKAMAVal = kama.Add(v)
		}
	}
}
