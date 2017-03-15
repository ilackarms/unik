package openstack

import (
	"github.com/ilackarms/unik/pkg/providers"
)

func (p *OpenstackProvider) GetConfig() providers.ProviderConfig {
	return providers.ProviderConfig{
		UsePartitionTables: true,
	}
}
