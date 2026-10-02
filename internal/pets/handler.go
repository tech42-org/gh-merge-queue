package pets

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Handler exposes the pet store over HTTP.
type Handler struct {
	store *Store
}

// NewHandler returns a Handler backed by store.
func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

// Register mounts the pet routes on mux.
func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /pets", h.list)
	mux.HandleFunc("POST /pets", h.create)
	mux.HandleFunc("GET /pets/{id}", h.get)
	mux.HandleFunc("PUT /pets/{id}", h.update)
	mux.HandleFunc("DELETE /pets/{id}", h.delete)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	pets := h.store.List()
	writeJSON(w, http.StatusOK, pets)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	p, err := h.store.Get(id)
	if err != nil {
		writeStoreError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, p)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	p, ok := decodePet(w, r)
	if !ok {
		return
	}
	writeJSON(w, http.StatusCreated, h.store.Create(p))
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	p, ok := decodePet(w, r)
	if !ok {
		return
	}
	updated, err := h.store.Update(id, p)
	if err != nil {
		writeStoreError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}
	if err := h.store.Delete(id); err != nil {
		writeStoreError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

const maxNameLen = 50

func validate(p Pet) error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("name is required")
	}
	if len(p.Name) > maxNameLen {
		return fmt.Errorf("name must be at most %d characters", maxNameLen)
	}
	return nil
}

func decodePet(w http.ResponseWriter, r *http.Request) (Pet, bool) {
	var p Pet
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return Pet{}, false
	}
	if err := validate(p); err != nil {
		writeError(w, http.StatusUnprocessableEntity, err.Error())
		return Pet{}, false
	}
	return p, true
}

func parseID(w http.ResponseWriter, r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id < 1 {
		writeError(w, http.StatusBadRequest, "invalid pet id")
		return 0, false
	}
	return id, true
}

func writeStoreError(w http.ResponseWriter, err error) {
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeError(w, http.StatusInternalServerError, "internal error")
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
