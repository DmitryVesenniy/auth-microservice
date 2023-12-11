package repository

import (
	"fmt"

	"github.com/DmitryVesenniy/auth-microservice/internal/domain/models"
)

type RepositoryInMemory struct {
	lastIndex int64
	Users     []models.User
}

func (r *RepositoryInMemory) CreateUser(user models.User) (models.User, error) {
	_, err := r.GetUser(user.Email)

	if err == nil {
		return models.User{}, fmt.Errorf("user already exist")
	}

	r.lastIndex++

	user.ID = r.lastIndex

	r.Users = append(r.Users, user)

	return user, nil
}

func (r *RepositoryInMemory) GetUser(email string) (models.User, error) {
	for _, user := range r.Users {
		if user.Email == email {
			return user, nil
		}
	}

	return models.User{}, fmt.Errorf("user not found")
}

func (r *RepositoryInMemory) ListUser(email string) ([]models.User, error) {
	return r.Users, nil
}

func (r *RepositoryInMemory) DelUser(userID int64) error {

	index := -1

	for i, user := range r.Users {
		if user.ID == userID {
			index = i
			break
		}
	}

	if index < 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}
