package order

type Notifier interface {
	NotifyAll(order any)
}

type Service struct {
	notifier Notifier
}

func NewService(notifier Notifier) *Service {
	return &Service{
		notifier: notifier,
	}
}

func (s *Service) ProcessOrder() {
	s.notifier.NotifyAll("order_processed")
}
