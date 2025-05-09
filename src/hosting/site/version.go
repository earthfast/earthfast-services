package site

import (
	"net/http"

	"armada-node/model"
)

type notFoundVersion struct{}

func (v notFoundVersion) IsProject(p *model.Project) bool {
	return p == nil || p.Content == ""
}

func (v notFoundVersion) Start() {}

func (v notFoundVersion) Stop() {}

func (v notFoundVersion) Delete() error { return nil }

func (v notFoundVersion) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (v notFoundVersion) String() string {
	return "HTTPStatusCodeSite[code=404]"
}
