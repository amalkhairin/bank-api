package services

import (
	"errors"
	"time"
	"bank-api/src/models"
	"bank-api/src/repositories"
)

type PaymentService struct {
	paymentRepo *repositories.PaymentRepository
	userRepo    *repositories.UserRepository
}

// NewPaymentService creates a new PaymentService
func NewPaymentService(paymentRepo *repositories.PaymentRepository, userRepo *repositories.UserRepository) *PaymentService {
	return &PaymentService{
		paymentRepo: paymentRepo,
		userRepo:    userRepo,
	}
}

// create payment
func (s *PaymentService) CreatePayment(userID string, amount float64, description string) error {
	user, err := s.userRepo.GetUser(userID)
	if err != nil {
		return err
	}

	if user.Balance < amount {
		return errors.New("insufficient balance")
	}

	payment := &models.Payment{
		UserID:      userID,
		Amount:      amount,
		Description: description,
	}

	err = s.paymentRepo.CreatePayment(payment)
	if err != nil {
		return err
	}

	user.Balance -= amount
	err = s.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}

	s.paymentRepo.AddHistory(models.History{
		UserID: user.ID,
		Action: "payment",
		Time:   time.Now().Format(time.RFC3339),
	})

	return nil
}

// get user payments
func (s *PaymentService) GetUserPayments(userID string) ([]models.Payment, error) {
	return s.paymentRepo.GetPaymentsByUserID(userID)
}
