package virtualbox

import (
	"github.com/ilackarms/unik/pkg/providers"
)

func (p *VirtualboxProvider) GetConfig() providers.ProviderConfig {
	return providers.ProviderConfig{
		UsePartitionTables: true,
	}
}
