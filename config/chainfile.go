package config

import (
	"time"

	"github.com/olympus-protocol/ogen/primitives"
)

// ChainFile represents the on-disk chain file used to initialize the chain.
type ChainFile struct {
	Validators         []primitives.ValidatorInitialization `json:"validators"`
	GenesisTime        int64                                `json:"genesis_time"`
	InitialConnections []string                             `json:"initial_connections"`
}

// ToInitializationParameters converts the chain configuration file to initialization
// parameters.
func (cf *ChainFile) ToInitializationParameters() primitives.InitializationParameters {
	ip := primitives.InitializationParameters{
		InitialValidators: cf.Validators,
		GenesisTime:       time.Unix(cf.GenesisTime, 0),
	}

	if cf.GenesisTime == 0 {
		ip.GenesisTime = time.Now().Add(5 * time.Second)
	}

	return ip
}
