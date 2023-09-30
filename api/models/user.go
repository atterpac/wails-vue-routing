package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id uuid.UUID `json:"id" gorm:"primaryKey;unique;not null"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Event []*Event `json:"events" gorm:"many2many:user_events;"`
	// Crud Tracking
	ModifiedAt time.Time `json:"modified_at"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Id == uuid.Nil {
		u.Id = uuid.New()
	}
	return nil
}


func GetUserById(db *gorm.DB, id uuid.UUID) (User, error) {
	var user User
	err := db.First(&user, id).Error
	if err != nil {
		log.Println("Error getting user: ", err)
		return User{}, err
	}
	return user, nil
}


func (u *User) GetId() uuid.UUID {
	return u.Id
}

func (u User) Create(db *gorm.DB) (User, error) {
	err := db.Create(&u).Error
	if err != nil {
		log.Println("Error creating user: ", err)
		return User{}, err
	}
	return u, nil	
}

func (u User) Update(db *gorm.DB) (User, error) {
	err := db.Save(&u).Error
	if err != nil {
		log.Println("Error updating user: ", err)
		return User{}, err
	}
	return u, nil	
}

func (u User) Delete(db *gorm.DB) (User, error) {
	err := db.Delete(&u).Error
	if err != nil {
		log.Println("Error deleting user: ", err)
		return User{}, err
	}
	return u, nil	
}

func DeleteById(db *gorm.DB, id uuid.UUID) (error) {
	err := db.Delete(&User{}, id).Error
	if err != nil {
		log.Println("Error deleting user: ", err)
		return err
	}
	return nil
}

func GetAllUsers(db *gorm.DB) ([]User, error){
	var users []User
	err := db.Find(&users).Error
	if err != nil {
		log.Println("Error getting all users: ", err)
		return []User{}, err
	}
	return users, nil
}
