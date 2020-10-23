package manager

import (
	"github.com/dhiguero/grpc-proto-manager/internal/app/gpm/config"
	"github.com/rs/zerolog/log"
)

// GPM structure with the manager main loop.
type GPM struct {
	cfg config.ServiceConfig
}

// NewManager creates a new GPM entity.
func NewManager(cfg config.ServiceConfig) *GPM {
	return &GPM{cfg: cfg}
}

// Run triggers the execution of the command.
func (gpm *GPM) Run(basePath string) error {
	log.Debug().Msg("Launching GPM")
	if err := gpm.cfg.IsValid(); err != nil {
		log.Fatal().Err(err).Msg("invalid configuration options")
	}
	gpm.cfg.Print()
	return nil
}
