package pets

// Pet is a pet in the store.
type Pet struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Species string `json:"species"`
	Age     int    `json:"age"`
}
