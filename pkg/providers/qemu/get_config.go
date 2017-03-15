package qemu

import (
	"github.com/ilackarms/unik/pkg/providers"
)

func (p *QemuProvider) GetConfig() providers.ProviderConfig {
	return providers.ProviderConfig{
		UsePartitionTables: true,
	}
}
