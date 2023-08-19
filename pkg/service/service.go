package service

import "fmt"

type Service interface{}

type UserService interface {
	UpdateUserStatus(userID string) error
}

type OrderService struct {
	userService UserService
}

func NewOrderService(userService UserService) *OrderService {
	return &OrderService{userService: userService}
}

type CreateOrderCmd struct {
	UserID     string
	PromoCode  string
	ProductIDs []string
}

func (s *OrderService) CreateOrder(cmd CreateOrderCmd) error {
	if len(cmd.ProductIDs) == 0 {
		return fmt.Errorf("no products in order")
	}

	// order creation logic...

	err := s.userService.UpdateUserStatus(cmd.UserID)

	if err != nil {
		return fmt.Errorf("error updating user status")
	}

	return nil
}
