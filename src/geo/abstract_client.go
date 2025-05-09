package geo

import (
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"
)

type AbstractClient struct {
	Client  *http.Client
	BaseURL string
	APIKey  string
	logger  *zap.Logger
}

type GeolocationClient interface {
	Get(ipAddr string) (Coordinate, error)
}

func NewAbstractClient(ipLookupAPIKey string, logger *zap.Logger) (AbstractClient, error) {
	if logger == nil {
		logger = zap.NewNop()
	}

	if ipLookupAPIKey == "" {
		logger.Warn("Missing IP lookup API key, geolocation will be unavailable")
		return AbstractClient{}, errors.New("missing IP lookup API key")
	}

	client := AbstractClient{
		Client: &http.Client{
			Timeout: 250 * time.Millisecond,
		},
		BaseURL: "https://ipgeolocation.abstractapi.com/v1/",
		APIKey:  ipLookupAPIKey,
		logger:  logger,
	}

	logger.Info("Initialized geolocation client",
		zap.String("provider", "AbstractAPI"),
		zap.Duration("timeout", client.Client.Timeout))

	return client, nil
}

func (c AbstractClient) Get(ipAddr string) (Coordinate, error) {
	if net.ParseIP(ipAddr).IsPrivate() || ipAddr == "127.0.0.1" {
		c.logger.Debug("Skipping geolocation for private/local IP", zap.String("ip", ipAddr))
		return Coordinate{}, nil
	}

	u, err := url.Parse(c.BaseURL)
	if err != nil {
		return Coordinate{}, err
	}

	q := u.Query()
	q.Set("api_key", c.APIKey)
	q.Set("ip_address", ipAddr)
	u.RawQuery = q.Encode()

	// Redact API key from logs
	logURL := *u
	logQuery := logURL.Query()
	logQuery.Set("api_key", "[REDACTED]")
	logURL.RawQuery = logQuery.Encode()

	c.logger.Debug("Making geolocation request",
		zap.String("url", logURL.String()),
		zap.String("ip", ipAddr))

	resp, err := c.Client.Get(u.String())
	if err != nil {
		return Coordinate{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Coordinate{}, errors.New("received non-OK HTTP status")
	}

	var info Coordinate
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return Coordinate{}, err
	}

	if info.Latitude == 0 && info.Longitude == 0 {
		return Coordinate{}, errors.New("no coordinates returned")
	}

	return info, nil
}
