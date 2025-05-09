package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"armada-node/geo"

	"go.uber.org/zap"
)

const (
	minLatitude  = -90.0
	maxLatitude  = 90.0
	minLongitude = -180.0
	maxLongitude = 180.0
)

const (
	ErrMissingLatitude  = "missing X-Geoip-Latitude"
	ErrMissingLongitude = "missing X-Geoip-Longitude"
)

type GeoIPHandler func(geo.Coordinate, http.ResponseWriter, *http.Request)

func GeoIP(logger *zap.Logger, c geo.GeolocationClient, next GeoIPHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		coord, err := extractCoordinateFromHeaders(r)
		if err != nil {
			logger.Warn("Geolocation failed", zap.Error(err))

			// if missing Geoip headers, try to find IP with fallback API
			if err.Error() == ErrMissingLatitude || err.Error() == ErrMissingLongitude {
				coord, err = lookupLatAndLon(r.RemoteAddr, c)
				if err != nil {
					logger.Warn("IP Lookup failed", zap.Error(err))
				}
			}
		}

		err = validateCoordinate(coord)
		if err != nil {
			logger.Warn("Invalid coordinate", zap.Error(err))
			// set coord to zero value if invalid
			coord = geo.Coordinate{}
		}

		next(coord, w, r)
	})
}

func lookupLatAndLon(ipAddr string, c geo.GeolocationClient) (geo.Coordinate, error) {
	coord, err := c.Get(ipAddr)

	if err != nil {
		return geo.Coordinate{}, err
	}

	return coord, nil
}

func extractCoordinateFromHeaders(r *http.Request) (geo.Coordinate, error) {
	var coord geo.Coordinate

	if xLatitude := r.Header.Get("X-Geoip-Latitude"); xLatitude == "" {
		return geo.Coordinate{}, errors.New(ErrMissingLatitude)
	} else if lat, err := strconv.ParseFloat(xLatitude, 64); err != nil {
		return geo.Coordinate{}, fmt.Errorf("parsing X-Geoip-Latitude: %v", err)
	} else {
		coord.Latitude = lat
	}

	if xLongitude := r.Header.Get("X-Geoip-Longitude"); xLongitude == "" {
		return geo.Coordinate{}, errors.New(ErrMissingLongitude)
	} else if lon, err := strconv.ParseFloat(xLongitude, 64); err != nil {
		return geo.Coordinate{}, fmt.Errorf("parsing X-Geoip-Longitude: %v", err)
	} else {
		coord.Longitude = lon
	}

	return coord, nil
}

func validateCoordinate(coord geo.Coordinate) error {
	if coord.Latitude < minLatitude || coord.Latitude > maxLatitude {
		return fmt.Errorf("latitude is out of bounds: %f", coord.Latitude)
	}

	if coord.Longitude < minLongitude || coord.Longitude > maxLongitude {
		return fmt.Errorf("longitude is out of bounds: %f", coord.Longitude)
	}

	return nil
}
