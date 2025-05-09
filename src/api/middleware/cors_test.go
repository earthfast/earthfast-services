package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORS(t *testing.T) {
	middleware := CORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}))

	rw := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	middleware.ServeHTTP(rw, r)

	wantHeaders := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "*",
		"Access-Control-Allow-Methods": "GET, HEAD, OPTIONS",
	}
	for name, wantVal := range wantHeaders {
		if gotVal := rw.Result().Header.Get(name); gotVal != wantVal {
			t.Errorf("Incorrect response header %q: got %q, want %q", name, gotVal, wantVal)
		}
	}
}
