package stonks

import (
	http "net/http"
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
}

type StonkInfo struct {
	ID string
	// TODO: Sort by timestamps!
	History []Match
	// TODO: Add history of all users for this stonk
	// TODO: Add this users current info
}

type Match struct {
	UserSell  string
	UserBuy   string
	Quantity  int
	TimeStamp int64
}

func (s *StonksService) GetStonkInfo(w http.ResponseWriter, r *http.Request, id string) (StonkInfo, *Err) {
	/*
		if strings.Contains(identifier, "@") {
			return s.validateAndExistsEmail(w, r, identifier)
		}
		return s.validateAndExistsLoyalty(w, r, identifier)
	*/
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
