package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Model for Gorm to use in Database
type Fixture struct {
	Id            uuid.UUID   `json:"id" gorm:"primaryKey;unique;not null"`
	IsVWX         bool        `json:"is_vwx,omitempty"`
	VWX_UID       string      `json:"uid,omitempty"`
	IsGDTF        bool        `json:"is_gdtf,omitempty"`
	GDTF_UID      string      `json:"gdtf_uid,omitempty"`
	InstType      string      `json:"instrument_type,omitempty"`
	Wattage       string      `json:"wattage,omitempty"`
	Purpose       string      `json:"purpose,omitempty"`
	Position      string      `json:"position,omitempty"`
	UnitNumber    string      `json:"unit_number,omitempty"`
	Color         string      `json:"color,omitempty"`
	Dimmer        string      `json:"dimmer,omitempty"`
	Channel       string      `json:"channel,omitempty"`
	Universe      string      `json:"universe,omitempty"`
	UDimmer       string      `json:"u_dimmer,omitempty"`
	CircuitNumber string      `json:"circuit_number,omitempty"`
	CircuitName   string      `json:"circuit_name,omitempty" `
	System        string      `json:"system,omitempty" `
	NumChannels   string      `json:"num_channels,omitempty"`
	Weight        string      `json:"weight,omitempty"`
	Focus         string      `json:"focus,omitempty"`
	FixtureID     string      `json:"fixture_id,omitempty"`
	Address       string      `json:"address,omitempty"`
	Voltage       string      `json:"voltage,omitempty"`
	FixtureMode   string      `json:"fixture_mode,omitempty" `
	Accessories   []Accessory `json:"accessories" gorm:"foreignKey:FixtureID"`
	XLocmm        float32     `json:"x_loc_mm,omitempty" `
	YLocmm        float32     `json:"y_loc_mm,omitempty" `
	ZLocmm        float32     `json:"z_loc_mm,omitempty" `
	EventID       uuid.UUID   `json:"event_id" gorm:"type:uuid;not null"`
	Event         Event       `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (f *Fixture) BeforeCreate() error {
	if f.Id == uuid.Nil {
		f.Id = uuid.New()
	}
	return nil
}

func (f Fixture) Create(db *gorm.DB) (Fixture, error) {
	err := db.Create(&f).Error
	if err != nil {
		return Fixture{}, err
	}
	return f, nil
}

func (f Fixture) Update(db *gorm.DB) (Fixture, error) {
	err := db.Save(&f).Error
	if err != nil {
		return Fixture{}, err
	}
	return f, nil
}

func (f Fixture) Delete(db *gorm.DB) (Fixture, error) {
	err := db.Delete(&f).Error
	if err != nil {
		return Fixture{}, err
	}
	return f, nil
}

func GetFixtureById(db *gorm.DB, id uuid.UUID) (Fixture, error) {
	var fixture Fixture
	err := db.First(&fixture, id).Error
	if err != nil {
		return Fixture{}, err
	}
	return fixture, nil
}
