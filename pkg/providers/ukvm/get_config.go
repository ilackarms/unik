package ukvm

import (
	"github.com/ilackarms/unik/pkg/providers"
)

func (p *UkvmProvider) GetConfig() providers.ProviderConfig {
	return providers.ProviderConfig{
		UsePartitionTables: true,
	}
}
