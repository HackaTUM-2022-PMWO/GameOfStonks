package stonks

type Prices map[StonkName][]DataPoint

type DataPoint struct {
	Time  int
	Value float64
}

func NewPrices(stonksPrices map[StonkName]float64) Prices {
	out := make(Prices, len(stonksPrices))
	for name, price := range stonksPrices {
		timeSeries := make([]DataPoint, 1000)
		timeSeries[0] = DataPoint{
			Time:  0,
			Value: price, // TODO: Make configurable
		}
		out[name] = timeSeries
	}
	return out
}
