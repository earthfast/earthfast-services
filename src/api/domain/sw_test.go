package domain

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"

	"armada-node/geo"
	"armada-node/geo/geotest"
	"armada-node/model"
	"armada-node/model/modeltest"

	"go.uber.org/zap/zaptest"
)

var (
	fakeTemplates = Templates{
		ServiceWorker: template.Must(template.New("main.js.tmpl").Parse(`
		var projectID = "{{.ProjectID}}";
		var bootstrapNodes = "{{.BootstrapNodes}}";
		var contentNodes = "{{.ContentNodes}}";
	`)),
	}

	erroringTemplates = Templates{
		ServiceWorker: template.Must(template.New("main.js.tmpl").Parse(`
		var foo = "{{.DoesNotExist}}";
	`)),
	}
)

func TestServiceWorkerHandler(t *testing.T) {
	logger := zaptest.NewLogger(t)

	project := &model.Project{ID: model.ID{1, 2, 3}, Name: "hello-world"}
	node0 := &model.Node{
		Host:      "node0.armadanetwork.com",
		Region:    geo.NorthAmerica.ID,
		ProjectID: project.ID,
	}
	node1 := &model.Node{
		Host:      "node1.armadanetwork.com",
		Region:    geo.NorthAmerica.ID,
		ProjectID: project.ID,
	}
	m := modeltest.NewClient().WithContentNodes(node0, node1).WithProjects(project)

	resolver := NewStaticResolver(map[string]model.ID{"example.com": project.ID})

	cases := []struct {
		name         string
		templates    Templates
		url          string
		headers      http.Header
		wantCode     int
		wantContains []string
	}{
		{
			name:      "Service worker",
			templates: fakeTemplates,
			url:       "http://example.com/earthfast-sw.js",
			wantCode:  http.StatusOK,
			wantContains: []string{
				fmt.Sprintf("contentNodes = \"%s,%s\"", node0.Host, node1.Host),
				fmt.Sprintf("projectID = \"%x\"", project.ID),
			},
		},
		{
			name:      "Service worker unknown domain",
			templates: fakeTemplates,
			url:       "http://foo.com/earthfast-sw.js",
			wantCode:  http.StatusNotFound,
		},
		{
			name:      "Service worker template failure",
			templates: erroringTemplates,
			url:       "http://example.com/earthfast-sw.js",
			wantCode:  http.StatusInternalServerError,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			h := newSWHandler(logger, m, resolver, tc.templates, &geotest.AbstractClientMock{})

			rw := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, tc.url, nil)
			r.Header = tc.headers
			h.ServeHTTP(rw, r)
			if rw.Code != tc.wantCode {
				t.Fatalf("Unexpected HTTP status code: got %d, want %d", rw.Code, tc.wantCode)
			}
			for _, want := range tc.wantContains {
				if got := rw.Body.String(); !strings.Contains(got, want) {
					t.Errorf("Unexpected response body: got %q, want it to contain %q", got, want)
				}
			}
		})
	}
}

func TestHandler_ServiceWorker_Geo(t *testing.T) {
	logger := zaptest.NewLogger(t)

	project0 := &model.Project{ID: modeltest.RandomID(t), Name: "hello-world"}
	europe0 := &model.Node{
		Host:      "node0.armadanetwork.com",
		Region:    geo.Europe.ID,
		ProjectID: project0.ID,
	}
	europe1 := &model.Node{
		Host:      "node1.armadanetwork.com",
		Region:    geo.Europe.ID,
		ProjectID: project0.ID,
	}
	northAmerica := &model.Node{
		Host:      "node2.armadanetwork.com",
		Region:    geo.NorthAmerica.ID,
		ProjectID: project0.ID,
	}
	asia := &model.Node{
		Host:      "node3.armadanetwork.com",
		Region:    geo.Asia.ID,
		ProjectID: project0.ID,
	}
	unknown0 := &model.Node{
		Host:      "node4.armadanetwork.com",
		Region:    "",
		ProjectID: project0.ID,
	}
	unknown1 := &model.Node{
		Host:      "node5.armadanetwork.com",
		Region:    "foo",
		ProjectID: project0.ID,
	}

	resolver := NewStaticResolver(map[string]model.ID{"example.com": project0.ID})

	templates := Templates{
		ServiceWorker: template.Must(template.New("main.js.tmpl").Parse("{{.ContentNodes}}")),
	}

	cases := []struct {
		name      string
		nodes     []*model.Node
		userCoord geo.Coordinate
		wantNodes []string
	}{
		{
			name:  "No user location returns largest region",
			nodes: []*model.Node{europe0, europe1, northAmerica, asia},
			wantNodes: []string{
				europe0.Host,
				europe1.Host,
			},
		},
		{
			name:  "Unknown regions are grouped together",
			nodes: []*model.Node{unknown0, unknown1},
			wantNodes: []string{
				unknown0.Host,
				unknown1.Host,
			},
		},
		{
			name:      "Nearest region is chosen (London)",
			nodes:     []*model.Node{europe0, europe1, northAmerica, asia, unknown0},
			userCoord: geotest.London,
			wantNodes: []string{
				europe0.Host,
				europe1.Host,
			},
		},
		{
			name:      "Nearest region is chosen (Tokyo)",
			nodes:     []*model.Node{europe0, europe1, northAmerica, asia, unknown0},
			userCoord: geotest.Tokyo,
			wantNodes: []string{
				asia.Host,
			},
		},
		{
			name:      "Nearest region is chosen (New York)",
			nodes:     []*model.Node{europe0, europe1, northAmerica, asia, unknown0},
			userCoord: geotest.NewYork,
			wantNodes: []string{
				northAmerica.Host,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := modeltest.NewClient().WithContentNodes(tc.nodes...).WithProjects(project0)
			h := newSWHandler(logger, m, resolver, templates, &geotest.AbstractClientMock{})

			rw := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://example.com/earthfast-sw.js", nil)
			if !tc.userCoord.IsZero() {
				r.Header.Set("X-Geoip-Latitude", fmt.Sprintf("%f", tc.userCoord.Latitude))
				r.Header.Set("X-Geoip-Longitude", fmt.Sprintf("%f", tc.userCoord.Longitude))
			}

			h.ServeHTTP(rw, r)

			if rw.Code != http.StatusOK {
				t.Fatalf("Unexpected HTTP status code: got %d, want %d", rw.Code, http.StatusOK)
			}
			wantStr := strings.Join(tc.wantNodes, ",")
			if got := rw.Body.String(); got != wantStr {
				t.Errorf("Unexpected response: got %q, want %q", got, wantStr)
			}
		})
	}

}
