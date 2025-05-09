package model

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

var (
	ZeroID = ID{}
)

type ID [32]byte

func (id ID) IsZero() bool {
	return id == ZeroID
}

func (id ID) Hex() string {
	return hex.EncodeToString(id[:])
}

// ProjectType represents the type of project deployment
type ProjectType string

const (
	ProjectTypeStatic ProjectType = "static"
	ProjectTypeNextJS ProjectType = "nextjs"
)

type ProjectMetadata struct {
	Type      ProjectType `json:"type"`
	Port      int         `json:"port,omitempty"`
	BundleURL string      `json:"bundleUrl,omitempty"`
}

// ParseMetadata parses the project metadata string into structured data
func (p *Project) ParseMetadata() (*ProjectMetadata, error) {
	var metadata ProjectMetadata
	if p.Metadata == "" {
		return &metadata, nil
	}
	err := json.Unmarshal([]byte(p.Metadata), &metadata)
	return &metadata, err
}

type Project struct {
	ID       ID
	Name     string
	Content  string
	Checksum string
	Metadata string
	// TODO: Consider adding:
	// ParsedMetadata ProjectMetadata // Parsed version of Metadata
}

type Node struct {
	ID        ID
	Host      string
	ProjectID ID
	Region    string
}

func (a *Node) Equal(b *Node) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.ID == b.ID && a.Host == b.Host && a.ProjectID == b.ProjectID && a.Region == b.Region
}

type Client interface {
	GetNode(ctx context.Context, id ID) (*Node, error)
	GetProject(ctx context.Context, id ID) (*Project, error)

	ContentNodes(ctx context.Context, projectID ID) ([]*Node, error)
}

// ParseID converts a hexadecimal string representation of an ID into an ID.
// The input string may optionally be prefixed with "0x".
func ParseID(s string) (ID, error) {
	s = strings.TrimPrefix(s, "0x")
	idBytes, err := hex.DecodeString(s)
	if err != nil {
		return ZeroID, fmt.Errorf("decoding as hexadecimal: %v", err)
	}
	if n := len(idBytes); n != 32 {
		return ZeroID, fmt.Errorf("invalid length: got %d bytes, want 32", n)
	}
	var id ID
	copy(id[:], idBytes[:])
	return id, nil
}
