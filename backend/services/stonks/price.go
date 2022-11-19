package stonks

type Prices map[string][]DataPoint

type DataPoint struct {
	Time  int
	Value float64
}

func NewPrices(stonksPrices map[string]float64) Prices {
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

// TODO: Needs to be called when the matcher returned a new match
// TODO: Prices can not go below 0
func updatePrice() {

}

// TODO
