package repo

import (
	"fmt"
	"strings"
)

// RepositoryType defines a type for all supported repositories.
type RepositoryType int

const (
	// GitHub cloud.
	GitHub RepositoryType = iota
)

// RepositoryTypeToString map associating type to its string representation.
var RepositoryTypeToString = map[RepositoryType]string{
	GitHub: "github",
}

// RepositoryTypeToEnum map associating string representation with type.
var RepositoryTypeToEnum = map[string]RepositoryType{
	"github": GitHub,
}

// Provider defines the common interface for different repository managers (e.g., GitHub)
type Provider interface {
	// Clone a given repository to a path
	Clone(repoURL string, outputPath string) error
}

// NewRepoProvider factory method to instantiate a repository provider for a given system.
func NewRepoProvider(repoProviderName string) (Provider, error) {
	provider, exists := RepositoryTypeToEnum[strings.ToLower(repoProviderName)]
	if !exists {
		return nil, fmt.Errorf("Provider not found for %s", repoProviderName)
	}
	switch provider {
	case GitHub:
		return NewGitHubCmdProvider()
	}
	return nil, fmt.Errorf("No provider implementation found for %s", repoProviderName)
}
