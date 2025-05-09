package middleware

import (
	"net/http"

	"go.opencensus.io/plugin/ochttp"
)

func OpenCensus(next http.Handler) http.Handler {
	return &ochttp.Handler{
		Handler:          next,
		IsPublicEndpoint: true,
	}
}
