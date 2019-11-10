package gota

// WSMAConstructor is for passing to algorithms with swappable moving average implementations.
type WSMAConstructor struct {
	WarmupType WarmupType
}
func (c WSMAConstructor) New(inTimePeriod int) AlgSimple {
	return NewWSMA(inTimePeriod, c.WarmupType)
}

// Wilder's Smoothing - Moving Average
func NewWSMA(inTimePeriod int, warmType WarmupType) *EMA {
	// http://etfhq.com/blog/2010/08/19/wilders-smoothing/
	// Documentation has Wilder's Smoothing calculated very different from an EMA. However the calculation is effectively
	// the same, but with a different alpha.
	ema := NewEMA(inTimePeriod, warmType)
	ema.alpha = float64(1) / float64(inTimePeriod)
	return ema
}

/*
type WSMA struct {
	prev float64
	count int
	warmCount int
}

func NewWSMA(inTimePeriod int) *WSMA {
	return &WSMA{
		warmCount: inTimePeriod,
	}
}

func (wsma *WSMA) Add(v float64) float64 {
	if wsma.count < wsma.warmCount {
		wsma.prev = ((wsma.prev * float64(wsma.count)) + v) / float64(wsma.count + 1)
		wsma.count++
		return wsma.prev
	}

	wsma.prev = (wsma.prev * float64(wsma.count) - wsma.prev + v) / float64(wsma.count)
	return wsma.prev
}

func (wsma WSMA) Last() float64 {
	return wsma.prev
}

func (wsma WSMA) Warmed() bool {
	return wsma.count == wsma.warmCount
}

func (wsma WSMA) WarmCount() int {
	return wsma.warmCount
}
*/
