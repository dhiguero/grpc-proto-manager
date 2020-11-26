package repo

import (
	"github.com/rs/zerolog/log"
)

// GitHubCmdProvider structure with the implementation to manage a GitHub system.
// This implementation relies on leveraging the existing git command on the system. In the
// future another provider making use of the API or golang SDK should be added :)
type GitHubCmdProvider struct {
	GHCommon
}

// NewGitHubCmdProvider creates a new provider connecting to GitHub.
func NewGitHubCmdProvider() (Provider, error) {
	log.Debug().Msg("Using GitHubCmdProvider")
	return &GitHubCmdProvider{
		GHCommon: GHCommon{
			UseSSH:            true,
			SetPusherUserName: false,
			SetPusherEmail:    false,
		},
	}, nil
}

// ConfigurePusher prepares the system to use a particular username/email to appear as the pusher of the commits.
func (gh *GitHubCmdProvider) ConfigurePusher(username string, email string) error {
	log.Debug().Str("username", username).Str("email", email).Msg("setting pusher information")
	// Notice that in GitHub, it is recommended to setup this information per repository, therefore this action
	// will be executed on per-repo basis before the commit & push information.

	// To setup this, the following commands will be issued:
	// git config -f /tmp/grpc-internal-agenda-go/.git/config user.name \"Your Name\"
	// git config -f /tmp/grpc-internal-agenda-go/.git/config user.email "my.name@server.com"

	// Alternatively, each git command may be configured with a given user name. However, given that this is executed from
	// a local temporal copy, there should not be any collateral impact in configuring the local repo. The alternative would be:
	// git -c user.name="Your name" -c user.email="my.name@server.com" commit -m "Commit message" ...
	gh.SetPusherUserName = true
	gh.PusherUserName = username
	gh.SetPusherEmail = true
	gh.PusherEmail = email

	return nil
}
