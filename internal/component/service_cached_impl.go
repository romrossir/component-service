package component

import (
	"sync"
)

// CachedService wraps a Service and adds caching for Get and List.
type CachedService struct {
	Backend Service // Underlying service (e.g., DefaultService)

	mu        sync.RWMutex
	getCache  map[int64]*Component
	listCache []Component
	listValid bool
}

// NewCachedService initializes a new CachedService.
func NewCachedService(backend Service) *CachedService {
	return &CachedService{
		Backend:   backend,
		getCache:  make(map[int64]*Component),
		listCache: nil,
		listValid: false,
	}
}

// Create adds a new component and clears caches.
func (s *CachedService) Create(c *Component) error {
	if err := s.Backend.Create(c); err != nil {
		return err
	}
	s.invalidate()
	return nil
}

// Get returns a component by ID, using cache.
func (s *CachedService) Get(id int64) (*Component, error) {
	s.mu.RLock()
	cached, found := s.getCache[id]
	s.mu.RUnlock()
	if found {
		return cached, nil
	}

	// Cache miss: get from backend
	c, err := s.Backend.Get(id)
	if err != nil {
		return nil, err
	}

	s.mu.Lock()
	s.getCache[id] = c
	s.mu.Unlock()

	return c, nil
}

// Update modifies a component and clears caches.
func (s *CachedService) Update(id int64, c *Component) error {
	if err := s.Backend.Update(id, c); err != nil {
		return err
	}
	s.invalidate()
	return nil
}

// Delete removes a component and clears caches.
func (s *CachedService) Delete(id int64) error {
	if err := s.Backend.Delete(id); err != nil {
		return err
	}
	s.invalidate()
	return nil
}

// List returns all components, using cache.
func (s *CachedService) List() ([]Component, error) {
	s.mu.RLock()
	if s.listValid {
		cachedCopy := make([]Component, len(s.listCache))
		copy(cachedCopy, s.listCache)
		s.mu.RUnlock()
		return cachedCopy, nil
	}
	s.mu.RUnlock()

	// Cache miss
	list, err := s.Backend.List()
	if err != nil {
		return nil, err
	}

	s.mu.Lock()
	s.listCache = list
	s.listValid = true
	s.mu.Unlock()

	copied := make([]Component, len(list))
	copy(copied, list)
	return copied, nil
}

// invalidate clears the caches (called after writes).
func (s *CachedService) invalidate() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.getCache = make(map[int64]*Component)
	s.listCache = nil
	s.listValid = false
}
