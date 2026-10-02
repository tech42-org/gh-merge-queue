package pets

import (
	"errors"
	"testing"
)

func TestStoreCreateAssignsIncrementingIDs(t *testing.T) {
	s := NewStore()
	a := s.Create(Pet{Name: "A"})
	b := s.Create(Pet{Name: "B"})
	if a.ID != 1 || b.ID != 2 {
		t.Fatalf("got ids %d, %d; want 1, 2", a.ID, b.ID)
	}
}

func TestStoreGetMissing(t *testing.T) {
	s := NewStore()
	if _, err := s.Get(42); !errors.Is(err, ErrNotFound) {
		t.Fatalf("got %v, want ErrNotFound", err)
	}
}

func TestStoreUpdateKeepsID(t *testing.T) {
	s := NewStore()
	p := s.Create(Pet{Name: "A"})
	got, err := s.Update(p.ID, Pet{ID: 999, Name: "B"})
	if err != nil {
		t.Fatal(err)
	}
	if got.ID != p.ID || got.Name != "B" {
		t.Fatalf("got %+v", got)
	}
}

func TestStoreDelete(t *testing.T) {
	s := NewStore()
	p := s.Create(Pet{Name: "A"})
	if err := s.Delete(p.ID); err != nil {
		t.Fatal(err)
	}
	if err := s.Delete(p.ID); !errors.Is(err, ErrNotFound) {
		t.Fatalf("second delete: got %v, want ErrNotFound", err)
	}
}
