package middleware

import (
	"armada-node/geo"
	"armada-node/geo/geotest"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap/zaptest"
)

func TestGeoIP(t *testing.T) {
	cases := []struct {
		name                  string
		headers               map[string]string
		want                  geo.Coordinate
		seedGeolocationClient geo.Coordinate
	}{
		{
			name:                  "No geo headers",
			headers:               map[string]string{},
			want:                  geo.Coordinate{},
			seedGeolocationClient: geo.Coordinate{},
		},
		{
			name: "Malformed latitude",
			headers: map[string]string{
				"X-Geoip-Latitude":  "abc",
				"X-Geoip-Longitude": "1.234",
			},
			want:                  geo.Coordinate{},
			seedGeolocationClient: geo.Coordinate{},
		},
		{
			name: "Malformed longitude",
			headers: map[string]string{
				"X-Geoip-Latitude":  "1.234",
				"X-Geoip-Longitude": "0-1",
			},
			want:                  geo.Coordinate{},
			seedGeolocationClient: geo.Coordinate{},
		},
		{
			name: "Invalid latitude",
			headers: map[string]string{
				"X-Geoip-Latitude":  "100.234",
				"X-Geoip-Longitude": "1.234",
			},
			want:                  geo.Coordinate{},
			seedGeolocationClient: geo.Coordinate{},
		},
		{
			name: "Invalid longitude",
			headers: map[string]string{
				"X-Geoip-Latitude":  "1.234",
				"X-Geoip-Longitude": "-270.234",
			},
			want:                  geo.Coordinate{},
			seedGeolocationClient: geo.Coordinate{},
		},
		{
			name:    "Missing headers but retrieved from lookupLatAndLon",
			headers: map[string]string{},
			want: geo.Coordinate{
				Latitude:  1.234,
				Longitude: -5.678,
			},
			seedGeolocationClient: geo.Coordinate{
				Latitude:  1.234,
				Longitude: -5.678,
			},
		},
		{
			name: "Success",
			headers: map[string]string{
				"X-Geoip-Latitude":  "1.234",
				"X-Geoip-Longitude": "-5.678",
			},
			want: geo.Coordinate{
				Latitude:  1.234,
				Longitude: -5.678,
			},
			seedGeolocationClient: geo.Coordinate{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			logger := zaptest.NewLogger(t)
			abstractClientMock := geotest.NewAbstractClientMock(tc.seedGeolocationClient.Latitude, tc.seedGeolocationClient.Longitude)
			middleware := GeoIP(logger, abstractClientMock, func(coord geo.Coordinate, w http.ResponseWriter, r *http.Request) {
				if coord != tc.want {
					t.Errorf("Unexpected Coordinate: got %+v, want %+v", coord, tc.want)
				}
			})

			rw := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/", nil)
			for key, val := range tc.headers {
				r.Header.Set(key, val)
			}

			middleware.ServeHTTP(rw, r)
		})
	}
}
