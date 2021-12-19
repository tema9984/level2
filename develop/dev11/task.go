package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"gitub.com/tema9984/dev11/config"
	"gitub.com/tema9984/dev11/internal/calendar"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API:
POST /create_event
POST /update_event
POST /delete_event
GET /events_for_day
GET /events_for_week
GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/
const configPath = "config/config.json"

func main() {
	cfg, err := config.ParseConfig(configPath)
	if err != nil {
		logrus.Fatal(err)
	}
	mux := http.NewServeMux()
	svc := calendar.NewService(cfg)

	mux.HandleFunc("/create_event", svc.CreateEvent)
	mux.HandleFunc("/update_event", svc.UpdateEvent)
	mux.HandleFunc("/delete_event", svc.DeleteEvent)
	mux.HandleFunc("/events_for_day", svc.GetEventsForDay)
	mux.HandleFunc("/events_for_week", svc.GetEventsForWeek)
	mux.HandleFunc("/events_for_month", svc.GetEventsForMonth)
	hand := svc.Logging(mux)
	cors := cors(hand)
	s := &http.Server{
		Addr:         ":" + cfg.WebPort,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      cors,
	}
	log.Fatal(s.ListenAndServe())
}
func handlerDay(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "page")
}
func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Accept")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,POST,GET")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, req)
		w.WriteHeader(http.StatusOK)
	})
}
