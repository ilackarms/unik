package gcloud

import (
	"github.com/ilackarms/unik/pkg/providers"
)

func (p *GcloudProvider) GetConfig() providers.ProviderConfig {
	return providers.ProviderConfig{
		UsePartitionTables: false,
	}
}
