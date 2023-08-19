package service_test

import (
	"course/pkg/service"
	"testing"
)

type mockUserService struct {
	f func(userID string) error
}

func (s *mockUserService) UpdateUserStatus(userID string) error {
	return s.f(userID)
}

func TestCreateOrder(t *testing.T) {
	inputCmd := service.CreateOrderCmd{
		UserID:     "user1",
		PromoCode:  "promo1",
		ProductIDs: []string{"prod1", "prod2"},
	}

	mock := &mockUserService{
		f: func(userID string) error {
			if userID != inputCmd.UserID {
				t.Logf("expected userID %s, got %s", inputCmd.UserID, userID)
				t.Fail()
			}

			return nil
		},
	}

	orderService := service.NewOrderService(mock)

	wantErr := false

	err := orderService.CreateOrder(inputCmd)
	if err != nil == !wantErr {
		t.Logf("expected error %v, got %v", wantErr, err)
		t.Fail()
	}
}
