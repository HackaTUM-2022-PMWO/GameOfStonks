package stonks

import (
	"errors"
	http "net/http"
)

// TODO: Need to update player NetWorth
func (s *StonksService) update() error {
	// drain the updates chanel until it is empty
	updated := false
	for {
		select {
		case matches := <-s.matchUpdateCh:
			// update the stonk prices
			time := make(map[StonkName]int, len(s.prices))
			for stonkName, stonkPrices := range s.prices {
				time[stonkName] = stonkPrices.LatestTime()
			}
			for _, match := range matches {
				stonkName := StonkName(match.Stonk)
				s.prices[stonkName] = append(s.prices[stonkName], DataPoint{
					Time:  time[stonkName],
					Value: (match.SellOrder.Price + match.BuyOrder.Price) / 2.,
				})
			}

			// update the users stock position
			for userId, user := range s.activeUsers {
				// FIXME: update the users stock position

				// since the stock prices might have been adjusted, we also have to re-evaluate the users net worth
				// FIXME: Adapt user's NetWorth (Money-ReservedMoney + Stonks_Quantity*Stonk_LastPrice)

				// update the datapoints
				// FIXME: update the users NetWorthTimeSeries-DataPoints

				s.activeUsers[userId] = user
			}

			updated = true
			// see if there are more updates
		default:
			if updated {
				// FIXME: Trigger SSE with new state
			}
			return nil
		}
	}
}

func userExists(r *http.Request, users map[string]User) (bool, string, error) {
	cookie, err := r.Cookie("user")
	if errors.Is(err, http.ErrNoCookie) {
		// nothing to do
	} else if err != nil {
		return false, "", err
	} else {
		// try to find the user by the id
		if _, ok := users[cookie.Value]; ok {
			return true, cookie.Value, nil
		}
	}

	return false, "", nil

}
