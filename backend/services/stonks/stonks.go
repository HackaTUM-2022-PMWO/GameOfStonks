package stonks

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

type Stonks struct {
	Bla bool
}

func (d *Stonks) Hello(name string) (string, *Err) {
	if name == "Peter" {
		return "", &Err{"fuck you Peter I do not like you"}
	}
	return "Hello from the server: " + name, nil
}

func (d *Stonks) HelloInterface(anything interface{}, anythingMap map[string]interface{}, anythingSlice []interface{}) {

}

func (d *Stonks) HelloNumberMaps(intMap map[int]string) (floatMap map[float64]string) {
	floatMap = map[float64]string{}
	for i, str := range intMap {
		floatMap[float64(i)] = str
	}
	return
}

func (d *Stonks) HelloScalarError() (err *ScalarError) {
	return
}

func (d *Stonks) nothingInNothinOut() {

}
