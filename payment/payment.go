package payment

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// PaymentProcessor определяет общий интерфейс для всех платежных провайдеров
type PaymentProcessor interface {
	Authorize(amount float64) (string, error)                       // Авторизация платежа
	Capture(transactionID string) error                             // Подтверждение платежа
	Refund(transactionID string, amount float64) error              // Возврат средств
	GetPaymentDetails(transactionID string) (PaymentDetails, error) // Получение деталей платежа
}

// PaymentDetails содержит информацию о платеже
type PaymentDetails struct {
	ID              string
	Amount          float64
	Currency        string
	Status          string
	CreatedAt       time.Time
	MethodOfPayment string
}

// Реализация StripeProcessor
type StripeProcessor struct {
	APIKey string
}

// Authorize создает транзакцию в Stripe
func (s *StripeProcessor) Authorize(amount float64) (string, error) {
	if amount <= 0 {
		return "", errors.New("amount must be positive")
	}
	if s.APIKey == "" {
		return "", errors.New("invalid API key")
	}
	return fmt.Sprintf("stripe_%d", time.Now().UnixNano()), nil
}

func (s *StripeProcessor) Capture(transactionID string) error {
	if transactionID == "" {
		return errors.New("transaction ID cannot be empty")
	}
	log.Printf("Capturing Stripe transaction: %s", transactionID)
	return nil
}

// Refund выполняет возврат средств
func (s *StripeProcessor) Refund(transactionID string, amount float64) error {
	if transactionID == "" {
		return errors.New("transaction ID cannot be empty")
	}
	if amount <= 0 {
		return errors.New("refund amount must be positive")
	}
	log.Printf("Refunding %.2f for Stripe transaction: %s", amount, transactionID)
	return nil
}

func ProcessPayment(processor PaymentProcessor, amount float64) (string, error) {
	transactionID, err := processor.Authorize(amount)
	if err != nil {
		return "", fmt.Errorf("authorization failed: %v", err)
	}

	if err := processor.Capture(transactionID); err != nil {
		return "", fmt.Errorf("capture failed: %v", err)
	}

	return transactionID, nil
}

func (s *StripeProcessor) GetPaymentDetails(transactionID string) (PaymentDetails, error) {
	return PaymentDetails{
		ID:              transactionID,
		Amount:          100.00, // Пример значения
		Currency:        "USD",
		Status:          "completed",
		CreatedAt:       time.Now(),
		MethodOfPayment: "stripe",
	}, nil
}
