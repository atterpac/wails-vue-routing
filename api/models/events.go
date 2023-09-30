package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	Location string `json:"location"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Users []User `json:"users" gorm:"many2many:user_events;"`
	Followspots []Followspot `json:"followspots"`
}

func (e *Event) BeforeCreate(tx *gorm.DB) error {
	if e.Id == uuid.Nil {
		e.Id = uuid.New()
	}
	return nil
}

func (e Event) Create(db *gorm.DB) (Event, error) {
	err := db.Create(&e).Error
	if err != nil {
		log.Println("Error creating event: ", err)
		return Event{}, err
	}
	return e, nil
}

func (e Event) Update(db *gorm.DB) (Event, error) {
	err := db.Save(&e).Error
	if err != nil {
		log.Println("Error updating event: ", err)
		return Event{}, err
	}
	return e, nil
}

func (e Event) Delete(db *gorm.DB) (Event, error) {
	err := db.Delete(&e).Error
	if err != nil {
		log.Println("Error deleting event: ", err)
		return Event{}, err
	}
	return e, nil
}

func GetEventById(db *gorm.DB, id uuid.UUID) (Event, error) {
	var event Event
	err := db.First(&event, id).Error
	if err != nil {
		log.Println("Error getting event: ", err)
		return Event{}, err
	}
	return event, nil
}

func GetAllEvents(db *gorm.DB) ([]Event, error) {
	var events []Event
	err := db.Find(&events).Error
	if err != nil {
		log.Println("Error getting events: ", err)
		return []Event{}, err
	}
	return events, nil
}

