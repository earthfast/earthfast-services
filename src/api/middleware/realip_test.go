package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRealIP(t *testing.T) {
	cases := []struct {
		name       string
		headers    map[string]string
		remoteAddr string
		wantIP     string
	}{
		{
			name:   "No IP",
			wantIP: "",
		},
		{
			name: "X-Real-IP",
			headers: map[string]string{
				"X-Real-IP":       "1.1.1.1",
				"X-Forwarded-For": "2.2.2.2",
			},
			remoteAddr: "3.3.3.3",
			wantIP:     "1.1.1.1",
		},
		{
			name: "X-Forwarded-For",
			headers: map[string]string{
				"X-Forwarded-For": "2.2.2.2",
			},
			remoteAddr: "3.3.3.3",
			wantIP:     "2.2.2.2",
		},
		{
			name: "X-Forwarded-For First",
			headers: map[string]string{
				"X-Forwarded-For": "1.1.1.1,2.2.2.2",
			},
			remoteAddr: "3.3.3.3",
			wantIP:     "1.1.1.1",
		},
		{
			name:       "RemoteAddr",
			remoteAddr: "3.3.3.3",
			wantIP:     "3.3.3.3",
		},
		{
			name:       "RemoteAddr host:port",
			remoteAddr: "3.3.3.3:999",
			wantIP:     "3.3.3.3",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			middleware := RealIP(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.RemoteAddr != tc.wantIP {
					t.Errorf("Incorrect IP: got %s, want %s", r.RemoteAddr, tc.wantIP)
				}
			}))

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			for key, val := range tc.headers {
				r.Header.Set(key, val)
			}
			r.RemoteAddr = tc.remoteAddr

			middleware.ServeHTTP(w, r)
		})
	}
}
