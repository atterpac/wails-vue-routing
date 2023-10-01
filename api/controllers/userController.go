package controllers

import (
	"log"
	"yagami/api/models"

	"github.com/google/uuid"
)

func (a *App) GetUsers() []models.User {
	u, err := models.GetAllUsers(a.Db)
	if err != nil {
		log.Println("Error Getting Users: ", err)
		return nil
	}
	return u
}

func (a *App) UpdateUser(user models.User) models.User {
	updateUser, err := user.Update(a.Db)
	if err != nil {
		log.Println("Error Updating User: ", err)
		return user
	}
	return updateUser
}

func (a *App) CreateUser(user models.User) models.User {
	createdUser, err := user.Create(a.Db)
	if err != nil {
		log.Println("Error Creating User: ", err)
		return user
	}
	return createdUser
}

func (a *App) DeleteUser(user models.User) models.User {
	deletedUser, err := user.Delete(a.Db)
	if err != nil {
		log.Println("Error Deleting User: ", err)
		return user
	}
	return deletedUser
}

func (a *App) GetUserById(id uuid.UUID) models.User {
	user, err := models.GetUserById(a.Db, id)
	if err != nil {
		log.Println("Error Getting User By Id: ", err)
		return models.User{}
	}
	return user
}

