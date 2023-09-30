package models

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FollowspotCue struct {
	Id uuid.UUID `json:"id" gorm:"primaryKey;unique;not null"`
	Number int `json:"number"`
	Frame int `json:"frame"`
	Size string `json:"size"`
	Target string `json:"target"`
	Description string `json:"description"`
	Location string `json:"location"`
	Notes string `json:"notes"`
	FollowspotId uuid.UUID `json:"followspot_id"`
	Followspot Followspot `json:"followspot"`
}


func (fsq *FollowspotCue) BeforeCreate(tx *gorm.DB) error {
	if fsq.Id == uuid.Nil {
		fsq.Id = uuid.New()
	}
	return nil
}

func GetFollowspotCueById(db *gorm.DB, id uuid.UUID) (*FollowspotCue, error) {
	var cue FollowspotCue
	err := db.First(&cue, id).Error
	if err != nil {
		log.Println("Error getting user: ", err)
		return &FollowspotCue{}, err
	}
	return &cue, nil
}

func (fsq *FollowspotCue) Create(db *gorm.DB) (*FollowspotCue, error) {
	err := db.Create(&fsq).Error
	if err != nil {
		log.Println("Error creating user: ", err)
		return &FollowspotCue{}, err
	}
	return fsq, nil	
}

func (fsq *FollowspotCue) Update(db *gorm.DB) (*FollowspotCue, error) {
	err := db.Save(&fsq).Error
	if err != nil {
		log.Println("Error updating user: ", err)
		return &FollowspotCue{}, err
	}
	return fsq, nil	
}

func (fsq *FollowspotCue) Delete(db *gorm.DB) (*FollowspotCue, error) {
	err := db.Delete(&fsq).Error
	if err != nil {
		log.Println("Error deleting user: ", err)
		return &FollowspotCue{}, err
	}
	return fsq, nil	
}

func DeleteFollowspotCueById(db *gorm.DB, id uuid.UUID) (error) {
	err := db.Delete(&User{}, id).Error
	if err != nil {
		log.Println("Error deleting user: ", err)
		return err
	}
	return nil
}

func GetAllFollowspotCues(db *gorm.DB) ([]FollowspotCue, error){
	var cues []FollowspotCue
	err := db.Find(&cues).Error
	if err != nil {
		log.Println("Error getting all users: ", err)
		return []FollowspotCue{}, err
	}
	return cues, nil
}
