package pets

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newTestServer(t *testing.T) *httptest.Server {
	t.Helper()
	mux := http.NewServeMux()
	NewHandler(NewStore()).Register(mux)
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv
}

func do(t *testing.T, method, url, body string) *http.Response {
	t.Helper()
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { resp.Body.Close() })
	return resp
}

func TestPetLifecycle(t *testing.T) {
	srv := newTestServer(t)

	steps := []struct {
		method, path, body string
		want               int
	}{
		{"GET", "/pets", "", http.StatusOK},
		{"POST", "/pets", `{"name":"Rex","species":"dog"}`, http.StatusCreated},
		{"GET", "/pets/1", "", http.StatusOK},
		{"PUT", "/pets/1", `{"name":"Rexy","species":"dog"}`, http.StatusOK},
		{"DELETE", "/pets/1", "", http.StatusNoContent},
		{"GET", "/pets/1", "", http.StatusNotFound},
	}
	for _, s := range steps {
		resp := do(t, s.method, srv.URL+s.path, s.body)
		if resp.StatusCode != s.want {
			t.Fatalf("%s %s: got %d, want %d", s.method, s.path, resp.StatusCode, s.want)
		}
	}
}

func TestCreateValidation(t *testing.T) {
	srv := newTestServer(t)

	cases := map[string]struct {
		body string
		want int
	}{
		"malformed json": {`{`, http.StatusBadRequest},
		"missing name":   {`{"species":"cat"}`, http.StatusUnprocessableEntity},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			resp := do(t, "POST", srv.URL+"/pets", tc.body)
			if resp.StatusCode != tc.want {
				t.Fatalf("got %d, want %d", resp.StatusCode, tc.want)
			}
		})
	}
}

func TestInvalidID(t *testing.T) {
	srv := newTestServer(t)

	resp := do(t, "GET", srv.URL+"/pets/abc", "")
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("got %d, want %d", resp.StatusCode, http.StatusBadRequest)
	}
}
