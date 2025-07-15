package entities_test

import (
	"testing"

	"github.com/otaviomart1ns/backend-challenge/internal/domain/entities"
)

func TestPaymentType_IsValid(t *testing.T) {
	tests := []struct {
		name     string
		payment  entities.PaymentType
		expected bool
	}{
		{
			name:     "Valid CREDIT_CARD",
			payment:  entities.CreditCard,
			expected: true,
		},
		{
			name:     "Invalid DEBIT_CARD",
			payment:  "DEBIT_CARD",
			expected: false,
		},
		{
			name:     "Empty string",
			payment:  "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.payment.IsValid()
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
