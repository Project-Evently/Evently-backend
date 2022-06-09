package repository

import (
	"context"
	"errors"
	"github.com/Project-Evently/Evently-backend/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

type EventDbSql struct {
	pool *pgxpool.Pool
}

func NewEventDbSql(pool *pgxpool.Pool) *EventDbSql {
	return &EventDbSql{
		pool: pool,
	}
}

func (r *EventDbSql) ReadEvent(EventId int) (*entity.Event, error) {
	var event *entity.Event

	var eventId int
	var clubId int
	var instituteId int
	var description string
	var eventDateIST string
	var eventTimeIST string
	var eventLocation string
	var eventOrganizer string
	var eventContact string
	var eventLink string
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	row := tx.QueryRow(context.Background(), "select event_id, club_id, institute_id, description, event_date, event_time, event_location, event_organizer, event_contact, event_link from events where event_id = $1", EventId)

	err = row.Scan(&eventId, &clubId, &instituteId, &description, &eventDateIST, &eventTimeIST, &eventLocation, &eventOrganizer, &eventContact, &eventLink)
	event = &entity.Event{
		EventId:        eventId,
		ClubId:         clubId,
		InstituteId:    instituteId,
		Description:    description,
		EventDateIST:   eventDateIST,
		EventTimeIST:   eventTimeIST,
		EventLocation:  eventLocation,
		EventOrganizer: eventOrganizer,
		EventContact:   eventContact,
		EventLink:      eventLink,
	}

	if event != nil {
		return nil, errors.New("error while fetching data from events table")
	}
	tx.Commit(context.Background())
	return event, nil
}

func (r *EventDbSql) ReadListInstitute(InstituteId int) ([]*entity.Event, error) {
	EventList := make([]*entity.Event, 1)

	var event *entity.Event

	var eventId int
	var clubId int
	var instituteId int
	var description string
	var eventDateIST string
	var eventTimeIST string
	var eventLocation string
	var eventOrganizer string
	var eventContact string
	var eventLink string

	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(context.Background(), "select event_id, club_id, institute_id, description, event_date, event_time, event_location, event_organizer, event_contact, event_link from events where institute_id = $1", InstituteId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&eventId, &clubId, &instituteId, &description, &eventDateIST, &eventTimeIST, &eventLocation, &eventOrganizer, &eventContact, &eventLink)
		if err != nil {
			return nil, err
		}
		event = &entity.Event{
			EventId:        eventId,
			ClubId:         clubId,
			InstituteId:    instituteId,
			Description:    description,
			EventDateIST:   eventDateIST,
			EventTimeIST:   eventTimeIST,
			EventLocation:  eventLocation,
			EventOrganizer: eventOrganizer,
			EventContact:   eventContact,
			EventLink:      eventLink,
		}

		EventList = append(EventList, event)
	}

	return EventList, nil
}

func (r *EventDbSql) ReadListClub(ClubId int) ([]*entity.Event, error) {
	EventList := make([]*entity.Event, 1)

	var event *entity.Event

	var eventId int
	var clubId int
	var instituteId int
	var description string
	var eventDateIST string
	var eventTimeIST string
	var eventLocation string
	var eventOrganizer string
	var eventContact string
	var eventLink string

	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(context.Background(), "select event_id, club_id, institute_id, description, event_date, event_time, event_location, event_organizer, event_contact, event_link from events where club_id = $1", ClubId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&eventId, &clubId, &instituteId, &description, &eventDateIST, &eventTimeIST, &eventLocation, &eventOrganizer, &eventContact, &eventLink)
		if err != nil {
			return nil, err
		}
		event = &entity.Event{
			EventId:        eventId,
			ClubId:         clubId,
			InstituteId:    instituteId,
			Description:    description,
			EventDateIST:   eventDateIST,
			EventTimeIST:   eventTimeIST,
			EventLocation:  eventLocation,
			EventOrganizer: eventOrganizer,
			EventContact:   eventContact,
			EventLink:      eventLink,
		}

		EventList = append(EventList, event)
	}
	tx.Commit(context.Background())
	return EventList, nil
}

func (r *EventDbSql) WriteEvent(event *entity.Event) error {
	tx, err := r.pool.Begin(context.Background())
	if err != nil {
		return err
	}

	ct, err := tx.Exec(context.Background(), "INSERT INTO events (club_id, institute_id, description, event_date, event_time, event_location, event_organizer, event_contact, event_link) values ($1,$2,$3,$4,$5,$6,$7,$8,$9)", event.ClubId, event.InstituteId, event.Description, event.EventDateIST, event.EventTimeIST, event.EventLocation, event.EventOrganizer, event.EventContact, event.EventLink)
	if err != nil {
		return err
	}

	if ct.RowsAffected() < 0 {
		return errors.New("error while adding new event in event table")
	}

	tx.Commit(context.Background())
	return nil
}
