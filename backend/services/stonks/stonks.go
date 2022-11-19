package stonks

import (
	"errors"
	http "net/http"
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

type ScalarError string

func (se *ScalarError) Error() string {
	return string(*se)
}

type ScalarInPlace string

type StonksService struct {
	l *zap.Logger

	orderP store.OrderPersistor
	matchP store.MatchPersistor

	// users are only held ephemeraly
	waitingUsers []User
	activeUsers  []User
}

func NewStonksService(
	l *zap.Logger,
	orderP store.OrderPersistor,
	matchP store.MatchPersistor,
) *StonksService {
	return &StonksService{
		l:            l,
		orderP:       orderP,
		matchP:       matchP,
		waitingUsers: make([]User, 0, 5),
		activeUsers:  make([]User, 0, 5),
	}
}

type User struct {
	id string // NOTE: private on purpose
	// TODO: Probably need to add the ID without leaking it to other users (impersenation!)
	Name string

	// TODO: Deduct the money once an order is placed not when it is executed!
	money float64
}

type StonkInfo struct {
	ID string

	// TODO: Add the graph data
	// History map[]
	// TODO: Sort by timestamps!
	MatchHistory []Match
	Orders       []Order
}

type Match struct {
	UserSell  string
	UserBuy   string
	Quantity  int
	TimeStamp int64
}

type Order struct {
	User      string
	OrderType OrderType
	Quantity  int
	TimeStamp int64
}

type OrderType string

const (
	OrderTypeSell = "sell"
	OrderTypeBuy  = "buy"
)

func (s *StonksService) NewUser(w http.ResponseWriter, r *http.Request, name string) *Err {
	if r.Method != http.MethodPost {
		return &Err{"you have to post"}
	}

	cookie, err := r.Cookie("user")
	if errors.Is(err, http.ErrNoCookie) {
		// nothing to do
	} else {
		// try to find the user by the id
		for _, u := range s.waitingUsers {
			if u.id == cookie.Value {
				return &Err{"user already registered"}
			}
		}
	}

	u := User{
		id:   uuid.New().String(),
		Name: name,
	}
	s.waitingUsers = append(s.waitingUsers, u)

	// Set a cookie
	cookie = &http.Cookie{Name: "user", Value: u.id, Expires: time.Now().Add(time.Hour * 24 * 7)}
	http.SetCookie(w, cookie)
	return nil
}

// TODO: Actually this should be an SSE
func (s *StonksService) StartSession(w http.ResponseWriter, r *http.Request, id string) ([]User, *Err) {
	// TODO: Need to clear the users after one round
	if len(s.activeUsers) != 0 {
		return nil, &Err{"other session still active"}
	}

	// make the waitingUsers the active ones
	s.activeUsers = s.waitingUsers

	return s.activeUsers, nil
}

func (s *StonksService) GetStonkInfo(w http.ResponseWriter, r *http.Request, stonk string) (StonkInfo, *Err) {
	if r.Method != http.MethodGet {
		return StonkInfo{}, &Err{"you have to get"}
	}

	// FIXME: Somehow verify the user

	// TODO: Get the data from the collections
	storeOrders, err := s.orderP.GetOrders(r.Context(), store.Stonk(stonk), nil)
	if err != nil {
		return StonkInfo{}, &Err{"unable to retrieve orders"}
	}

	// transform the orders
	orders := ordersToStonksVo(storeOrders)

	storeMatches, err := s.matchP.GetMatches(r.Context(), store.Stonk(stonk), nil)
	if err != nil {
		return StonkInfo{}, &Err{"unable to retrieve orders"}
	}

	// transform the orders
	matches := matchsToStonksVo(storeMatches)

	// TODO: Transform the data
	// TODO: return the shit

	return StonkInfo{
		Orders:       orders,
		MatchHistory: matches,
	}, nil
}
