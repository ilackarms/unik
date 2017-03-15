package aws

import (
	"github.com/ilackarms/unik/pkg/providers"
)

func (p *AwsProvider) GetConfig() providers.ProviderConfig {
	return providers.ProviderConfig{
		UsePartitionTables: false,
	}
}
