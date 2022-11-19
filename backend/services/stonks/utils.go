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
			// set the previous value as a default, just to make sure the clock advances for the stock
			for stonkName, stonkPrices := range s.prices {
				time := stonkPrices.LatestTime() + 1
				value := stonkPrices.LatestValue()
				s.prices[stonkName] = append(s.prices[stonkName], DataPoint{
					Time:  time,
					Value: value,
				})
			}
			// update the value to the actual new on if there is a match for this stock
			for _, match := range matches {
				stonkName := StonkName(match.Stonk)
				dataPoints := s.prices[stonkName]
				dataPoints[len(dataPoints)-1].Value = (match.SellOrder.Price + match.BuyOrder.Price) / 2.
				s.prices[stonkName] = dataPoints

				// update the users stock position
				for userId, user := range s.activeUsers {
					if userId == match.BuyOrder.User.ID {				// if buyer
						user.Stonks[stonkName] += match.Quantity
						user.ReservedStonks[stonkName] -= match.Quantity
						user.ReservedMoney += float64(match.Quantity) * s.prices[stonkName].LatestValue()
					} else if userId != match.SellOrder.User.ID {		// elif seller
						user.Stonks[stonkName] -= 1
					} else {
						continue
					}

					// since the stock prices might have been adjusted, we also have to re-evaluate the users net worth
					user.NetWorth = user.Money
					for stonk, num := range user.Stonks {
						user.NetWorth += float64(num) * s.prices[stonk].LatestValue()
					}

					// update the NetWorthTimeSeries-DataPoints
					nextD := DataPoint{s.prices[stonkName].LatestTime(), user.NetWorth}
					user.NetWorthTimeSeries = append(user.NetWorthTimeSeries, nextD)

					s.activeUsers[userId] = user
				}
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
