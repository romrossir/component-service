package component

// DefaultService provides default implementation of Service using a Repository.
type DefaultService struct {
	Repo Repository // Underlying data access implementation
}

// Create creates a new component.
func (s *DefaultService) Create(c *Component) error {
	return s.Repo.Create(c)
}

// Get retrieves a component by ID.
func (s *DefaultService) Get(id int64) (*Component, error) {
	return s.Repo.GetByID(id)
}

// Update updates an existing component.
func (s *DefaultService) Update(id int64, c *Component) error {
	return s.Repo.Update(id, c)
}

// Delete removes a component.
func (s *DefaultService) Delete(id int64) error {
	return s.Repo.Delete(id)
}

// List returns all components.
func (s *DefaultService) List() ([]Component, error) {
	return s.Repo.List()
}
