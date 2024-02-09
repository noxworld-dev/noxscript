package nstest

import "encoding/json"

type Storage struct {
	Sessions    map[string]*Store
	Persistents map[string]*Store
}

func (s *Storage) Session(name string) *Store {
	st := s.Sessions[name]
	if st != nil {
		return st
	}
	st = &Store{Persistent: false, Name: name}
	if s.Sessions == nil {
		s.Sessions = make(map[string]*Store)
	}
	s.Sessions[name] = st
	return st
}

func (s *Storage) Persistent(name string) *Store {
	st := s.Persistents[name]
	if st != nil {
		return st
	}
	st = &Store{Persistent: true, Name: name}
	if s.Persistents == nil {
		s.Persistents = make(map[string]*Store)
	}
	s.Persistents[name] = st
	return st
}

type Store struct {
	Persistent bool
	Name       string
	Data       map[string][]byte
}

func (s *Store) Get(key string, obj any) error {
	data := s.Data[key]
	if len(data) == 0 {
		return nil
	}
	if err := json.Unmarshal(data, obj); err != nil {
		panic(err)
	}
	return nil
}

func (s *Store) Set(key string, obj any) error {
	data, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	if s.Data == nil {
		s.Data = make(map[string][]byte)
	}
	s.Data[key] = data
	return nil
}
