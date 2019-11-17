package gota

type AlgSingle interface {
	// Add adds a new sample value to the algorithm and returns the computed value.
	Add(float64) float64
	// Warmed indicates whether the algorithm has enough data to generate accurate results.
	Warmed() bool
	// WarmCount returns the number of samples that must be provided for the algorithm to be fully "warmed".
	WarmCount() int
}

type AlgTriple interface {
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


// AlgSingleConstructor is the signature accepted by various algorithms which accept any moving average, or other AlgSingle, implementation.
type AlgSingleConstructor func(inTimePeriod int) AlgSingle
type AlgTripleConstructor func(inTimePeriod int) AlgTriple
type AlgQuadConstructor func(inTimePeriod int) AlgQuad
type AlgQuad2PerConstructor func(inTimePeriodShort, inTimePeriodLong int) AlgQuad
