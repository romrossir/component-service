package component

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// Handler handles HTTP requests for component CRUD operations.
type Handler struct {
	Service Service // Business logic interface
}

// ServeHTTP dispatches requests to appropriate method handlers.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/components":
		h.list(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/components":
		h.create(w, r)
	case strings.HasPrefix(r.URL.Path, "/components/"):
		idStr := strings.TrimPrefix(r.URL.Path, "/components/")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid ID", http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			h.get(w, r, id)
		case http.MethodPut:
			h.update(w, r, id)
		case http.MethodDelete:
			h.delete(w, r, id)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	default:
		http.NotFound(w, r)
	}
}

// create handles POST /components.
func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var c Component
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if err := h.Service.Create(&c); err != nil {
		http.Error(w, "failed to create", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(c)
}

// get handles GET /components/{id}.
func (h *Handler) get(w http.ResponseWriter, _ *http.Request, id int64) {
	c, err := h.Service.Get(id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(c)
}

// update handles PUT /components/{id}.
func (h *Handler) update(w http.ResponseWriter, r *http.Request, id int64) {
	var c Component
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if err := h.Service.Update(id, &c); err != nil {
		http.Error(w, "update failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// delete handles DELETE /components/{id}.
func (h *Handler) delete(w http.ResponseWriter, _ *http.Request, id int64) {
	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "delete failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// list handles GET /components.
func (h *Handler) list(w http.ResponseWriter, _ *http.Request) {
	components, err := h.Service.List()
	if err != nil {
		http.Error(w, "failed to list", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(components)
}
