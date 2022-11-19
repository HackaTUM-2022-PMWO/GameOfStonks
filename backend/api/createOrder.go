package handler

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"

	"github.com/hackaTUM/GameOfStonks/store"
	"go.uber.org/zap"
)

type CreateOrderCmd struct {
	// TODO: Actually the ID should be set by the server
	Order store.Order `yaml:"order"`
}

func CreateOrderHandler(
	l *zap.Logger,
	store store.OrderPersistor,
) http.HandlerFunc {
	l = l.With(zap.String("endpoint", "create_oder"))

	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Validate the request

		// read the command
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			l.Error("unable to read body", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
		}

		// unmarshal
		var cmd CreateOrderCmd
		err = json.Unmarshal(data, &cmd)
		if err != nil {
			l.Error("unable to unmarshal command")
			w.WriteHeader(http.StatusBadRequest)
		}

		// TODO: validate the order

		// TODO: If the order does contain an ID

		// store the order
		err = store.AddOrder(&cmd.Order)
		if err != nil {
			l.Error("unable to persist order")
			w.WriteHeader(http.StatusBadRequest)
		}

		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}
}
