package models

import (
	"encoding/xml"
	"fmt"

	"github.com/google/uuid"
	guuid "github.com/google/uuid"

	"gorm.io/gorm"
)

// Model for Gorm to use in Database
type VWXFixture struct {
	Id                   uuid.UUID   `json:"id" gorm:"primaryKey;unique;not null"`
	Action               string      `json:"action"`
	TimeStamp            int         `json:"time_stamp"`
	AppStamp             string      `json:"app_stamp"`
	UID                  string      `json:"uid,omitempty"`
	InstType             string      `json:"instrument_type,omitempty"`
	Wattage              string      `json:"wattage,omitempty"`
	Purpose              string      `json:"purpose,omitempty"`
	Position             string      `json:"position,omitempty"`
	UnitNumber           string      `json:"unit_number,omitempty"`
	Color                string      `json:"color,omitempty"`
	Dimmer               string      `json:"dimmer,omitempty"`
	Channel              string      `json:"channel,omitempty"`
	Universe             string      `json:"universe,omitempty"`
	UDimmer              string      `json:"u_dimmer,omitempty"`
	CircuitNumber        string      `json:"circuit_number,omitempty"`
	CircuitName          string      `json:"circuit_name,omitempty" `
	System               string      `json:"system,omitempty" `
	UserField1           string      `json:"user_field_1,omitempty" `
	UserField2           string      `json:"user_field_2,omitempty" `
	UserField3           string      `json:"user_field_3,omitempty" `
	UserField4           string      `json:"user_field_4,omitempty" `
	UserField5           string      `json:"user_field_5,omitempty" `
	UserField6           string      `json:"user_field_6,omitempty"`
	NumChannels          string      `json:"num_channels,omitempty"`
	FrameSize            string      `json:"frame_size,omitempty"`
	FieldAngle           string      `json:"field_angle,omitempty"`
	FieldAngle2          string      `json:"field_angle_2,omitempty"`
	BeamAngle            string      `json:"beam_angle,omitempty"`
	BeamAngle2           string      `json:"beam_angle_2,omitempty" `
	Weight               string      `json:"weight,omitempty"`
	Gobo1                string      `json:"gobo_1,omitempty"`
	Gobo1Rotation        string      `json:"gobo_1_rot,omitempty"`
	Gobo2                string      `json:"gobo_2,omitempty" `
	Gobo2Rotation        string      `json:"gobo_2_rot,omitempty" `
	Mark                 string      `json:"mark,omitempty"`
	LampRotationAngle    string      `json:"lamp_rot_angle,omitempty" `
	Focus                string      `json:"focus,omitempty"`
	XRotation            string      `json:"x_rot,omitempty" `
	YRotation            string      `json:"y_rot,omitempty" `
	FixtureID            string      `json:"fixture_id,omitempty"`
	UseVerticalBeam      string      `json:"use_vertical_beam,omitempty" `
	Address              string      `json:"address,omitempty"`
	Voltage              string      `json:"voltage,omitempty"`
	BreakerID            string      `json:"breaker_id,omitempty" `
	Time                 string      `json:"time,omitempty"`
	Cost                 string      `json:"cost,omitempty" `
	VerticalFocusAngle   string      `json:"vertical_focus_angle,omitempty" `
	HorizontalFocusAngle string      `json:"horizontal_focus_angle,omitempty" `
	AngleToFace          string      `json:"angle_to_face,omitempty"`
	OffAxisAngle         string      `json:"off_axis_angle,omitempty"`
	ThrowDistance        string      `json:"throw_distance,omitempty"`
	FixtureMode          string      `json:"fixture_mode,omitempty" `
	EnableZRotation      string      `json:"enable_z_rot,omitempty" `
	ZRotation            string      `json:"z_rot,omitempty" `
	Accessories          []Accessory `json:"accessories" gorm:"foreignKey:FixtureID"`
	XLocmm               float32     `json:"x_loc_mm,omitempty" `
	YLocmm               float32     `json:"y_loc_mm,omitempty" `
	ZLocmm               float32     `json:"z_loc_mm,omitempty" `
	EventID              guuid.UUID  `json:"event_id" gorm:"type:uuid;not null"`
	Event                Event       `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Accessory struct {
	gorm.Model
	UID        string `json:"id" xml:"UID"`
	DeviceType string `json:"device_type" xml:"Device_Type"`
	SymbolName string `json:"symbol_name" xml:"Symbol_Name"`
	InstType   string `json:"instrument_type" xml:"Inst_Type"`
	FixtureID  int    `json:"fixture_id" gorm:"foreignKey:FixtureID"`
}

// Structs needed for parsing the XML
type Fixtures struct {
	XMLName        xml.Name       `xml:"SLData"`
	InstrumentData InstrumentData `xml:"InstrumentData"`
}

type InstrumentData struct {
	Action      string        `xml:"Action"`
	AppStamp    string        `xml:"AppStamp"`
	VWVersion   string        `xml:"VWVersion"`
	VWBuild     string        `xml:"VWBuild"`
	AutoRot2D   string        `xml:"AutoRot2D"`
	Instruments []Instruments `xml:",any"`
}

type Instruments struct {
	Action               string      `xml:"Action"`
	TimeStamp            int         `xml:"TimeStamp"`
	AppStamp             string      `xml:"AppStamp"`
	UID                  string      `xml:"UID"`
	InstType             string      `xml:"Inst_Type"`
	Wattage              string      `xml:"Wattage"`
	Purpose              string      `xml:"Purpose"`
	Position             string      `xml:"Position"`
	UnitNumber           string      `xml:"Unit_Number"`
	Color                string      `xml:"Color"`
	Dimmer               string      `xml:"Dimmer"`
	Channel              string      `xml:"Channel"`
	Universe             string      `xml:"Universe"`
	UDimmer              string      `xml:"U_Dimmer"`
	CircuitNumber        string      `xml:"Circuit_Number"`
	CircuitName          string      `xml:"Circuit_Name"`
	System               string      `xml:"System"`
	UserField1           string      `xml:"User_Field_1"`
	UserField2           string      `xml:"User_Field_2"`
	UserField3           string      `xml:"User_Field_3"`
	UserField4           string      `xml:"User_Field_4"`
	UserField5           string      `xml:"User_Field_5"`
	UserField6           string      `xml:"User_Field_6"`
	NumChannels          string      `xml:"Num_Channels"`
	FrameSize            string      `xml:"Frame_Size"`
	FieldAngle           string      `xml:"Field_Angle"`
	FieldAngle2          string      `xml:"Field_Angle_2"`
	BeamAngle            string      `xml:"Beam_Angle"`
	BeamAngle2           string      `xml:"Beam_Angle_2"`
	Weight               string      `xml:"Weight"`
	Gobo1                string      `xml:"Gobo_1"`
	Gobo1Rotation        string      `xml:"Gobo_1_Rotation"`
	Gobo2                string      `xml:"Gobo_2"`
	Gobo2Rotation        string      `xml:"Gobo_2_Rotation"`
	Mark                 string      `xml:"Mark"`
	LampRotationAngle    string      `xml:"Lamp_Rotation_Angle"`
	Focus                string      `xml:"Focus"`
	XRotation            string      `xml:"X_Rotation"`
	YRotation            string      `xml:"Y_Rotation"`
	FixtureID            string      `xml:"FixtureID"`
	UseVerticalBeam      string      `xml:"Use_Vertical_Beam"`
	Address              string      `xml:"Address"`
	Voltage              string      `xml:"Voltage"`
	BreakerID            string      `xml:"BreakerID"`
	Time                 string      `xml:"Time"`
	Cost                 string      `xml:"Cost"`
	VerticalFocusAngle   string      `xml:"Vertical_Focus_Angle"`
	HorizontalFocusAngle string      `xml:"Horizontal_Focus_Angle"`
	AngleToFace          string      `xml:"Angle_To_Face"`
	OffAxisAngle         string      `xml:"Off_Axis_Angle"`
	ThrowDistance        string      `xml:"Throw_Distance"`
	FixtureMode          string      `xml:"Fixture_Mode" `
	EnableZRotation      string      `xml:"Enable_Z_Rotation"`
	ZRotation            string      `xml:"Z_Rotation"`
	Accessories          Accessories `xml:"Accessories"`
	XLocmm               float32     `xml:"X_Location_mm"`
	YLocmm               float32     `xml:"Y_Location_mm"`
	ZLocmm               float32     `xml:"Z_Location_mm"`
}

type Accessories struct {
	Items []Accessory `xml:",any"`
}

func (f *VWXFixture) Save(db *gorm.DB) (*VWXFixture, error) {
	err := db.Save(f).Error
	if err != nil {
		fmt.Println("Error Saving: ", err)
		return f, err
	}
	return f, nil
}

func (f *VWXFixture) Delete(db *gorm.DB) (*VWXFixture, error) {
	err := db.Delete(f).Error
	if err != nil {
		fmt.Println("Error Deleting")
		return f, err
	}
	return f, nil
}

func (f *VWXFixture) FindFixturesByEvent(db *gorm.DB, event guuid.UUID) ([]VWXFixture, error) {
	var fixtures []VWXFixture
	err := db.Where("event_id = ?", event).Find(&fixtures).Error
	if err != nil {
		fmt.Println("Error Deleting")
		return nil, err
	}
	return fixtures, nil
}

func PushXML(db *gorm.DB, fixtures []VWXFixture) {
	i := 0
	for i < len(fixtures) {
		fixture, err := fixtures[i].Save(db)
		if err != nil {
			fmt.Println("Error %w", err)
		}
		fmt.Println("Saved Fixture: ", fixture)
		i++
	}
}

func ConvertToFixtures(instruments []Instruments) []VWXFixture {
	var fixtures []VWXFixture

	for _, instrument := range instruments {
		fixture := VWXFixture{
			Action:               instrument.Action,
			TimeStamp:            instrument.TimeStamp,
			AppStamp:             instrument.AppStamp,
			UID:                  instrument.UID,
			InstType:             instrument.InstType,
			Wattage:              instrument.Wattage,
			Purpose:              instrument.Purpose,
			Position:             instrument.Position,
			UnitNumber:           instrument.UnitNumber,
			Color:                instrument.Color,
			Dimmer:               instrument.Dimmer,
			Channel:              instrument.Channel,
			Universe:             instrument.Universe,
			UDimmer:              instrument.UDimmer,
			CircuitNumber:        instrument.CircuitNumber,
			CircuitName:          instrument.CircuitName,
			System:               instrument.System,
			UserField1:           instrument.UserField1,
			UserField2:           instrument.UserField2,
			UserField3:           instrument.UserField3,
			UserField4:           instrument.UserField4,
			UserField5:           instrument.UserField5,
			UserField6:           instrument.UserField6,
			NumChannels:          instrument.NumChannels,
			FrameSize:            instrument.FrameSize,
			FieldAngle:           instrument.FieldAngle,
			FieldAngle2:          instrument.FieldAngle2,
			BeamAngle:            instrument.BeamAngle,
			BeamAngle2:           instrument.BeamAngle2,
			Weight:               instrument.Weight,
			Gobo1:                instrument.Gobo1,
			Gobo1Rotation:        instrument.Gobo1Rotation,
			Gobo2:                instrument.Gobo2,
			Gobo2Rotation:        instrument.Gobo2Rotation,
			Mark:                 instrument.Mark,
			LampRotationAngle:    instrument.LampRotationAngle,
			Focus:                instrument.Focus,
			XRotation:            instrument.XRotation,
			YRotation:            instrument.YRotation,
			FixtureID:            instrument.FixtureID,
			UseVerticalBeam:      instrument.UseVerticalBeam,
			Address:              instrument.Address,
			Voltage:              instrument.Voltage,
			BreakerID:            instrument.BreakerID,
			Time:                 instrument.Time,
			Cost:                 instrument.Cost,
			VerticalFocusAngle:   instrument.VerticalFocusAngle,
			HorizontalFocusAngle: instrument.HorizontalFocusAngle,
			AngleToFace:          instrument.AngleToFace,
			OffAxisAngle:         instrument.OffAxisAngle,
			ThrowDistance:        instrument.ThrowDistance,
			FixtureMode:          instrument.FixtureMode,
			EnableZRotation:      instrument.EnableZRotation,
			ZRotation:            instrument.ZRotation,
			Accessories:          instrument.Accessories.Items,
			XLocmm:               instrument.XLocmm,
			YLocmm:               instrument.YLocmm,
			ZLocmm:               instrument.ZLocmm,
		}

		fixtures = append(fixtures, fixture)
	}

	return fixtures
}
