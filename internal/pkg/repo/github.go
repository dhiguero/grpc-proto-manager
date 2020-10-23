package repo

import (
	"fmt"
	"os/exec"

	"github.com/rs/zerolog/log"
)

// GitHubCmdProvider structure with the implementation to manage a GitHub system.
// This implementation relies on leveraging the existing git command on the system. In the
// future another provider making use of the API or golang SDK should be added :)
type GitHubCmdProvider struct {
}

// NewGitHubCmdProvider creates a new provider connecting to GitHub.
func NewGitHubCmdProvider() (Provider, error) {
	log.Debug().Msg("Using GitHubCmdProvider")
	return &GitHubCmdProvider{}, nil
}

// Clone a given repository to a path
// git clone git@github.com:whatever folder-name
func (gh *GitHubCmdProvider) Clone(repoURL string, outputPath string) error {
	// TODO Check output path exists.
	cmdArgs := []string{"clone", fmt.Sprintf("git@%s", repoURL), outputPath}

	cmd := exec.Command("git", cmdArgs...)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("unable to clone repo %s due to %w", repoURL, err)
	}
	log.Debug().Str("output", string(stdoutStderr)).Msg("repo successfully cloned")
	return nil
}
