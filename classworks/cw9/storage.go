package main

type Storage struct {
	cities []string
}

func (s *Storage) CreateCity(city string) {
	s.cities = append(s.cities, city)
}

func (s *Storage) GetCities() []string {
	return s.cities
}
