package pets

import (
	"errors"
	"sort"
	"sync"
)

// ErrNotFound is returned when a pet does not exist.
var ErrNotFound = errors.New("pet not found")

// Store is an in-memory, concurrency-safe pet store.
type Store struct {
	mu     sync.RWMutex
	nextID int64
	pets   map[int64]Pet
}

// NewStore returns an empty store.
func NewStore() *Store {
	return &Store{nextID: 1, pets: make(map[int64]Pet)}
}

// All returns all pets ordered by ID.
func (s *Store) All() []Pet {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]Pet, 0, len(s.pets))
	for _, p := range s.pets {
		out = append(out, p)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

// Get returns the pet with the given ID.
func (s *Store) Get(id int64) (Pet, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	p, ok := s.pets[id]
	if !ok {
		return Pet{}, ErrNotFound
	}
	return p, nil
}

// Create stores a new pet and assigns it an ID.
func (s *Store) Create(p Pet) Pet {
	s.mu.Lock()
	defer s.mu.Unlock()

	p.ID = s.nextID
	s.nextID++
	s.pets[p.ID] = p
	return p
}

// Update replaces the pet with the given ID.
func (s *Store) Update(id int64, p Pet) (Pet, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.pets[id]; !ok {
		return Pet{}, ErrNotFound
	}
	p.ID = id
	s.pets[id] = p
	return p, nil
}

// Delete removes the pet with the given ID.
func (s *Store) Delete(id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.pets[id]; !ok {
		return ErrNotFound
	}
	delete(s.pets, id)
	return nil
}
