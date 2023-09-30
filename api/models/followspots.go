package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Followspot struct {
	Id uuid.UUID `json:"id" gorm:"primaryKey;unique;not null"`
	Number int `json:"number"`
	Fixture string `json:"fixture"`
	Location string `json:"location"`
	Notes string `json:"notes"`
	Frames []Frame `json:"frames"`
	EventId uuid.UUID `json:"event_id"`
	FollowspotCues []FollowspotCue `json:"followspot_cues"`

	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Frame struct {
	Id uuid.UUID `json:"id" gorm:"primaryKey;unique;not null"`
	Index int `json:"index"`
	Label string `json:"label"`
	Number int `json:"number"`
	FollowspotId uuid.UUID `json:"followspot_id"`
}


func (f *Followspot) BeforeCreate(tx *gorm.DB) error {
	if f.Id == uuid.Nil {
		f.Id = uuid.New()
	}
	return nil
}

func GetFollowspotById(db *gorm.DB, id uuid.UUID) (*Followspot, error) {
	var followspot Followspot
	err := db.First(&followspot, id).Error
	if err != nil {
		return &Followspot{}, err
	}
	return &followspot, nil
}

func (f *Followspot) Create(db *gorm.DB) (*Followspot, error) {
	err := db.Create(&f).Error
	if err != nil {
		return &Followspot{}, err
	}
	return f, nil
}

func (f *Followspot) Update(db *gorm.DB) (*Followspot, error) {
	err := db.Save(&f).Error
	if err != nil {
		return &Followspot{}, err
	}
	return f, nil
}

func (f *Followspot) Delete(db *gorm.DB) (*Followspot, error) {
	err := db.Delete(&f).Error
	if err != nil {
		return &Followspot{}, err
	}
	return f, nil
}

func GetAllFollowspots(db *gorm.DB) ([]Followspot, error) {
	var followspots []Followspot
	err := db.Find(&followspots).Error
	if err != nil {
		return []Followspot{}, err
	}
	return followspots, nil
}
