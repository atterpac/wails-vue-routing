package controllers

import (
	"log"
	"yagami/api/models"

	"github.com/google/uuid"
)


func (a *App) GetFollowspotCues(eventId uuid.UUID) ([]models.FollowspotCue, error) {
	q, err := models.GetFollowspotCuesByEvent(a.Db, eventId)
	if err != nil {
		log.Println("Error Getting FollowspotCues: ", err)
		return nil, err
	}
	return q, nil
}

func (a *App) UpdateFollowspotCue(followspot models.FollowspotCue) (models.FollowspotCue, error) {
	q, err := followspot.Update(a.Db)
	if err != nil {
		log.Println("Error Updating FollowspotCue: ", err)
		return followspot, err
	}
	return *q, nil
}

func (a *App) CreateFollowspotCue(followspot models.FollowspotCue) (models.FollowspotCue, error) {
	q, err := followspot.Create(a.Db)
	if err != nil {
		log.Println("Error Creating FollowspotCue: ", err)
		return followspot, err
	}
	return *q, nil
}

func (a *App) DeleteFollowspotCue(followspot models.FollowspotCue) (models.FollowspotCue, error) {
	q, err := followspot.Delete(a.Db)
	if err != nil {
		log.Println("Error Deleting FollowspotCue: ", err)
		return followspot, err
	}
	return *q, nil
}

func (a *App) GetFollowspotCueById(id uuid.UUID) (models.FollowspotCue, error) {
	q, err := models.GetFollowspotCueById(a.Db, id)
	if err != nil {
		log.Println("Error Getting FollowspotCue By Id: ", err)
		return models.FollowspotCue{}, err
	}
	return *q, nil
}


