package domain

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"armada-node/model"
)

type DomainToProjectMapping struct {
	URL       string `json:"url"`
	ProjectID string `json:"projectId"`
}

// DownloadAndParseJSON downloads a JSON file from a public URL and parses it.
func DownloadAndParseJSON(url string) (map[string]model.ID, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	// Create client with custom transport
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var projects []DomainToProjectMapping
	if err := json.Unmarshal(body, &projects); err != nil {
		return nil, err
	}

	projectsMap := make(map[string]model.ID)
	for _, project := range projects {
		projectID, err := model.ParseID(project.ProjectID)
		if err != nil {
			return nil, fmt.Errorf("parsing project ID: %v", err)
		}
		projectsMap[project.URL] = projectID
		fmt.Printf("Mapped domain %s to project %s\n", project.URL, projectID.Hex())
	}

	return projectsMap, nil
}
