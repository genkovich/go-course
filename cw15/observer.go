package cw15

type Registry struct {
	observers map[string]Observer
}

type Register interface {
	Register(o Observer)
	Unregister(o Observer)
	NotifyAll(subject any)
}

type Observer interface {
	GetNotified(subject any)
	GetID() string
}

func NewRegistry() *Registry {
	return &Registry{
		observers: make(map[string]Observer),
	}
}

func (s *Registry) NotifyAll(subject any) {
	for _, o := range s.observers {
		o.GetNotified(subject)
	}
}

func (s *Registry) Register(o Observer) {
	s.observers[o.GetID()] = o
}

func (s *Registry) Unregister(o Observer) {
	delete(s.observers, o.GetID())
}
