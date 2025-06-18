package component

// Repository defines the interface for accessing component data.
type Repository interface {
	Create(c *Component) error            // Create a new component
	GetByID(id int64) (*Component, error) // Retrieve a component by its ID
	Update(id int64, c *Component) error  // Update an existing component
	Delete(id int64) error                // Delete a component
	List() ([]Component, error)           // List all components
}
