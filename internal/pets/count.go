package pets

import "net/http"

func (h *Handler) count(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]int{"count": len(h.store.List())})
}
