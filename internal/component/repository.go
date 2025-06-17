package component

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) List() ([]Component, error) {
	rows, err := r.DB.Query("SELECT id, name, version FROM component")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var components []Component
	for rows.Next() {
		var c Component
		if err := rows.Scan(&c.ID, &c.Name, &c.Version); err != nil {
			return nil, err
		}
		components = append(components, c)
	}
	return components, nil
}

func (r *Repository) Get(id int) (*Component, error) {
	var c Component
	err := r.DB.QueryRow("SELECT id, name, version FROM component WHERE id = $1", id).
		Scan(&c.ID, &c.Name, &c.Version)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *Repository) Create(c *Component) error {
	return r.DB.QueryRow("INSERT INTO component (name, version) VALUES ($1, $2) RETURNING id", c.Name, c.Version).
		Scan(&c.ID)
}

func (r *Repository) Update(c *Component) error {
	_, err := r.DB.Exec("UPDATE component SET name = $1, version = $2 WHERE id = $3", c.Name, c.Version, c.ID)
	return err
}

func (r *Repository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM component WHERE id = $1", id)
	return err
}
