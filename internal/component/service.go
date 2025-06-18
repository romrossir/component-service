package component

// Service defines business logic operations for components.
type Service interface {
	Create(c *Component) error           // Create a new component
	Get(id int64) (*Component, error)    // Get a component by ID
	Update(id int64, c *Component) error // Update an existing component
	Delete(id int64) error               // Delete a component
	List() ([]Component, error)          // List all components
}
