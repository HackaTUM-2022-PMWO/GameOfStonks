package stonks

import (
	"errors"
	http "net/http"
	"sort"
	time "time"

	"github.com/google/uuid"
	"github.com/hackaTUM/GameOfStonks/store"
	"go.uber.org/zap"
)

type Err struct {
	Message string `json:"message"`
}

func (e *Err) Error() string {
	return e.Message
}

// type ScalarError string

// func (se *ScalarError) Error() string {
// 	return string(*se)
// }

// type ScalarInPlace string

type StonksService struct {
	l *zap.Logger

	// configuration values
	//
	startMoney float64

	// Time series data for all the stonks
	prices Prices

	orderP store.OrderPersistor
	matchP store.MatchPersistor

	matchUpdateCh <-chan []*store.Match

	// users are only held ephemeraly
	waitingUsers map[string]User
	activeUsers  map[string]User
}

func NewStonksService(
	l *zap.Logger,
	initialStonkPrices map[StonkName]float64,
	startMoney float64,
	orderP store.OrderPersistor,
	matchP store.MatchPersistor,
	matchUpdateCh <-chan []*store.Match,
) *StonksService {
	return &StonksService{
		l:             l.With(zap.String("component", "service")),
		prices:        NewPrices(initialStonkPrices),
		startMoney:    startMoney,
		orderP:        orderP,
		matchP:        matchP,
		matchUpdateCh: matchUpdateCh,

		waitingUsers: make(map[string]User, 5),
		activeUsers:  make(map[string]User, 5),
	}
}

type User struct {
	id string // NOTE: private on purpose
	// TODO: Probably need to add the ID without leaking it to other users (impersenation!)
	Name string

	// TODO: Deduct the money once an order is placed not when it is executed!
	Money float64

	// TODO: Need to create a users NetWorth (i.e. money current values of stonks)
}

type Match struct {
	UserSell  string
	UserBuy   string
	Quantity  int
	TimeStamp int64
}

type Order struct {
	UserName  string
	OrderType OrderType
	Quantity  int
	TimeStamp int64
}

type OrderType string

const (
	OrderTypeSell = "sell"
	OrderTypeBuy  = "buy"
)

type StonkInfo struct {
	Name         StonkName
	TimeSeries   []DataPoint
	MatchHistory []Match
	UserOrders   []Order
	Orders       []Order
}

type PlaceOrderCmd struct {
	Stonk     StonkName
	Quantity  int
	Price     float64
	OrderType OrderType
}

func (s *StonksService) NewUser(w http.ResponseWriter, r *http.Request, name string) ([]User, *Err) {
	if r.Method != http.MethodPost {
		return nil, &Err{"you gotta post wlad"}
	}

	if exists, userId, err := userExists(r, s.waitingUsers); err != nil {
		s.l.Error("unable to read user cookie", zap.Error(err))
		return nil, &Err{"unable to read user cookie"}
	} else if exists {
		s.l.Error("user already exists", zap.String("user_id", userId))
		return nil, &Err{"user already registered"}
	}

	u := User{
		id:    uuid.New().String(),
		Name:  name,
		Money: s.startMoney,
	}
	s.waitingUsers[u.id] = u

	// Set a cookie
	cookie := &http.Cookie{Name: "user", Value: u.id, Expires: time.Now().Add(time.Hour * 24 * 7)}
	http.SetCookie(w, cookie)

	users := make([]User, 0, len(s.waitingUsers))
	for _, u := range s.waitingUsers {
		users = append(users, u)
	}

	// sort the users
	sort.Slice(users, func(i, j int) bool {
		return users[i].Name < users[j].Name
	})
	return users, nil
}

// TODO: Actually this should be an SSE
func (s *StonksService) StartSession(w http.ResponseWriter, r *http.Request, id string) ([]User, *Err) {
	if r.Method != http.MethodPost {
		return nil, &Err{"you gotta post wlad"}
	}

	// TODO: Need to clear the users after one round
	if len(s.activeUsers) != 0 {
		s.l.Error("session already active",
			zap.Int("waiting_users_len", len(s.waitingUsers)),
			zap.Int("active_users_len", len(s.activeUsers)),
		)
		return nil, &Err{"other session still active"}
	}

	// make the waitingUsers the active ones
	s.activeUsers = s.waitingUsers

	users := make([]User, 0, len(s.activeUsers))
	for _, u := range s.activeUsers {
		users = append(users, u)
	}

	// sort the users
	sort.Slice(users, func(i, j int) bool {
		return users[i].Name < users[j].Name
	})
	return users, nil
}

