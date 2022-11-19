package stonks

import (
	http "net/http"
	time "time"

	"github.com/google/uuid"
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
	waitingUsers []User
	activeUsers  []User
}

type User struct {
	id string // NOTE: private on purpose
	// TODO: Probably need to add the ID without leaking it to other users (impersenation!)
	Name string
}

type StonkInfo struct {
	ID string
	// TODO: Sort by timestamps!
	History []Match
	Orders  []Order
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

	u := User{
		id:   uuid.New().String(),
		Name: name,
	}
	s.waitingUsers = append(s.waitingUsers, u)

	// Set a cookie
	cookie := http.Cookie{Name: "user", Value: u.id, Expires: time.Now().Add(time.Hour * 24 * 7)}
	http.SetCookie(w, &cookie)
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

func (s *StonksService) GetStonkInfo(w http.ResponseWriter, r *http.Request, id string) (StonkInfo, *Err) {
	// FIXME: Somehow verify the user

	// TODO: Get the data from the collections

	// TODO: Transform the data
	// TODO: return the shit

	return StonkInfo{}, nil
}

func (s *StonksService) Hello(name string) (string, *Err) {
	if name == "Peter" {
		return "", &Err{"fuck you Peter I do not like you"}
	}
	return "Hello from the server: " + name, nil
}

func (s *StonksService) HelloInterface(anything interface{}, anythingMap map[string]interface{}, anythingSlice []interface{}) {

}

func (s *StonksService) HelloNumberMaps(intMap map[int]string) (floatMap map[float64]string) {
	floatMap = map[float64]string{}
	for i, str := range intMap {
		floatMap[float64(i)] = str
	}
	return
}

func (s *StonksService) HelloScalarError() (err *ScalarError) {
	return
}

func (s *StonksService) nothingInNothinOut() {

}
