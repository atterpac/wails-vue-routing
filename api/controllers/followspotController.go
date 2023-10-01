package controllers

import (
	"log"
	"yagami/api/models"

	"github.com/google/uuid"
)

func (a *App) GetFollowspots(eventId uuid.UUID) []models.FollowspotCue {
	f, err := models.GetFollowspotCuesByEvent(a.Db, eventId)
	if err != nil {
		log.Println("Error Getting FollowspotCues: ", err)
		return nil
	}
	return f
}

func (a *App) UpdateFollowspot(followspot models.FollowspotCue) (models.FollowspotCue, error) {
	updateFollowspot, err := followspot.Update(a.Db)
	if err != nil {
		log.Println("Error Updating FollowspotCue: ", err)
		return followspot, err
	}
	return *updateFollowspot, nil
}

func (a *App) CreateFollowspot(followspot models.FollowspotCue) (models.FollowspotCue, error) {
	createdFollowspot, err := followspot.Create(a.Db)
	if err != nil {
		log.Println("Error Creating FollowspotCue: ", err)
		return followspot, err
	}
	return *createdFollowspot, nil
}
	
func (a *App) DeleteFollowspot(followspot models.FollowspotCue) (models.FollowspotCue, error) {
	deletedFollowspot, err := followspot.Delete(a.Db)
	if err != nil {
		log.Println("Error Deleting FollowspotCue: ", err)
		return followspot, err
	}
	return *deletedFollowspot, nil
}

func (a *App) GetFollowspotById(id uuid.UUID) (models.FollowspotCue, error) {
	followspot, err := models.GetFollowspotCueById(a.Db, id)
	if err != nil {
		log.Println("Error Getting FollowspotCue By Id: ", err)
		return models.FollowspotCue{}, err
	}
	return *followspot, nil
}