func (s *StonksService) GetStonkInfo(w http.ResponseWriter, r *http.Request, stonk StonkName) (StonkInfo, *Err) {
	if r.Method != http.MethodPost {
		return StonkInfo{}, &Err{"you gotta post wlad"}
	}

	// verify the user
	exists, userId, err := userExists(r, s.activeUsers)
	if err != nil {
		s.l.Warn("unable to read user cookie", zap.Error(err))
		// only warn - we still can return most of the result
	} else if !exists {
		s.l.Warn("user is not an active user", zap.String("user_id", userId))
		// only warn - we still can return most of the result
	}

	// make sure the stonk is valid and actually set
	if !stonk.IsValid() {
		s.l.Error("invalid stonk name", zap.String("stonk", string(stonk)))
		return StonkInfo{}, &Err{"invalid stonk name"}
	}

	// Get the data from the collections
	storeOrders, err := s.orderP.GetOrders(r.Context(), string(stonk), nil)
	if err != nil {
		s.l.Error("unable to retrieve orders", zap.String("stonk", string(stonk)), zap.Error(err))
		return StonkInfo{}, &Err{"unable to retrieve orders"}
	}

	userStoreOrders := make([]*store.Order, 0, len(storeOrders))
	if userId != "" {
		newStoreOrders := make([]*store.Order, 0, len(storeOrders))
		for _, o := range storeOrders {
			if o.User.ID == userId {
				userStoreOrders = append(userStoreOrders, o)
			} else {
				newStoreOrders = append(newStoreOrders, o)
			}
		}
		storeOrders = newStoreOrders
	}

	// transform the orders
	orders := ordersToStonksVo(storeOrders)
	userOrders := ordersToStonksVo(userStoreOrders)
	// sort orders (newest first)
	sort.Slice(orders, func(i, j int) bool {
		return orders[i].TimeStamp > orders[j].TimeStamp
	})
	sort.Slice(userOrders, func(i, j int) bool {
		return userOrders[i].TimeStamp > userOrders[j].TimeStamp
	})

	storeMatches, err := s.matchP.GetMatches(r.Context(), string(stonk), nil)
	if err != nil {
		s.l.Error("unable to retrieve orders", zap.String("stonk", string(stonk)), zap.Error(err))
		return StonkInfo{}, &Err{"unable to retrieve orders"}
	}

	// transform the orders
	matches := matchsToStonksVo(storeMatches)
	// sort matches (newest first)
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].TimeStamp > matches[j].TimeStamp
	})

	// update the prices before we retrieve them
	err = s.update()
	if err != nil {
		s.l.Error("unable to update", zap.Error(err))
		return StonkInfo{}, &Err{"unable to update"}
	}

	ts, ok := s.prices[stonk]
	if !ok {
		s.l.Error("no time series found", zap.String("stonk", string(stonk)))
		return StonkInfo{}, &Err{"time series not found"}
	}

	return StonkInfo{
		TimeSeries:   ts,
		Orders:       orders,
		UserOrders:   userOrders,
		MatchHistory: matches,
	}, nil
}

func (s *StonksService) PlaceOrder(w http.ResponseWriter, r *http.Request, cmd PlaceOrderCmd) *Err {
	if r.Method != http.MethodPost {
		return &Err{"you gotta post wlad"}
	}

	// verify the user
	exists, userId, err := userExists(r, s.activeUsers)
	if err != nil {
		s.l.Error("unable to read user cookie", zap.Error(err))
		return &Err{"unable to read user cookie"}
	} else if !exists {
		s.l.Error("user is not an active user", zap.String("user_id", userId))
		return &Err{"user is not an active user"}
	}

	// Make sure the stonk exists
	if !cmd.Stonk.IsValid() {
		s.l.Error("user is not an active user",
			zap.String("user_id", userId),
			zap.Float64("price", cmd.Price),
		)
		return &Err{"invalid stonk"}
	}

	// make sure the price is not negative
	if cmd.Price < 0. {
		return &Err{"negative price"}
	}

	user, ok := s.activeUsers[userId]
	if !ok {
		s.l.Error("user is not an active user", zap.String("user_id", userId))
		return &Err{"user is not an active user"}
	}
	if user.Money < (cmd.Price * float64(cmd.Quantity)) {
		s.l.Error("user has insufficient funds",
			zap.String("stonk", string(cmd.Stonk)),
			zap.Float64("price", cmd.Price),
			zap.Int("quantity", cmd.Quantity),
			zap.Error(err),
		)
		return &Err{"user has insufficient funds"} // TODO: Create separate error
	}

	// create a store order object
	order := store.Order{
		Id:       uuid.New().String(),
		Stonk:    string(cmd.Stonk),
		Quantity: cmd.Quantity,
		Price:    cmd.Price,
		Type:     orderTypeToStore(cmd.OrderType),
		User: store.User{
			ID:   user.id,
			Name: user.Name,
		},
		Time: time.Now(),
	}

	// insert the order
	err = s.orderP.InsertOrder(r.Context(), order)
	if err != nil {
		s.l.Error("unable to insert order", zap.Error(err))
		return &Err{"unable to insert order"}
	}

	return nil
}

