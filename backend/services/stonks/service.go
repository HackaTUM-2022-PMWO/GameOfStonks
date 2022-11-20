package stonks

import (
	http "net/http"
	"sort"
	"sync"
	time "time"

	"github.com/google/uuid"
	"github.com/hackaTUM/GameOfStonks/bot"
	"github.com/hackaTUM/GameOfStonks/store"
	"go.uber.org/zap"
)

type Err struct {
	Message string `json:"message"`
}

func (e *Err) Error() string {
	return e.Message
}

type StonksService struct {
	l *zap.Logger

	// configuration values
	//
	startMoney    float64
	startStonks   map[StonkName]int
	roundDuration time.Duration

	// Time series data for all the stonks
	prices Prices

	orderP store.OrderPersistor
	matchP store.MatchPersistor

	// bots that influence the market
	bots []bot.Bot

	matchUpdateCh <-chan []*store.Match

	sseCh chan State

	// users are only held ephemeraly
	waitingUsers map[string]*User
	activeUsers  map[string]*User
}

func NewStonksService(
	l *zap.Logger,
	initialStonkPrices map[StonkName]float64,
	startMoney float64,
	startStonks map[StonkName]int,
	roundDuration time.Duration,
	orderP store.OrderPersistor,
	matchP store.MatchPersistor,
	bots []bot.Bot,
	matchUpdateCh <-chan []*store.Match,
	sseCh chan State,
) *StonksService {
	service := &StonksService{
		l:             l.With(zap.String("component", "service")),
		prices:        NewPrices(initialStonkPrices),
		startMoney:    startMoney,
		startStonks:   startStonks,
		roundDuration: roundDuration,
		orderP:        orderP,
		matchP:        matchP,
		bots:          bots,
		matchUpdateCh: matchUpdateCh,
		sseCh:         sseCh,

		waitingUsers: make(map[string]*User, 5),
		activeUsers:  make(map[string]*User, 5),
	}

	return service
}

type State struct {
	Start         []*User       `json:"start,omitempty"`
	RoundDuration time.Duration `json:"roundDuration,omitempty"`
	Reload        bool          `json:"reload"`
	Finish        []*User       `json:"finish,omitempty"`
}

type User struct {
	mu sync.Mutex

	id   string // NOTE: private on purpose
	Name string

	Money         float64
	ReservedMoney float64

	Stonks         map[StonkName]int
	ReservedStonks map[StonkName]int

	NetWorth           float64
	NetWorthTimeSeries DataPoints
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
	Price     float64
	TimeStamp int64
}

type OrderType string

