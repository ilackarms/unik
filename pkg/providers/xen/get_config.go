package xen

import (
	"github.com/ilackarms/unik/pkg/providers"
)

func (p *XenProvider) GetConfig() providers.ProviderConfig {
	return providers.ProviderConfig{
		UsePartitionTables: false,
	}
}