type UpdateOrderCmd struct {
	Id       string
	Quantity int
	Price    float64
}

// FIXME: Implement
// NOTE: Update order with a quantity of 0 deletes the order
// func (s *StonksService) UpdateOrder(w http.ResponseWriter, r *http.Request, cmd UpdateOrderCmd) *Err {
// 	if r.Method != http.MethodPost {
// 		return &Err{"you gotta post wlad"}
// 	}

// 	// verify the user
// 	exists, userId, err := userExists(r, s.activeUsers)
// 	if err != nil {
// 		s.l.Error("unable to read user cookie", zap.Error(err))
// 		return &Err{"unable to read user cookie"}
// 	} else if !exists {
// 		s.l.Error("user is not an active user", zap.String("user_id", userId))
// 		return &Err{"user is not an active user"}
// 	}

// 	// Make sure the stonk exists
// 	if !cmd.Stonk.IsValid() {
// 		s.l.Error("user is not an active user",
// 			zap.String("user_id", userId),
// 			zap.Float64("price", cmd.Price),
// 		)
// 		return &Err{"invalid stonk"}
// 	}

// 	// make sure the price is not negative
// 	if cmd.Price < 0. {
// 		return &Err{"negative price"}
// 	}

// 	user, ok := s.activeUsers[userId]
// 	if !ok {
// 		s.l.Error("user is not an active user", zap.String("user_id", userId))
// 		return &Err{"user is not an active user"}
// 	}
// 	if user.Money < (cmd.Price * float64(cmd.Quantity)) {
// 		s.l.Error("user has insufficient funds",
// 			zap.String("stonk", string(cmd.Stonk)),
// 			zap.Float64("price", cmd.Price),
// 			zap.Int("quantity", cmd.Quantity),
// 			zap.Error(err),
// 		)
// 		return &Err{"user has insufficient funds"} // TODO: Create separate error
// 	}

// 	// create a store order object
// 	order := store.Order{
// 		Id:       uuid.New().String(),
// 		Stonk:    string(cmd.Stonk),
// 		Quantity: cmd.Quantity,
// 		Price:    cmd.Price,
// 		Type:     orderTypeToStore(cmd.OrderType),
// 		User: store.User{
// 			ID:   user.id,
// 			Name: user.Name,
// 		},
// 		Time: time.Now(),
// 	}

// 	// insert the order
// 	err = s.orderP.InsertOrder(r.Context(), order)
// 	if err != nil {
// 		s.l.Error("unable to insert order", zap.Error(err))
// 		return &Err{"unable to insert order"}
// 	}

// 	return nil
// }

// TODO: Add functions for:
// - UpdateOrder
// - DeleteOrder
// - GetUserInfo (users current portfolie + others)

// TODO: Add SSE for:
// - StartSession(?)
// - Order has been matched
// - NewGameState (how the fuck?)
// - SessionFinished (need to include leaderboard)

//---------------------------------------------------------------------------
// ~ utils
//---------------------------------------------------------------------------

// TODO: Need to update player NetWorth
func (s *StonksService) update() error {
	// drain the updates chanel until it is empty
	for {
		select {
		case matches := <-s.matchUpdateCh:
			// update the stonk prices
			time := make(map[StonkName]int, len(s.prices))
			for stonkName, stonkPrices := range s.prices {
				time[stonkName] = stonkPrices[len(stonkPrices)-1].Time
			}
			for _, match := range matches {
				stonkName := StonkName(match.Stonk)
				s.prices[stonkName] = append(s.prices[stonkName], DataPoint{
					Time:  time[stonkName],
					Value: (match.SellOrder.Price + match.BuyOrder.Price) / 2.,
				})
			}

			// FIXME: Adapt player NetWorth
			// s.activeUsers[]

			// see if there are more updates
		default:
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
