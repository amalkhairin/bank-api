package repositories

import (
	"encoding/json"
	"errors"
	"os"
	
	"bank-api/src/models"
)

type UserRepository struct {
	dataPath string
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(dataPath string) *UserRepository {
	return &UserRepository{dataPath: dataPath}
}

// get user
func (r *UserRepository) GetUser(id string) (*models.User, error) {
	users, err := r.loadUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	users, err := r.loadUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}

// update user
func (r *UserRepository) UpdateUser(user *models.User) error {
	users, err := r.loadUsers()
	if err != nil {
		return err
	}

	for i, u := range users {
		if u.ID == user.ID {
			users[i] = *user
			return r.saveUsers(users)
		}
	}

	return errors.New("user not found")
}

// add history
func (r *UserRepository) AddHistory(history models.History) error {
	histories, err := r.loadHistories()
	if err != nil {
		return err
	}

	histories = append(histories, history)
	return r.saveHistories(histories)
}

// load user from json file
func (r *UserRepository) loadUsers() ([]models.User, error) {
	file, err := os.ReadFile(r.dataPath + "/users.json")
	if err != nil {
		return nil, err
	}



	var users []models.User
	err = json.Unmarshal(file, &users)
	return users, err
}

// save user to json file
func (r *UserRepository) saveUsers(users []models.User) error {
	data, err := json.Marshal(users)
	if err != nil {
		return err
	}

	return os.WriteFile(r.dataPath+"/users.json", data, 0644)
}

// load history from json file
func (r *UserRepository) loadHistories() ([]models.History, error) {
	file, err := os.ReadFile(r.dataPath + "/history.json")
	if err != nil {
		return nil, err
	}

	var histories []models.History
	err = json.Unmarshal(file, &histories)
	return histories, err
}

// save history to json file
func (r *UserRepository) saveHistories(histories []models.History) error {
	data, err := json.Marshal(histories)
	if err != nil {
		return err
	}

	return os.WriteFile(r.dataPath+"/history.json", data, 0644)
}
