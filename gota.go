package gota

type AlgSimple interface {
	// Add adds a new sample value to the algorithm and returns the computed value.
	Add(float64) float64
	// Warmed indicates whether the algorithm has enough data to generate accurate results.
	Warmed() bool
	// WarmCount returns the number of samples that must be provided for the algorithm to be fully "warmed".
	WarmCount() int
}

type AlgTri interface {
	// Add adds a new sample value to the algorithm and returns the computed value.
	// Commonly the parameters are (high, low, close).
	Add(float64, float64, float64) float64
	// Warmed indicates whether the algorithm has enough data to generate accurate results.
	Warmed() bool
	// WarmCount returns the number of samples that must be provided for the algorithm to be fully "warmed".
	WarmCount() int
}

type AlgQuad interface {
	// Add adds a new sample value to the algorithm and returns the computed value.
	// Commonly the parameters are (high, low, close, volume).
	Add(float64, float64, float64, float64) float64
	// Warmed indicates whether the algorithm has enough data to generate accurate results.
	Warmed() bool
	// WarmCount returns the number of samples that must be provided for the algorithm to be fully "warmed".
	WarmCount() int
}


// AlgSimpleConstructor is the signature accepted by various algorithms which accept any moving average, or other AlgSimple, implementation.
type AlgSimpleConstructor func(inTimePeriod int) AlgSimple
type AlgTriConstructor func(inTimePeriod int) AlgTri
type AlgQuadConstructor func(inTimePeriod int) AlgQuad
type AlgQuad2PerConstructor func(inTimePeriodShort, inTimePeriodLong int) AlgQuad
