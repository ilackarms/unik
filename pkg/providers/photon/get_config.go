package photon

import (
	"github.com/ilackarms/unik/pkg/providers"
)

func (p *PhotonProvider) GetConfig() providers.ProviderConfig {
	return providers.ProviderConfig{
		UsePartitionTables: true,
	}
}
