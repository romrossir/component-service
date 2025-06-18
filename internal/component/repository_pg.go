package component

import (
	"database/sql"
)

// PostgresRepository implements the Repository interface using PostgreSQL.
type PostgresRepository struct {
	DB *sql.DB
}

// Create inserts a new component into the database.
func (r *PostgresRepository) Create(c *Component) error {
	return r.DB.QueryRow(
		"INSERT INTO component (name, parent_id) VALUES ($1, $2) RETURNING id",
		c.Name, c.ParentID,
	).Scan(&c.ID)
}

// GetByID retrieves a component by ID.
func (r *PostgresRepository) GetByID(id int64) (*Component, error) {
	var c Component
	err := r.DB.QueryRow("SELECT id, name, parent_id FROM component WHERE id = $1", id).
		Scan(&c.ID, &c.Name, &c.ParentID)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// Update modifies a component with the given ID.
func (r *PostgresRepository) Update(id int64, c *Component) error {
	_, err := r.DB.Exec("UPDATE component SET name = $1, parent_id = $2 WHERE id = $3", c.Name, c.ParentID, id)
	return err
}

// Delete removes a component from the database.
func (r *PostgresRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM component WHERE id = $1", id)
	return err
}

// List retrieves all components.
func (r *PostgresRepository) List() ([]Component, error) {
	rows, err := r.DB.Query("SELECT id, name, parent_id FROM component")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var components []Component
	for rows.Next() {
		var c Component
		if err := rows.Scan(&c.ID, &c.Name, &c.ParentID); err != nil {
			return nil, err
		}
		components = append(components, c)
	}
	return components, nil
}
