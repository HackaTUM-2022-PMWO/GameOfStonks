package stonks

import (
	"errors"
	http "net/http"
	"sort"
	"time"

	"go.uber.org/zap"
)

func (s *StonksService) update() bool {
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

			// Add a new entry to the users NetWorthTimeSeries-DataPoints
			for userId, user := range s.activeUsers {
				user.mu.Lock()
				user.NetWorthTimeSeries = append(user.NetWorthTimeSeries, DataPoint{
					Time:  user.NetWorthTimeSeries.LatestTime() + 1,
					Value: user.NetWorthTimeSeries.LatestValue(),
				})
				s.activeUsers[userId] = user
				user.mu.Unlock()
			}

			// update the value to the actual new on if there is a match for this stock
			for _, match := range matches {
				stonkName := StonkName(match.Stonk)
				dataPoints := s.prices[stonkName]
				dataPoints[len(dataPoints)-1].Value = (match.SellOrder.Price + match.BuyOrder.Price) / 2.
				s.prices[stonkName] = dataPoints

				// update the users stock position
				for userId, user := range s.activeUsers {
					user.mu.Lock()
					if userId == match.BuyOrder.User.ID { // if buyer
						user.Stonks[stonkName] += match.Quantity
						user.ReservedStonks[stonkName] -= match.Quantity
						user.ReservedMoney -= float64(match.Quantity) * match.BuyOrder.Price
						user.Money -= float64(match.Quantity) * match.BuyOrder.Price
					} else if userId != match.SellOrder.User.ID { // elif seller
						user.Stonks[stonkName] -= match.Quantity
						user.Money += float64(match.Quantity) * match.SellOrder.Price
					} else {
						user.mu.Unlock()
						continue
					}

					// since the stock prices might have been adjusted, we also have to re-evaluate the users net worth
					user.NetWorth = user.Money
					for stonk, num := range user.Stonks {
						user.NetWorth += float64(num) * s.prices[stonk].LatestValue()
					}

					// update the latest NetWorthTimeSeries-DataPoints
					user.NetWorthTimeSeries[len(user.NetWorthTimeSeries)-1].Value = user.NetWorth

					s.activeUsers[userId] = user
					user.mu.Unlock()
				}
			}

			// see if there are more updates
		default:
			if updated {
				return true // FIXME: Trigger SSE with new state
			}
			return false
		}
	}
}

func userExists(r *http.Request, users map[string]*User) (bool, string, error) {
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

func (s *StonksService) startSession() {
	// TODO: Need to clear the users after one round
	if len(s.activeUsers) != 0 {
		s.l.Error("session already active",
			zap.Int("waiting_users_len", len(s.waitingUsers)),
			zap.Int("active_users_len", len(s.activeUsers)),
		)

		// FIXME: Somehow need to handle the case when a session is already
		// 			active and enough people are in the waiting room again
		// return &Err{"other session still active"}
	}

	// make the waitingUsers the active ones
	s.activeUsers = s.waitingUsers

	users := make([]*User, 0, len(s.activeUsers))
	for _, u := range s.activeUsers {
		users = append(users, u)
	}

	// sort the users
	sort.Slice(users, func(i, j int) bool {
		return users[i].Name < users[j].Name
	})

	// make sure all users are up to date
	_ = s.update()

	// Start a timer for the end
	time.AfterFunc(s.roundDuration, func() {
		users := make([]*User, 0, len(s.activeUsers))
		for _, u := range s.activeUsers {
			users = append(users, u)
		}

		// sort the users - highest NetWorth first
		sort.Slice(users, func(i, j int) bool {
			return users[i].NetWorth > users[j].NetWorth
		})

		state := State{
			Start:  nil,
			Reload: false, // the front-end will start the game so no need to reload the current page
			Finish: users,
		}
		s.sseCh <- state

		s.activeUsers = make(map[string]*User, len(s.activeUsers))
	})

	state := State{
		Start:  users,
		Reload: false, // the front-end will start the game so no need to reload the current page
		Finish: nil,
	}
	s.sseCh <- state
}
