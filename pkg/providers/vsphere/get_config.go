package vsphere

import (
	"github.com/ilackarms/unik/pkg/providers"
)

func (p *VsphereProvider) GetConfig() providers.ProviderConfig {
	return providers.ProviderConfig{
		UsePartitionTables: true,
	}
}
