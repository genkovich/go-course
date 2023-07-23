package pubsub

type Subscriber interface {
	GetNotified(event any) error
	GetId() string
}

type message struct {
	event   any
	errChan chan<- error
}

type Publisher interface {
	PublishEvent() any
}

type Service struct {
	subscribers   chan Subscriber
	eventsChannel chan message
	stop          chan struct{}
}

func NewService() *Service {
	s := &Service{
		subscribers:   make(chan Subscriber),
		eventsChannel: make(chan message),
		stop:          make(chan struct{}),
	}

	go s.ProcessEvents()

	return s
}

func (s *Service) ProcessEvents() {
	subscribers := make(map[string]Subscriber)

	for {
		select {
		case e := <-s.eventsChannel:
			for _, sub := range subscribers {
				if err := sub.GetNotified(e.event); err != nil {
					e.errChan <- err
				}
			}
			close(e.errChan)
		case sub := <-s.subscribers:
			subscribers[sub.GetId()] = sub
		case <-s.stop:
			s.stop <- struct{}{}
			return
		}
	}
}

func (s *Service) AddSub(sub Subscriber) {
	s.subscribers <- sub
}

func (s *Service) Publish(event any) <-chan error {
	errCh := make(chan error)
	
	s.eventsChannel <- message{
		event:   event,
		errChan: errCh,
	}

	return errCh
}

func (s *Service) Stop() {
	s.stop <- struct{}{}
	<-s.stop
}
