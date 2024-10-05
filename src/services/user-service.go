package services

import (
	"errors"
	"time"

	"bank-api/src/models"
	"bank-api/src/repositories"
	"bank-api/src/utils"
)

type UserService struct {
	repo *repositories.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// login service
func (s *UserService) Login(username, password string) (string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	s.repo.AddHistory(models.History{
		UserID: user.ID,
		Action: "login",
		Time:   time.Now().Format(time.RFC3339),
	})

	return token, nil
}

// func (s *UserService) Payment(userID string, amount float64) error {
// 	user, err := s.repo.GetUser(userID)
// 	if err != nil {
// 		return err
// 	}

// 	if user.Balance < amount {
// 		return errors.New("insufficient balance")
// 	}

// 	user.Balance -= amount
// 	err = s.repo.UpdateUser(user)
// 	if err != nil {
// 		return err
// 	}

// 	s.repo.AddHistory(models.History{
// 		UserID: user.ID,
// 		Action: "payment",
// 		Amount: amount,
// 		Time:   time.Now().Format(time.RFC3339),
// 	})

// 	return nil
// }

// logout service
func (s *UserService) Logout(userID string) error {
	_, err := s.repo.GetUser(userID)
	if err != nil {
		return err
	}

	s.repo.AddHistory(models.History{
		UserID: userID,
		Action: "logout",
		Time:   time.Now().Format(time.RFC3339),
	})

	return nil
}
