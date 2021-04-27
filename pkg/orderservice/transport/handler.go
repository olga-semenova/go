package transport

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
)

type kitty struct {
	Name string `json:"n"`
}
type Order struct {
	Name     string `json:"name"`
	Id       string `json:"id"`
	Quantity string `json:"quantity"`
}

func Router() http.Handler {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/hello-world", helloWorld).Methods(http.MethodGet)
	s.HandleFunc("/get-kitty", getKitty).Methods(http.MethodGet)
	s.HandleFunc("/order", createOrder).Methods(http.MethodPost)
	return r
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var msg Order
	err = json.Unmarshal(b, &msg)
	if err != nil {
	}

	log.WithFields(log.Fields{
		"method": r.Method,
	}).Info("create order")
}

func helloWorld(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func logMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":        r.Method,
			"url":           r.URL,
			"remoteAddress": r.RemoteAddr,
			"userAgent":     r.UserAgent(),
		}).Info("got a new request")
		h.ServeHTTP(w, r)
	})
}

func getKitty(w http.ResponseWriter, _ *http.Request) {
	cat := kitty{"кот"}
	b, err := json.Marshal(cat)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err = io.WriteString(w, string(b)); err != nil {
		log.WithField("err", err).Error("write response error")
	}
}
