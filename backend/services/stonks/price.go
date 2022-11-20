package stonks

type Prices map[StonkName]DataPoints

type DataPoints []DataPoint

type DataPoint struct {
	Time  int
	Value float64
}

func (dp DataPoints) LatestTime() int {
	// We initialize the datapoints, so there will be always a value in it, but
	// let's make at least sure we don't get an out-of-bounds exception
	if len(dp) == 0 {
		return 0
	}
	return dp[len(dp)-1].Time
}

func (dp DataPoints) LatestValue() float64 {
	// We initialize the datapoints, so there will be always a value in it, but
	// let's make at least sure we don't get an out-of-bounds exception
	if len(dp) == 0 {
		return 0.0
	}
	return dp[len(dp)-1].Value
}

func NewPrices(stonksPrices map[StonkName]float64) Prices {
	out := make(Prices, len(stonksPrices))
	for name, price := range stonksPrices {
		timeSeries := make([]DataPoint, 1, 1000)
		timeSeries[0] = DataPoint{
			Time:  0,
			Value: price, // TODO: Make configurable
		}
		out[name] = timeSeries
	}
	return out
}
