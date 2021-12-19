package store

import (
	"github.com/sirupsen/logrus"
	"gitub.com/tema9984/dev11/internal/model"
)

type EventRepository struct {
	store *Store
}

func (rep *EventRepository) AddEvent(e model.Event) (model.Event, error) {
	var err error
	var id int
	str := "insert into calendar(event_name, event_date, descriptions) values ($1, $2, $3) RETURNING id"
	err = rep.store.db.QueryRow(str, e.EventName, e.EventDate, e.Descripton).Scan(&id)
	if err != nil {
		logrus.Error(err)
	}
	return e, err
}
func (rep *EventRepository) UpdateEvent(e model.Event) error {
	var err error
	str := "update calendar set event_name=$1, event_date=$2, descriptions=$3  where id = $4"
	_, err = rep.store.db.Exec(str, e.EventName, e.EventDate, e.Descripton, e.Id)
	if err != nil {
		logrus.Error(err)
	}
	return err
}
func (rep *EventRepository) DeleteEvent(id int) error {
	var err error
	str := "delete from calendar where id = $1"
	_, err = rep.store.db.Exec(str, id)
	if err != nil {
		logrus.Error(err)
	}
	return err
}
func (rep *EventRepository) getEvents(date, query string) ([]model.Event, error) {
	rows, err := rep.store.db.Query(query, date)
	if err != nil {
		return []model.Event{}, err
	}
	defer rows.Close()
	events := []model.Event{}
	for rows.Next() {
		e := model.Event{}
		err = rows.Scan(&e.Id, &e.EventName, &e.EventDate, &e.Descripton)
		if err != nil {
			logrus.Warning(err)
			continue
		}
		events = append(events, e)
	}

	return events, nil
}
func (rep *EventRepository) EventsForDay(date string) ([]model.Event, error) {
	return rep.getEvents(date, "select * from calendar where event_date=$1")
}
func (rep *EventRepository) EventsForWeek(date string) ([]model.Event, error) {
	return rep.getEvents(date, `select * from calendar where EXTRACT(WEEk FROM calendar.event_date) = EXTRACT(WEEk FROM $1::TIMESTAMP)`)
}
func (rep *EventRepository) EventsForMonth(date string) ([]model.Event, error) {

	return rep.getEvents(date, `select * from calendar where EXTRACT(Month FROM calendar.event_date) = EXTRACT(Month FROM $1::TIMESTAMP)`)
}
