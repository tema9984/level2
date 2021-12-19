package calendar

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"gitub.com/tema9984/dev11/config"
	"gitub.com/tema9984/dev11/internal/model"
	"gitub.com/tema9984/dev11/internal/store"
)

type Service interface {
	GetEventsForDay(w http.ResponseWriter, r *http.Request)
	GetEventsForWeek(w http.ResponseWriter, r *http.Request)
	GetEventsForMonth(w http.ResponseWriter, r *http.Request)
	CreateEvent(w http.ResponseWriter, r *http.Request)
	UpdateEvent(w http.ResponseWriter, r *http.Request)
	DeleteEvent(w http.ResponseWriter, r *http.Request)
	Logging(next http.Handler) http.Handler
	Close()
}

type service struct {
	conf   *config.Config
	store  *store.Store
	logger *logrus.Logger
}

// NewService ...
func NewService(cfg *config.Config) Service {
	svc := &service{conf: cfg}
	svc.configureService()
	return svc
}
func (svc *service) configureService() {
	err := svc.configureStore()
	svc.logger = logrus.New()
	if err != nil {
		logrus.Error(err)
	}
}

func (svc *service) configureStore() error {
	st := store.New(svc.conf)
	if err := st.Open(); err != nil {
		logrus.Error("ERROR IN OPEN DB")
		return err
	}

	svc.store = st
	return nil
}

////////////////////////////////
func (svc *service) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	var err error
	var jn []byte
	defer func() {
		if err != nil {
			w.Write(ErrToJson(model.ErrRespons{Err: err}))
		} else {
			w.Write(jn)
		}
	}()
	date, ok := r.URL.Query()["date"]
	if !ok || len(date[0]) != 10 {
		w.WriteHeader(http.StatusBadRequest)
		err = errors.New("BadRequest")
		return
	}
	ans, err := svc.store.Ev().EventsForDay(date[0])
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"desc": "store.Ev().EventsForDay(date[0]) err",
			"func": "GetEventsForDay",
		}).Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	jn = EventsToJson(ans)

}
func (svc *service) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.Write(ErrToJson(model.ErrRespons{Err: err}))
		}
	}()
	date, ok := r.URL.Query()["date"]
	if !ok || len(date[0]) != 10 {
		w.WriteHeader(http.StatusBadRequest)
		err = errors.New("BadRequest")
		return
	}
	ans, err := svc.store.Ev().EventsForWeek(date[0])
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"desc": "store.Ev().EventsForWeek(date[0]) err",
			"func": "GetEventsForWeek",
		}).Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	jn := EventsToJson(ans)
	w.Write(jn)

}
func (svc *service) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.Write(ErrToJson(model.ErrRespons{Err: err}))
		}
	}()
	date, ok := r.URL.Query()["date"]
	if !ok || len(date[0]) != 10 {
		w.WriteHeader(http.StatusBadRequest)
		err = errors.New("BadRequest")
		return
	}
	ans, err := svc.store.Ev().EventsForWeek(date[0])
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"desc": "store.Ev().EventsForMonth(date[0]) err",
			"func": "GetEventsForMonth",
		}).Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	jn := EventsToJson(ans)
	w.Write(jn)

}
func (svc *service) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	var err error
	var modEv model.Event
	defer func() {
		if err != nil {
			w.Write(ErrToJson(model.ErrRespons{Err: err}))
		} else {
			w.Write(EventsToJson([]model.Event{modEv}))
		}
	}()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"desc": "ioutil.ReadAll(r.Body) err",
			"func": "CreateEvent",
		}).Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	modEv = JsonToEvent(b)
	err = svc.store.Ev().DeleteEvent(modEv.Id)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"desc": "store.Ev().DeleteEvent(modEv.Id) err",
			"func": "DeleteEvent",
		}).Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
}
func (svc *service) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	var err error
	var modEv model.Event
	defer func() {
		if err != nil {
			w.Write(ErrToJson(model.ErrRespons{Err: err}))
		} else {
			w.Write(EventsToJson([]model.Event{modEv}))
		}
	}()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"desc": "ioutil.ReadAll(r.Body) err",
			"func": "UpdateEvent",
		}).Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	modEv = JsonToEvent(b)
	err = svc.store.Ev().UpdateEvent(modEv)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"desc": "store.Ev().UpdateEvent(modEv) err",
			"func": "UpdateEvent",
		}).Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
}
func (svc *service) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var err error
	var modEv model.Event
	defer func() {
		if err != nil {
			w.Write(ErrToJson(model.ErrRespons{Err: err}))
		} else {
			w.Write(EventsToJson([]model.Event{modEv}))
		}
	}()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"desc": "ioutil.ReadAll(r.Body) err",
			"func": "CreateEvent",
		}).Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	modEv = JsonToEvent(b)
	_, err = svc.store.Ev().AddEvent(modEv)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"desc": "store.Ev().DeleteEvent(id[0]) err",
			"func": "DeleteEvent",
		}).Error(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
}
func EventsToJson(ev []model.Event) []byte {
	r, err := json.Marshal(model.Respons{Result: ev})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"desc": "json.Marshal(ev) err",
			"func": "EventsToJson",
		}).Error(err)
		return []byte{}
	}
	return r
} ////////////////////////model.ErrRespons сделать везде респонсы для ошибок
func ErrToJson(er model.ErrRespons) []byte {
	r, err := json.Marshal(er)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"desc": "json.Marshal(ev) err",
			"func": "EventsToJson",
		}).Error(err)
		return []byte{}
	}
	return r
}
func JsonToEvent(evJn []byte) model.Event {
	ev := model.Event{}
	err := json.Unmarshal(evJn, &ev)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"desc": "err := json.Unmarshal(evJn, &ev) err",
			"func": "JsonToEvent",
		}).Error(err)
		return model.Event{}
	}
	return ev
}
func (svc *service) Close() {
	svc.store.Close()
}
func (svc *service) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		svc.logger.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}
