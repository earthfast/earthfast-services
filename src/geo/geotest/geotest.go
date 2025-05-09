package geotest

import "armada-node/geo"

var (
	London  = geo.Coordinate{Latitude: 51.507351, Longitude: -0.127758}
	NewYork = geo.Coordinate{Latitude: 40.730610, Longitude: -73.935242}
	Sydney  = geo.Coordinate{Latitude: -33.868820, Longitude: 151.209290}
	Tokyo   = geo.Coordinate{Latitude: 35.689487, Longitude: 139.691711}
)

type AbstractClientMock struct {
	geo.AbstractClient
	geo.Coordinate
}

func NewAbstractClientMock(lat float64, lon float64) AbstractClientMock {
	return AbstractClientMock{geo.AbstractClient{}, geo.Coordinate{Latitude: lat, Longitude: lon}}
}

func (c AbstractClientMock) Get(ipAddr string) (geo.Coordinate, error) {
	return c.Coordinate, nil
}
