package domain

import (
	"bytes"
	"fmt"
	"math"
	"net/http"
	"sort"
	"strings"

	"armada-node/api/middleware"
	"armada-node/geo"
	"armada-node/model"

	"go.uber.org/zap"
)

type swHandler struct {
	http.Handler

	m         model.Client
	resolver  Resolver
	templates Templates

	logger *zap.Logger
}

func newSWHandler(logger *zap.Logger, m model.Client, r Resolver, t Templates, c geo.GeolocationClient) *swHandler {
	h := &swHandler{
		m:         m,
		resolver:  r,
		templates: t,

		logger: logger,
	}

	h.Handler = middleware.GeoIP(logger, c, h.serveGeoHTTP)

	return h
}

func (h *swHandler) serveGeoHTTP(userCoord geo.Coordinate, w http.ResponseWriter, r *http.Request) {
	// Resolve the domain name being served into its associated project ID.
	projectID, err := h.resolver.ProjectForDomain(r.Context(), r.Host)
	if err != nil {
		h.logger.Error("Failed to resolve domain", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if projectID.IsZero() {
		h.logger.Error("Project not found for domain", zap.String("host", r.Host))
		http.Error(w, fmt.Sprintf("Project not found for domain: %s", r.Host), http.StatusNotFound)
		return
	}

	// Fetch all content nodes reserved by this project, across all regions.
	allNodes, err := h.m.ContentNodes(r.Context(), projectID)
	// print all nodes
	h.logger.Info("All nodes", zap.Any("nodes", allNodes))
	if err != nil {
		h.logger.Error("Failed to fetch content nodes", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Group the nodes by region.
	nodesByRegion := make(map[geo.Region][]*model.Node)
	for _, node := range allNodes {
		region := geo.GetRegion(node.Region)
		nodesByRegion[region] = append(nodesByRegion[region], node)
	}

	var chosenNodes []*model.Node
	if userCoord.IsZero() {
		// We were unable to geolocate the user so we can't pick a region.
		// Instead we choose to return the region that has the most nodes,
		// the idea being that it's likely the project's busiest region and
		// therefore the most likely to be where this user is visiting from.
		for _, nodes := range nodesByRegion {
			if len(nodes) > len(chosenNodes) {
				chosenNodes = nodes
			}
		}
	} else {
		// We were able to geolocate the user so we'll return the nodes from
		// the region they're closest to.
		minDistance := math.MaxFloat64
		for region, nodes := range nodesByRegion {
			dist := region.Distance(userCoord)
			if dist < minDistance {
				minDistance = dist
				chosenNodes = nodes
			}
		}
	}

	var contentNodes []string
	for _, node := range chosenNodes {
		contentNodes = append(contentNodes, node.Host)
	}

	// Sort the nodes. Any order is fine, but order matters because this list gets turned
	// into a string that's included in the returned service worker source.  If a browser
	// already has the service worker installed, it will compare the source code and consider
	// any difference to mean that a new version is available.
	sort.Strings(contentNodes)

	var buf bytes.Buffer
	data := struct {
		ProjectID      string
		BootstrapNodes string
		ContentNodes   string
	}{
		ProjectID:      projectID.Hex(),
		BootstrapNodes: "",
		ContentNodes:   strings.Join(contentNodes, ","),
	}
	if err := h.templates.ServiceWorker.Execute(&buf, data); err != nil {
		h.logger.Error("Failed to execute service worker template", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/javascript")
	if _, err := buf.WriteTo(w); err != nil {
		h.logger.Error("Failed to write response", zap.Error(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
