package controllers

import (
	"log"
	"yagami/api/models"

	"github.com/google/uuid"
)

func (a *App) GetEvents() []models.Event {
	e, err := models.GetAllEvents(a.Db)
	if err != nil {
		log.Println("Error Getting Events: ", err)
		return nil
	}
	return e
}

func (a *App) UpdateEvent(event models.Event) (models.Event, error) {
	e, err := event.Update(a.Db)
	if err != nil {
		log.Println("Error Updating Event: ", err)
		return event, err
	}
	return e, nil
}

func (a *App) CreateEvent(event models.Event) (models.Event, error) {
	e, err := event.Create(a.Db)
	if err != nil {
		log.Println("Error Creating Event: ", err)
		return event, err
	}
	return e, nil
}

func (a *App) DeleteEvent(event models.Event) (models.Event, error) {
	e, err := event.Delete(a.Db)
	if err != nil {
		log.Println("Error Deleting Event: ", err)
		return event, err
	}
	return e, nil
}

func (a *App) GetEventById(id uuid.UUID) (models.Event, error) {
	e, err := models.GetEventById(a.Db, id)
	if err != nil {
		log.Println("Error Getting Event By Id: ", err)
		return models.Event{}, err
	}
	return e, nil
}

func (a *App) SelectEvent(eventId uuid.UUID) {
	_, err := models.GetEventById(a.Db, eventId)
	if err != nil {
		log.Println("Error Getting Event By Id: ", err)
		return
	}
	a.CurrentEvent = eventId
}


