package component

// Component represents a node in a hierarchical structure.
type Component struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID *int   `json:"parent_id,omitempty"`
}