const (
	OrderTypeSell OrderType = "sell"
	OrderTypeBuy  OrderType = "buy"
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

type UpdateOrderCmd struct {
	Id       string
	Quantity int
	Price    float64
}

func (s *StonksService) NewUser(w http.ResponseWriter, r *http.Request, name string) ([]*User, *Err) {
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

	u := &User{
		id:                 uuid.New().String(),
		Name:               name,
		Money:              s.startMoney,
		Stonks:             make(map[StonkName]int, len(AllStonkNames)),
		ReservedStonks:     make(map[StonkName]int, len(AllStonkNames)),
		NetWorthTimeSeries: make(DataPoints, 0, 1000),
	}

	// set the number of starting stocks
	for s, i := range s.startStonks {
		u.Stonks[s] = i
	}

	// initialize the networth
	u.NetWorth = u.Money
	for stonk, num := range u.Stonks {
		u.NetWorth += float64(num) * s.prices[stonk].LatestValue()
	}

	// update the latest NetWorthTimeSeries-DataPoints
	u.NetWorthTimeSeries = append(u.NetWorthTimeSeries, DataPoint{
		Time:  0,
		Value: u.NetWorth,
	})

	s.waitingUsers[u.id] = u

	// Set a cookie
	cookie := &http.Cookie{
		Name:    "user",
		Value:   u.id,
		Expires: time.Now().Add(time.Hour * 24 * 7),
		Path:    "/",
	}
	http.SetCookie(w, cookie)

	users := make([]*User, 0, len(s.waitingUsers))
	for _, u := range s.waitingUsers {
		users = append(users, u)
	}

	// sort the users
	sort.Slice(users, func(i, j int) bool {
		return users[i].Name < users[j].Name
	})

	// TODO: Maybe change the condition before the presentation
	if len(s.waitingUsers) >= 2 {
		go func() {
			time.Sleep(500 * time.Millisecond)
			s.startSession()
		}()
	}

	return users, nil
}

func (s *StonksService) GetUserInfo(w http.ResponseWriter, r *http.Request) (*User, []*User, *Err) {
	if r.Method != http.MethodPost {
		return nil, nil, &Err{"you gotta post wlad"}
	}

	exists, userId, err := userExists(r, s.activeUsers)
	if err != nil {
		s.l.Error("unable to read user cookie", zap.Error(err))
		return nil, nil, &Err{"unable to read user cookie"}
	} else if !exists {
		s.l.Error("user is not an active user", zap.String("user_id", userId))
		return nil, nil, &Err{"user is not an active user"}
	}

	user, ok := s.activeUsers[userId]
	if !ok {
		s.l.Error("user is not an active user", zap.String("user_id", userId))
		return nil, nil, &Err{"user is not an active user"}
	}

	otherUsers := make([]*User, 0, len(s.activeUsers)-1)
	for _, u := range s.waitingUsers {
		if u.id != userId {
			otherUsers = append(otherUsers, u)
		}
	}

	// sort the users
	sort.Slice(otherUsers, func(i, j int) bool {
		return otherUsers[i].Name < otherUsers[j].Name
	})
	return user, otherUsers, nil
}

func (s *StonksService) GetStonkInfo(w http.ResponseWriter, r *http.Request, stonk StonkName) (StonkInfo, *Err) {
	if r.Method != http.MethodPost {
		return StonkInfo{}, &Err{"you gotta post wlad"}
	}

	// verify the user
	exists, userId, err := userExists(r, s.activeUsers)
	if err != nil {
		s.l.Warn("unable to read user cookie", zap.Error(err))
		return StonkInfo{}, &Err{"unable to read user cookie"}
	} else if !exists {
		s.l.Warn("user is not an active user", zap.String("user_id", userId))
		return StonkInfo{}, &Err{"user is not an active user"}
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
	newStoreOrders := make([]*store.Order, 0, len(storeOrders))
	for _, o := range storeOrders {
		if o.User.ID == userId {
			userStoreOrders = append(userStoreOrders, o)
		} else {
			newStoreOrders = append(newStoreOrders, o)
		}
	}
	storeOrders = newStoreOrders

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

	storeMatches, err := s.matchP.GetMatches(r.Context(), string(stonk))
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
	updated := s.update()
	if updated {
		state := State{
			Start:  nil,
			Reload: true,
			Finish: nil,
		}

		s.sseCh <- state
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

	// make sure the price is not non-positive
	if cmd.Price <= 0. {
		s.l.Error("order creation with non-positive price",
			zap.String("user_id", userId),
			zap.Int("quantity", cmd.Quantity),
			zap.Float64("price", cmd.Price),
		)
		return &Err{"non-positive price"}
	}

	// make sure the price is not non-positive
	if cmd.Quantity <= 0 {
		s.l.Error("order creation with non-positive quantity",
			zap.String("user_id", userId),
			zap.Int("quantity", cmd.Quantity),
			zap.Float64("price", cmd.Price),
		)
		return &Err{"non-positive quantity"}
	}

	user, ok := s.activeUsers[userId]
	if !ok {
		s.l.Error("user is not an active user", zap.String("user_id", userId))
		return &Err{"user is not an active user"}
	}
	// lock the user
	user.mu.Lock()
	defer user.mu.Unlock()

	totalPrice := (cmd.Price * float64(cmd.Quantity))
	if cmd.OrderType == OrderTypeBuy {
		// make sure the user has sufficient funds
		if (user.Money - user.ReservedMoney) < totalPrice {
			s.l.Error("user has insufficient funds",
				zap.String("stonk", string(cmd.Stonk)),
				zap.Float64("price", cmd.Price),
				zap.Int("quantity", cmd.Quantity),
				zap.Error(err),
			)
			return &Err{"user has insufficient funds"}
		}
	} else if cmd.OrderType == OrderTypeSell {
		// make sure the user has amount of the stock
		if _, ok := user.ReservedStonks[cmd.Stonk]; !ok {
			// init if not yet already intialized
			user.ReservedStonks[cmd.Stonk] = 0
		}

		if (user.Stonks[cmd.Stonk] - user.ReservedStonks[cmd.Stonk]) < cmd.Quantity {
			s.l.Error("user has insufficient stocks",
				zap.String("stonk", string(cmd.Stonk)),
				zap.Int("user_stonks", user.Stonks[cmd.Stonk]),
				zap.Int("user_reserved_stonks", user.ReservedStonks[cmd.Stonk]),
				zap.Int("quantity", cmd.Quantity),
				zap.Error(err),
			)
			return &Err{"user has insufficient stocks"}
		}

	} else {
		s.l.Error("unknown OrderType", zap.String("order_type", string(cmd.OrderType)))
		return &Err{"unknown OrderType"}
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

	if cmd.OrderType == OrderTypeBuy {
		// increase the reserved money of the user
		user.ReservedMoney = user.ReservedMoney + totalPrice
		s.activeUsers[userId] = user
	} else if cmd.OrderType == OrderTypeSell {
		// increase the reserved stocks of the user
		user.ReservedStonks[cmd.Stonk] += cmd.Quantity
		s.activeUsers[userId] = user

	} else {
		s.l.Error("unknown OrderType", zap.String("order_type", string(cmd.OrderType)))
		return &Err{"unknown OrderType"}
	}

	return nil
}

// NOTE: Update order with a quantity of 0 deletes the order
func (s *StonksService) UpdateOrder(w http.ResponseWriter, r *http.Request, cmd UpdateOrderCmd) *Err {
	if r.Method != http.MethodPost {
		return &Err{"you gotta post wlad"}
	}

	// make sure the id is not empty
	if cmd.Id == "" {
		s.l.Error("empty orderId")
		return &Err{"empty orderId"}
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

	// make sure the price is not non-positive
	if cmd.Price <= 0. {
		s.l.Error("order update with non-positive price",
			zap.String("user_id", userId),
			zap.String("order_id", cmd.Id),
			zap.Int("quantity", cmd.Quantity),
			zap.Float64("price", cmd.Price),
		)
		return &Err{"non-positive price"}
	}

	// make sure the price is not negative
	if cmd.Quantity < 0 {
		s.l.Error("order update with negative quantity",
			zap.String("user_id", userId),
			zap.String("order_id", cmd.Id),
			zap.Int("quantity", cmd.Quantity),
			zap.Float64("price", cmd.Price),
		)
		return &Err{"negative quantity"}
	}

	user, ok := s.activeUsers[userId]
	if !ok {
		s.l.Error("user is not an active user", zap.String("user_id", userId))
		return &Err{"user is not an active user"}
	}
	// lock the user
	user.mu.Lock()
	defer user.mu.Unlock()

	// Get the order
	order, err := s.orderP.GetOrder(r.Context(), cmd.Id)
	if err != nil {
		s.l.Error("unable to get previous order",
			zap.String("order_id", cmd.Id),
			zap.Error(err),
		)
		return &Err{"unable to get previous order"}
	}

	// if the quantity is set to 0 we are actually deleting the order
	if cmd.Quantity == 0 {
		s.l.Info("deleting order", zap.String("order_id", cmd.Id))

		err := s.orderP.DeleteOrder(r.Context(), cmd.Id)
		if err != nil {
			s.l.Error("unable to delete order",
				zap.String("order_id", cmd.Id),
				zap.Error(err),
			)
			return &Err{"unable to delete order"}
		}

		// increase the reserved money of the user
		previousTotal := order.Price * float64(order.Quantity)
		user.ReservedMoney = user.ReservedMoney - previousTotal
		s.activeUsers[userId] = user

		return nil
	} else {
		// make sure the user has sufficient fund and update the total reserved money
		previousTotal := order.Price * float64(order.Quantity)
		newTotalPrice := cmd.Price * float64(cmd.Quantity)
		if order.Type == store.OrderTypeBuy {
			if (user.Money - user.ReservedMoney + previousTotal) < newTotalPrice {
				s.l.Error("user has insufficient funds",
					zap.Float64("price", cmd.Price),
					zap.Int("quantity", cmd.Quantity),
					zap.Error(err),
				)
				return &Err{"user has insufficient funds"}
			}
		} else if order.Type == store.OrderTypeSell {
			// make sure the user has amount of the stock
			if _, ok := user.ReservedStonks[StonkName(order.Stonk)]; !ok {
				// init if not yet already intialized
				user.ReservedStonks[StonkName(order.Stonk)] = 0
			}

			if (user.Stonks[StonkName(order.Stonk)] - user.ReservedStonks[StonkName(order.Stonk)] + order.Quantity) < cmd.Quantity {
				s.l.Error("user has insufficient stocks",
					zap.String("stonk", string(StonkName(order.Stonk))),
					zap.Int("user_stonks", user.Stonks[StonkName(order.Stonk)]),
					zap.Int("user_reserved_stonks", user.ReservedStonks[StonkName(order.Stonk)]),
					zap.Int("quantity", cmd.Quantity),
					zap.Error(err),
				)
				return &Err{"user has insufficient stocks"}
			}
		} else {
			s.l.Error("unknown OrderType", zap.String("store_order_type", string(order.Type)))
			return &Err{"unknown OrderType"}
		}

		// update the previous order object
		newOrder := store.Order{
			Id:       order.Id,
			Stonk:    order.Stonk,
			Quantity: cmd.Quantity,
			Price:    cmd.Price,
			Type:     order.Type,
			User:     order.User,
			Time:     time.Now(),
		}

		// insert the order
		err = s.orderP.UpdateOrder(r.Context(), newOrder)
		if err != nil {
			s.l.Error("unable to update order",
				zap.String("order_id", cmd.Id),
				zap.Error(err),
			)
			return &Err{"unable to update order"}
		}

		if order.Type == store.OrderTypeBuy {
			// increase the reserved money of the user
			user.ReservedMoney = user.ReservedMoney - previousTotal + newTotalPrice
			s.activeUsers[userId] = user
		} else if order.Type == store.OrderTypeSell {
			// increase the reserved stocks of the user
			user.ReservedStonks[StonkName(order.Stonk)] += cmd.Quantity - order.Quantity
			s.activeUsers[userId] = user

		} else {
			s.l.Error("unknown OrderType", zap.String("store_order_type", string(order.Type)))
			return &Err{"unknown OrderType"}
		}

		return nil
	}
}
