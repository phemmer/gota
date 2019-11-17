package gota

type ADL struct {
	prev  float64
}

func NewADL() *ADL {
	return &ADL{}
}

func (adl *ADL) Add(high, low, close, volume float64) float64 {
	var mfv float64
	if high-low == 0 {
		mfv = 0
	} else {
		mfm := ((close - low) - (high - close)) / (high - low)
		mfv = mfm * volume
	}
	cadl := adl.prev + mfv
	adl.prev = cadl
	return cadl
}

func (adl ADL) Warmed() bool {
	return true
}

func (adl ADL) WarmCount() int {
	return 0
}

type ADO struct {
	adl      ADL
	emaShort AlgSingle
	emaLong  AlgSingle
}

func NewADOMA(inTimePeriodShort, inTimePeriodLong int, maNew AlgSingleConstructor) *ADO {
	return &ADO{
		emaShort: maNew(inTimePeriodShort),
		emaLong:  maNew(inTimePeriodLong),
	}
}

/*
type ADOConstructor struct {
	WarmupType WarmupType
}

func (c ADOConstructor) New(inTimePeriodShort, inTimePeriodLong int) AlgQuad {
	return NewADO(inTimePeriodShort, inTimePeriodLong, c.WarmupType)
}
*/

func NewADO(inTimePeriodShort, inTimePeriodLong int, warmupType WarmupType) *ADO {
	return NewADOMA(inTimePeriodShort, inTimePeriodLong, EMAConstructor{warmupType}.New)
}

func (cado *ADO) Add(high, low, close, volume float64) float64 {
	v := cado.adl.Add(high, low, close, volume)
	return cado.emaShort.Add(v) - cado.emaLong.Add(v)
}

func (cado ADO) Warmed() bool {
	return cado.emaLong.Warmed() && cado.emaShort.Warmed()
}

func (cado ADO) WarmCount() int {
	wc := cado.emaLong.WarmCount()
	wc2 := cado.emaShort.WarmCount()
	if wc2 > wc {
		return wc2
	}
	return wc
}
