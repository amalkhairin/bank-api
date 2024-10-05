package repositories

import (
	"encoding/json"
	"os"
	"time"

	"github.com/google/uuid"
	"bank-api/src/models"
)

type PaymentRepository struct {
	dataPath string
}

// NewPaymentRepository creates a new PaymentRepository
func NewPaymentRepository(dataPath string) *PaymentRepository {
	return &PaymentRepository{dataPath: dataPath}
}

// create payment
func (r *PaymentRepository) CreatePayment(payment *models.Payment) error {
	payments, err := r.loadPayments()
	if err != nil {
		return err
	}

	payment.ID = uuid.New().String()
	payment.CreatedAt = time.Now()

	payments = append(payments, *payment)
	return r.savePayments(payments)
}

// get user payments
func (r *PaymentRepository) GetPaymentsByUserID(userID string) ([]models.Payment, error) {
	payments, err := r.loadPayments()
	if err != nil {
		return nil, err
	}

	userPayments := []models.Payment{}
	for _, payment := range payments {
		if payment.UserID == userID {
			userPayments = append(userPayments, payment)
		}
	}

	return userPayments, nil
}

// load payment from json file
func (r *PaymentRepository) loadPayments() ([]models.Payment, error) {
	file, err := os.ReadFile(r.dataPath + "/payments.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Payment{}, nil
		}
		return nil, err
	}

	var payments []models.Payment
	err = json.Unmarshal(file, &payments)
	return payments, err
}

// save payment to json file
func (r *PaymentRepository) savePayments(payments []models.Payment) error {
	data, err := json.Marshal(payments)
	if err != nil {
		return err
	}

	return os.WriteFile(r.dataPath+"/payments.json", data, 0644)
}

func (r *PaymentRepository) AddHistory(history models.History) error {
	histories, err := r.loadHistories()
	if err != nil {
		return err
	}

	histories = append(histories, history)
	return r.saveHistories(histories)
}

// load history from json file
func (r *PaymentRepository) loadHistories() ([]models.History, error) {
	file, err := os.ReadFile(r.dataPath + "/history.json")
	if err != nil {
		return nil, err
	}

	var histories []models.History
	err = json.Unmarshal(file, &histories)
	return histories, err
}

// save history to json file
func (r *PaymentRepository) saveHistories(histories []models.History) error {
	data, err := json.Marshal(histories)
	if err != nil {
		return err
	}

	return os.WriteFile(r.dataPath+"/history.json", data, 0644)
}